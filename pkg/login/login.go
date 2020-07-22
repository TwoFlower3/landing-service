package login

import (
	"fmt"
	"net/smtp"
)

// Auth ...
type Auth struct {
	username, password string
}

// NewAuth ...
func NewAuth(username, password string) smtp.Auth {
	return &Auth{username, password}
}

// Start ...
func (a *Auth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

// Next ...
func (a *Auth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, fmt.Errorf("Unknown from server")
		}
	}
	return nil, nil
}
