package logger

import (
	"bytes"
	"io"
	"testing"
	"time"
)

func getFrozenIOLogger() (*IOLogger, *bytes.Buffer) {
	var buf bytes.Buffer // implements io.Writer
	w := io.Writer(&buf)
	return &IOLogger{
		output: w,
		logicalLogger: LogicalLogger{
			nowFunc: func() time.Time {
				return time.Date(2025, 12, 20, 10, 0, 0, 0, time.UTC)
			},
		},
	}, &buf
}

func TestIOLogger_Info(t *testing.T) {
	logger, buffer := getFrozenIOLogger()
	logger.Info("The message with formatting %q", "Some templated data")
	got := buffer.String()
	if expected := "[INFO] [2025-12-20T10:00:00Z] The message with formatting \"Some templated data\""; got != expected {
		t.Fatalf("Expected %q got %q", expected, got)
	}
}

func TestIOLogger_Error(t *testing.T) {
	logger, buffer := getFrozenIOLogger()
	logger.Error("The message with formatting %q", "Some templated data")
	got := buffer.String()
	if expected := "[ERROR] [2025-12-20T10:00:00Z] The message with formatting \"Some templated data\""; got != expected {
		t.Fatalf("Expected %q got %q", expected, got)
	}
}

func TestIOLogger_Debug(t *testing.T) {
	logger, buffer := getFrozenIOLogger()
	logger.Debug("The message with formatting %q", "Some templated data")
	got := buffer.String()
	if expected := "[DEBUG] [2025-12-20T10:00:00Z] The message with formatting \"Some templated data\""; got != expected {
		t.Fatalf("Expected %q got %q", expected, got)
	}
}
