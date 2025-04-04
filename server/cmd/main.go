package main

import (
	"context"
	"log"

	"github.com/alecthomas/kong"
	"github.com/xhhx-space/olycall-server/internal/app"
)

type startCmd struct {
	Port     int    `help:"Server port" env:"OT_SERVER_PORT" default:"8080" json:"port"`
	LogLevel string `help:"Logging level" env:"OT_SERVER_LOG_LEVEL" default:"INFO" json:"log_level"`
}

func main() {
	var cliArgs struct {
		Start startCmd `cmd:"" help:"Start the server"`
	}
	ctx := kong.Parse(&cliArgs)

	switch ctx.Command() {
	case "start":
		if err := handleStart(cliArgs.Start); err != nil {
			log.Fatalln(err)
		}
	default:
		log.Fatalln("Unknown command.")
	}
}

func handleStart(cmd startCmd) error {
	log.Printf("Starting server on port %d with log level %s\n", cmd.Port, cmd.LogLevel)
	ctx := context.Background()
	return app.Run(ctx, app.Config(cmd)) //nolint: wrapcheck
}
