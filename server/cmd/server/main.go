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

// nolint: lll
type startCmd struct {
	DSN                     string `env:"OC_SERVER_DSN"                        help:"DSN"                         json:"-"              required:""`
	RedisPort               int    `default:"6379"                             env:"OC_SERVER_REDIS_PORT"         help:"Redis port"     json:"redis_port"`
	RedisHost               string `default:"localhost"                        env:"OC_SERVER_REDIS_HOST"         help:"Redis host"     json:"redis_host"`
	RedisPassword           string `env:"OC_SERVER_REDIS_PASSWORD"             help:"Redis password"              json:"-"`
	RedisDB                 int    `default:"0"                                env:"OC_SERVER_REDIS_DB"           help:"Redis database" json:"redis_db"`
	GoogleOauth2ID          string `env:"OC_SERVER_GOOGLE_OAUTH2_ID"           help:"Google OAuth2 client ID"     json:"-"              name:"google-oauth2-id"           required:""`
	GoogleOauth2Secret      string `env:"OC_SERVER_GOOGLE_OAUTH2_SECRET"       help:"Google OAuth2 client secret" json:"-"              name:"google-oauth2-secret"       required:""`
	GoogleOauth2RedirectURL string `env:"OC_SERVER_GOOGLE_OAUTH2_REDIRECT_URL" help:"Google OAuth2 redirect URL"  json:"-"              name:"google-oauth2-redirect-url" required:""`
	Secret                  string `env:"OC_SERVER_SECRET"                     help:"Server secret"               json:"secret"         required:""`
	Port                    int    `default:"8080"                             env:"OC_SERVER_PORT"               help:"Server port"    json:"port"`
	LogLevel                string `default:"INFO"                             env:"OC_SERVER_LOG_LEVEL"          help:"Logging level"  json:"log_level"`
}

type migrateCmd struct {
	DSN  string `env:"OT_SERVER_DSN"                    help:"DSN" json:"-" required:""`
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
		//nolint: err113
		return errors.New("no migration direction specified. Use --up or --down")
	}

	return nil
}
