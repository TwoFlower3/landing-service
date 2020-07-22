package smtp

import (
	"fmt"
	"os"
	"testing"

	msg "github.com/twoflower3/interview-service/pkg/msgbuilder"
)

func TestSendMessageWithoutAttachment(t *testing.T) {
	smtp := NewSMTP("", "", "", "")

	message, _ := msg.GetMessage(
		"test",
		"test@test",
		"89991231212",
		"test_project",
		msg.NewAttachment("", nil),
	)

	if err := smtp.SendMessage(message); err != nil {
		t.Fatalf("%+v", err)
	}
}

func TestSendMessage(t *testing.T) {
	smtp := NewSMTP("", "", "", "")

	file, err := getFile("")
	if err != nil {
		t.Fatal(err)
	}

	message, _ := msg.GetMessage(
		"name_test",
		"test@test",
		"89991231212",
		"test_project",
		msg.NewAttachment(file.Name(), file),
	)

	if err := smtp.SendMessage(message); err != nil {
		t.Fatalf("%+v", err)
	}
}

func getFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("file cant read %+v", err)
	}

	return file, nil
}
