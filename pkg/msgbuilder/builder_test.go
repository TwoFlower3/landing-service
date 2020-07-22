package msgbuilder

import (
	"os"
	"testing"
)

func TestNewAttachment(t *testing.T) {
	attach := NewAttachment("", nil)
	if attach.Filename == "" {
		t.Fatal("attachment is empty")
	}
}

func TestNewMessage(t *testing.T) {
	name := "test"
	email := "test@test"
	number := "89991231212"
	project := "testProject"
	filename := "dasd"

	file, _ := os.Open("")

	attach := NewAttachment(filename, file)
	if attach.Filename == "" {
		t.Fatal("attachment is empty")
	}

	if _, err := GetMessage(
		name,
		email,
		number,
		project,
		attach,
	); err != nil {
		t.Fatal(err)
	}
}

func TestNewMessageWithoutAttachment(t *testing.T) {
	name := "test"
	email := "test@test"
	number := "89991231212"
	project := "testProject"

	attach := Attachment{}

	if _, err := GetMessage(
		name,
		email,
		number,
		project,
		attach,
	); err != nil {
		t.Fatal(err)
	}
}
