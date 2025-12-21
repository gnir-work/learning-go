package logger

import (
	"testing"
)

func TestMultiLogger_IO(t *testing.T) {
	firstLogger, firstBuffer := getFrozenIOLogger()
	secondLogger, secondBuffer := getFrozenIOLogger()
	logger := MultiLogger{
		loggers: []Logger{firstLogger, secondLogger},
	}
	logger.Info("The message with formatting %q", "Some templated data")
	for _, got := range []string{firstBuffer.String(), secondBuffer.String()} {
		if expected := "[INFO] [2025-12-20T10:00:00Z] The message with formatting \"Some templated data\""; got != expected {
			t.Fatalf("Expected %q got %q", expected, got)
		}
	}
}
