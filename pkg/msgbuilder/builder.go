package msgbuilder

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
)

const (
	tmpl = `Subject: Вакансия {{.Name}}` + "\r\nMIME-Version: 1.0\r\n" +
		"Content-Type: multipart/mixed; boundary={{.Placeholder}}\r\n" +
		"\r\n--{{.Placeholder}}\r\n" +
		"Content-Type: text/plain\r\n\r\n" + `
Имя: {{.Name}}
E-Mail: {{.Email}}
Телефон: {{.Number}}
Интересует проект: {{.Project}}
{{if (ne .Attach.Filename "")}}` +
		"\r\n--{{.Placeholder}}\r\n" +
		"Content-Type: application/octet-stream; name=\"{{.Attach.Filename}}\"\r\n" +
		"Content-Transfer-Encoding: base64\r\n" +
		// "Content-Length: 64353\r\n" +
		"Content-Disposition: attachment; filename=\"{{.Attach.Filename}}\"\r\n\r\n" +
		"{{.Attach.Content}}{{end}}" +
		"\r\n--{{.Placeholder}}--\r\n\r\n"

	newTmpl = `Имя: {{.Name}}
E-Mail: {{.Email}}
Телефон: {{.Number}}
Интересует проект: {{.Project}}`
)

// message ...
type message struct {
	Subject string
	Name    string
	Email   string
	Project string
	Number  string
	Attach  Attachment
}

// Message ...
type Message struct {
	Subject string
	Message []byte
	Attach  Attachment
}

// Attachment ...
type Attachment struct {
	Filename string
	Content  io.Reader
}

// NewAttachment ...
func NewAttachment(filename string, file io.Reader) Attachment {
	if file == nil {
		return Attachment{}
	}

	return Attachment{
		Filename: filename,
		Content:  file,
	}
}

// GetMessage ..
func GetMessage(name, email, number, project string, attach Attachment) (Message, error) {
	msg := message{
		Name:    name,
		Email:   email,
		Number:  number,
		Project: project,
		Attach:  attach,
	}

	var tpl bytes.Buffer

	t := template.Must(template.New("config").Parse(newTmpl))

	err := t.Execute(&tpl, msg)
	if err != nil {
		return Message{}, fmt.Errorf("execute error: %+v", err)
	}

	return Message{
		Subject: fmt.Sprintf("Вакансия %s", name),
		Message: tpl.Bytes(),
		Attach:  attach,
	}, nil
}
