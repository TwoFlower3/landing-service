package smtp

import (
	"fmt"
	"io"
	"net/smtp"

	"github.com/twoflower3/interview-service/pkg/login"
	msg "github.com/twoflower3/interview-service/pkg/msgbuilder"

	lib "github.com/AeroNotix/libsmtp"
)

const (
	smtpPort = "587"
)

// SMTP ...
type SMTP struct {
	hostname string

	fromLogin    string
	fromPassword string
	to           string
}

// NewSMTP ...
func NewSMTP(hostname, fromLogin, fromPassword, to string) *SMTP {
	return &SMTP{
		hostname:     hostname,
		fromLogin:    fromLogin,
		fromPassword: fromPassword,
		to:           to,
	}
}

// SendMessage ...
func (client *SMTP) SendMessage(message msg.Message) error {

	if err := client.checkConnection(); err != nil {
		return fmt.Errorf("check connection error: %+v", err)
	}

	if err := client.send(message); err != nil {
		return fmt.Errorf("send error: %+v", err)
	}

	return nil
}

func (client *SMTP) checkConnection() error {
	_, err := smtp.Dial(smtpAddress(client.hostname, smtpPort))
	if err != nil {
		return fmt.Errorf("connection to smtp server error: %+v", err)
	}

	return nil
}

func (client *SMTP) send(message msg.Message) error {
	auth := login.NewAuth(client.fromLogin, client.fromPassword)
	to := []string{client.to}

	attachment := func() map[string]io.Reader {
		att := map[string]io.Reader{}

		if message.Attach.Filename == "" {
			return att
		}

		att[message.Attach.Filename] = message.Attach.Content
		return att
	}

	if err := lib.SendMailWithAttachments(smtpAddress(client.hostname, smtpPort),
		&auth,
		client.fromLogin,
		message.Subject,
		to,
		message.Message,
		attachment(),
	); err != nil {
		return fmt.Errorf("send mail error: %+v", err)
	}

	return nil
}

func smtpAddress(hostname, port string) string {
	return fmt.Sprintf("%s:%s", hostname, smtpPort)
}
