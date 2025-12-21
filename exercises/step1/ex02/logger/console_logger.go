package logger

import (
	"os"
)

func NewConsoleLogger() *IOLogger {
	return NewIOLogger(os.Stdout)
}
