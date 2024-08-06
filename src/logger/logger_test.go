package logger

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	var infoBuffer bytes.Buffer

	log.SetOutput(&infoBuffer)
	SetVerbose(true)
	Info("This is an info message")
	if !strings.Contains(infoBuffer.String(), "This is an info message\n") {
		t.Errorf("Expected 'This is an info message\n', but got '%s'", infoBuffer.String())
	}

	infoBuffer.Reset()
	SetVerbose(false)
	Info("This message should not appear")
	if infoBuffer.String() != "" {
		t.Errorf("Expected no output, but got '%s'", infoBuffer.String())
	}
}
