package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/alecthomas/kong"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type startCmd struct {
	DSN                     string `help:"DSN" env:"OC_SERVER_DSN" required:"" json:"-"`
	RedisPort               int    `help:"Redis port" env:"OC_SERVER_REDIS_PORT" default:"6379" json:"redis_port"`
	RedisHost               string `help:"Redis host" env:"OC_SERVER_REDIS_HOST" default:"localhost" json:"redis_host"`
	RedisPassword           string `help:"Redis password" env:"OC_SERVER_REDIS_PASSWORD" json:"-"`
	GoogleOauth2ID          string `help:"Google OAuth2 client ID" env:"OC_SERVER_GOOGLE_OAUTH2_ID" name:"google-oauth2-id" required:"" json:"-"`                        //nolint: lll
	GoogleOauth2Secret      string `help:"Google OAuth2 client secret" env:"OC_SERVER_GOOGLE_OAUTH2_SECRET" name:"google-oauth2-secret" required:"" json:"-"`            //nolint: lll
	GoogleOauth2RedirectURL string `help:"Google OAuth2 redirect URL" env:"OC_SERVER_GOOGLE_OAUTH2_REDIRECT_URL" name:"google-oauth2-redirect-url" required:"" json:"-"` //nolint: lll
	Secret                  string `help:"Server secret" env:"OC_SERVER_SECRET" required:"" json:"secret"`
	Port                    int    `help:"Server port" env:"OC_SERVER_PORT" default:"8080" json:"port"`
	LogLevel                string `help:"Logging level" env:"OC_SERVER_LOG_LEVEL" default:"INFO" json:"log_level"`
}

type migrateCmd struct {
	DSN  string `help:"DSN" env:"OT_SERVER_DSN" required:"" json:"-"`
	Up   bool   `help:"Apply all up migrations"`
	Down bool   `help:"Apply down migration (rollback)"`
}

func main() {
	var cliArgs struct {
		Start   startCmd   `cmd:"" help:"Start the server"`
		Migrate migrateCmd `cmd:"" help:"Run database migrations"`
	}
	ctx := kong.Parse(&cliArgs)

	switch ctx.Command() {
	case "start":
		if err := handleStart(cliArgs.Start); err != nil {
			log.Fatalln(err)
		}
	case "migrate":
		if err := handleMigrate(cliArgs.Migrate); err != nil {
			log.Fatalln(err)
		}
	default:
		log.Fatalln("Unknown command. Use 'server' or 'migrate'.")
	}
}

func handleStart(cmd startCmd) error {
	jsonConfig, _ := json.MarshalIndent(cmd, "", "  ")
	log.Println(string(jsonConfig))

	log.Printf("Starting server on port %d with log level %s\n", cmd.Port, cmd.LogLevel)
	ctx := context.Background()
	return run(ctx, cmd) //nolint: wrapcheck
}

func handleMigrate(cmd migrateCmd) error {
	log.Printf("Running migrations on database: %s\n", cmd.DSN)

	conn, err := sql.Open("pgx", cmd.DSN)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer conn.Close()

	driver, err := pgx.WithInstance(conn, &pgx.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations", // Migration files directory
		"postgres", driver,
	)
	if err != nil {
		return fmt.Errorf("failed to initialize migrations: %w", err)
	}

	switch {
	case cmd.Up:
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("migration failed: %w", err)
		}
		log.Println("Migrations applied successfully.")
	case cmd.Down:
		if err := m.Steps(-1); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("rollback failed: %w", err)
		}
		log.Println("Last migration rolled back.")
	default:
		return errors.New("no migration direction specified. Use --up or --down")
	}

	return nil
}
