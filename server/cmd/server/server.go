package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"olycall-server/internal/core"
	"olycall-server/internal/in/rest"
	"olycall-server/pkg/ctxlogger"
	"olycall-server/pkg/redis"

	"github.com/jackc/pgx/v5/pgxpool"

	chatStorePostgres "olycall-server/internal/out/chatstore/postgres"
	connectionStoreInmemory "olycall-server/internal/out/connectionstore/inmemory"
	googleOAuthProviderHttp "olycall-server/internal/out/googleoauthprovider/http"
	notificationsProviderWebsocket "olycall-server/internal/out/notificationsprovider/websocket"
	oAuthStateStoreRedis "olycall-server/internal/out/oauthstatestore/redis"
	userStorePostgres "olycall-server/internal/out/userstore/postgres"
)

//	@title		Server API
//	@version	0.1

//	@host		localhost:8080
//	@BasePath	/api

//	@Accept		json
//	@Produce	json

func run(ctx context.Context, cfg startCmd) error {
	logLevel := slog.LevelInfo
	if err := logLevel.UnmarshalText([]byte(strings.ToUpper(cfg.LogLevel))); err != nil {
		return fmt.Errorf("failed to parse log level: %w", err)
	}

	h := ctxlogger.ContextHandler{
		Handler: slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: logLevel,
		}),
	}

	logger := slog.New(&h)

	config, err := pgxpool.ParseConfig(cfg.DSN)
	if err != nil {
		return fmt.Errorf("parse DSN: %w", err)
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return fmt.Errorf("connect to database: %w", err)
	}
	defer pool.Close()

	redisClient, err := redis.NewClient(
		ctx,
		cfg.RedisHost,
		cfg.RedisPort,
		cfg.RedisPassword,
		cfg.RedisDB,
	)
	if err != nil {
		return fmt.Errorf("failed to create redis client: %w", err)
	}
	defer redisClient.Close()

	oAuthStateStore := oAuthStateStoreRedis.NewOAuthStateStore(redisClient)

	userStore := userStorePostgres.New(pool)

	chatStore := chatStorePostgres.New(pool)
	connectionStore := connectionStoreInmemory.New()

	googleOAuthProvider := googleOAuthProviderHttp.NewGoogleOAuthProvider(
		cfg.GoogleOauth2ID,
		cfg.GoogleOauth2Secret,
		cfg.GoogleOauth2RedirectURL,
	)

	notificationsProvider := notificationsProviderWebsocket.New()

	svc := core.NewService(
		userStore,
		chatStore,
		oAuthStateStore,
		googleOAuthProvider,
		notificationsProvider,
		connectionStore,
		cfg.Secret,
	)

	controller := rest.NewController(
		svc,
		connectionStore,
		logger,
	)

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		Handler:           controller.GetMux(),
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Run the server in a goroutine
	srvErr := make(chan error, 1)

	go func() {
		logger.InfoContext(ctx, "starting server", "address", srv.Addr)

		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			srvErr <- fmt.Errorf("listen and serve: %w", err)
		} else {
			srvErr <- nil
		}
	}()

	select {
	case <-stop:
		logger.InfoContext(ctx, "shutdown signal received")

		// Context with timeout for graceful shutdown
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			return fmt.Errorf("shutdown: %w", err)
		}

		return nil

	case err := <-srvErr:
		return err
	}
}
