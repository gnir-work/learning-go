package logger

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func getFrozenFileLogger(t *testing.T) (*IOLogger, string) {
	dir := t.TempDir()

	path := filepath.Join(dir, "input.txt")
	file, err := os.Create(path)
	if err != nil {
		t.Fatalf("failed to create file %q: %v", path, err)
	}
	t.Cleanup(func() {
		_ = file.Close()
	})

	return &IOLogger{
		output: file,
		logicalLogger: LogicalLogger{
			nowFunc: func() time.Time {
				return time.Date(2025, 12, 20, 10, 0, 0, 0, time.UTC)
			},
		},
	}, path
}

func TestFileLogger_Info(t *testing.T) {
	logger, file := getFrozenFileLogger(t)

	logger.Info("The message with formatting %q", "Some templated data")
	got, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("failed to read log file %v", err)
	}
	if expected := "[INFO] [2025-12-20T10:00:00Z] The message with formatting \"Some templated data\""; string(got) != expected {
		t.Fatalf("Expected %q got %q", expected, got)
	}
}

func TestFileLogger_Error(t *testing.T) {
	logger, file := getFrozenFileLogger(t)

	logger.Error("The message with formatting %q", "Some templated data")
	got, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("failed to read log file %v", err)
	}
	if expected := "[ERROR] [2025-12-20T10:00:00Z] The message with formatting \"Some templated data\""; string(got) != expected {
		t.Fatalf("Expected %q got %q", expected, got)
	}
}

func TestFileLogger_Debug(t *testing.T) {
	logger, file := getFrozenFileLogger(t)

	logger.Debug("The message with formatting %q", "Some templated data")
	got, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("failed to read log file %v", err)
	}
	if expected := "[DEBUG] [2025-12-20T10:00:00Z] The message with formatting \"Some templated data\""; string(got) != expected {
		t.Fatalf("Expected %q got %q", expected, got)
	}
}
