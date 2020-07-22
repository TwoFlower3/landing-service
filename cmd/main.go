package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	cli "github.com/urfave/cli/v2"

	log "github.com/sirupsen/logrus"

	"github.com/twoflower3/interview-service/pkg/server"
)

var version string

const (
	defaultVersion = "v0.0.1-dev"
)

func main() {
	app := cli.NewApp()
	app.Name = "is"
	app.Usage = "Interview Service"
	app.Version = version
	app.Flags = flags
	app.Action = run

	if err := app.Run(os.Args); err != nil {
		log.WithError(err).Fatal("Runtime error")
		os.Exit(1)
	}
}

func run(ctx *cli.Context) error {
	if ctx.Bool(FLAG_DEBUG) {
		log.SetLevel(log.DebugLevel)
	}

	s, err := server.New(server.Options{
		Address: ctx.String(FLAG_HOST) + ":" + ctx.String(FLAG_PORT),

		SmptHostname:     ctx.String(FLAG_SMTP_HOST),
		MailFromLogin:    ctx.String(FLAG_ACCOUNT),
		MailFromPassword: ctx.String(FLAG_ACCOUNT_PASSWORD),
		MailTo:           ctx.String(FLAG_SEND_TO),
	})
	if err != nil {
		return fmt.Errorf("Initialize server error: %v", err)
	}

	// TODO: ADD exit or error
	go s.Start()
	waitShutdownSign()

	log.Info("Server shutdown initialized")
	return s.Shutdown(time.Second * 5)
}

func waitShutdownSign() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
