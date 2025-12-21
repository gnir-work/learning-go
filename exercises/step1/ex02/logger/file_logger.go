package logger

import (
	"os"
)

func NewFileLogger(file *os.File) *IOLogger {
	return NewIOLogger(file)
}
