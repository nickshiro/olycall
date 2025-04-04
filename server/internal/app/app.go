package app

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/xhhx-space/olycall-server/internal/controller"
	"github.com/xhhx-space/olycall-server/internal/service"
	"github.com/xhhx-space/olycall-server/pkg/ctxlogger"
)

type Config struct {
	Port     int    `json:"port"`
	LogLevel string `json:"log_level"`
}

//	@title		Server API
//	@version	0.1
//
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//
//	@host		localhost:8080
//	@BasePath	/api
//
//	@Accept		json
//	@Produce	json

func Run(ctx context.Context, cfg Config) error {
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

	s := service.New()

	controller := controller.NewController(s, logger)

	server := &http.Server{
		Addr:        fmt.Sprintf(":%d", cfg.Port),
		Handler:     controller.GetMux(),
		ReadTimeout: time.Minute,
	}

	// graceful shutdown
	serverCtx, serverStopCtx := context.WithCancel(ctx)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownCtx, c := context.WithTimeout(serverCtx, time.Minute)
		defer c()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	logger.InfoContext(serverCtx, "starting server", "address", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to serve: %w", err)
	}

	<-serverCtx.Done()

	return nil
}
