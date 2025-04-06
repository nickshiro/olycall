package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/oauth2"

	"olycall-server/internal/controller/rest"
	redisCache "olycall-server/internal/repository/cache/redis"
	"olycall-server/internal/repository/domain/postgres"
	"olycall-server/internal/service"
	"olycall-server/pkg/ctxlogger"
	"olycall-server/pkg/redis"
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

	domainRepo := postgres.NewRepo(pool)

	redisClient, err := redis.NewRedisClient(&redis.Config{
		Port:     cfg.RedisPort,
		Host:     cfg.RedisHost,
		Password: cfg.RedisPassword,
	})
	if err != nil {
		return fmt.Errorf("failed to create redis client: %w", err)
	}
	defer redisClient.Close()

	cacheRepo := redisCache.NewRepo(redisClient)

	googleOauth2Config := oauth2.Config{
		ClientID:     cfg.GoogleOauth2ID,
		ClientSecret: cfg.GoogleOauth2Secret,
		RedirectURL:  cfg.GoogleOauth2RedirectURL,
	}

	s := service.New(
		domainRepo,
		cacheRepo,
		googleOauth2Config,
		cfg.Secret,
	)

	controller := rest.NewController(s, logger)
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		Handler:           controller.GetMux(),
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
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
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
