package server

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/twoflower3/interview-service/pkg/datastore"
	"github.com/twoflower3/interview-service/pkg/smtp"
)

// Server dummy
type Server struct {
	*http.Server
}

// Options dummy
type Options struct {
	WriteTimeout time.Duration
	Address      string

	SmptHostname     string
	MailFromLogin    string
	MailFromPassword string
	MailTo           string
}

// New create server
func New(options Options) (*Server, error) {
	s := Server{
		Server: &http.Server{
			Addr: options.Address,
		},
	}

	log.Debugf("options: %+v", options)

	datastore.ClientSMTP = smtp.NewSMTP(
		options.SmptHostname,
		options.MailFromLogin,
		options.MailFromPassword,
		options.MailTo,
	)

	return &s, nil
}

// Start server
func (s *Server) Start() error {
	s.registerHandler()
	// server.log.WithField("Address", server.Address()).Info("Starting server")
	return s.ListenAndServe()
}

// Shutdown server
func (s *Server) Shutdown(sec time.Duration) error {
	// ctx, cancel := context.WithTimeout(context.Background(), sec*time.Second)
	// defer cancel()
	// return s.Shutdown(ctx)
	return nil
}

// Address of server
func (s *Server) Address() string {
	return s.Address()
}

func (s *Server) registerHandler() {
	s.Handler = createRouter(s)
}
