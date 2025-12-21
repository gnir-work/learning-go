package logger

import (
	"testing"
	"time"
)

func getFrozenLogger() LogicalLogger {
	return LogicalLogger{
		nowFunc: func() time.Time {
			return time.Date(2025, 12, 20, 10, 0, 0, 0, time.UTC)
		},
	}
}

func TestLogicalLogger_Info(t *testing.T) {
	logger := getFrozenLogger()
	log := logger.Info("The message with formatting %q", "Some templated data")
	if expected := "[INFO] [2025-12-20T10:00:00Z] The message with formatting \"Some templated data\""; log != expected {
		t.Fatalf("Expected %q got %q", expected, log)
	}
}

func TestLogicalLogger_Error(t *testing.T) {
	logger := getFrozenLogger()
	log := logger.Error("The message with formatting %q", "Some templated data")
	if expected := "[ERROR] [2025-12-20T10:00:00Z] The message with formatting \"Some templated data\""; log != expected {
		t.Fatalf("Expected %q got %q", expected, log)
	}
}

func TestLogicalLogger_Debug(t *testing.T) {
	logger := getFrozenLogger()
	log := logger.Debug("The message with formatting %q", "Some templated data")
	if expected := "[DEBUG] [2025-12-20T10:00:00Z] The message with formatting \"Some templated data\""; log != expected {
		t.Fatalf("Expected %q got %q", expected, log)
	}
}
