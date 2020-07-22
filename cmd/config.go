package main

import (
	cli "github.com/urfave/cli/v2"
)

const (
	FLAG_DEBUG            = "debug"
	FLAG_TEXTLOG          = "textlog"
	FLAG_TRACE            = "trace"
	FLAG_PORT             = "port"
	FLAG_HOST             = "host"
	FLAG_SMTP_HOST        = "smtp"
	FLAG_ACCOUNT          = "account"
	FLAG_ACCOUNT_PASSWORD = "account-password"
	FLAG_SEND_TO          = "send-to"
)

var flags = []cli.Flag{
	&cli.BoolFlag{
		EnvVars: []string{"DEBUG"},
		Name:    FLAG_DEBUG,
		Value:   true,
		Usage:   "start the server in debug mode",
	},
	&cli.BoolFlag{
		EnvVars: []string{"TEXTLOG"},
		Name:    FLAG_TEXTLOG,
		Value:   true,
		Usage:   "output log in text format",
	},
	&cli.BoolFlag{
		EnvVars: []string{"TRACE"},
		Name:    FLAG_TRACE,
		Usage:   "enable trace in output log",
	},
	&cli.StringFlag{
		EnvVars: []string{"HOST"},
		Name:    FLAG_HOST,
		Value:   "",
		Usage:   "Server address",
	},
	&cli.StringFlag{
		EnvVars: []string{"PORT"},
		Name:    FLAG_PORT,
		Value:   "8080",
		Usage:   "Server port",
	},
	&cli.StringFlag{
		EnvVars: []string{"SMTP_HOSTNAME"},
		Name:    FLAG_SMTP_HOST,
		Value:   "",
		Usage:   "smtp hostname",
	},
	&cli.StringFlag{
		EnvVars: []string{"LOGIN"},
		Name:    FLAG_ACCOUNT,
		Value:   "",
		Usage:   "login from",
	},
	&cli.StringFlag{
		EnvVars: []string{"PASSWORD"},
		Name:    FLAG_ACCOUNT_PASSWORD,
		Value:   "",
		Usage:   "password from",
	},
	&cli.StringFlag{
		EnvVars: []string{"SEND_MAIL"},
		Name:    FLAG_SEND_TO,
		Value:   "",
		Usage:   "mail to send",
	},
}
