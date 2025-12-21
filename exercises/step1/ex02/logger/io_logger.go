package logger

import (
	"io"
)

type Logger interface {
	Info(message string, args ...any)
	Error(message string, args ...any)
	Debug(message string, args ...any)
}

type IOLogger struct {
	output        io.Writer
	logicalLogger LogicalLogger
}

func NewIOLogger(output io.Writer) *IOLogger {
	logger := IOLogger{
		output:        output,
		logicalLogger: LogicalLogger{},
	}
	return &logger
}

func (logger *IOLogger) log(logLevel LogLevel, message string, args ...any) {
	_, _ = io.WriteString(logger.output, logger.logicalLogger.log(logLevel, message, args...))

}

func (logger *IOLogger) Info(message string, args ...any) {
	logger.log(INFO, message, args...)
}

func (logger *IOLogger) Debug(message string, args ...any) {
	logger.log(DEBUG, message, args...)
}

func (logger *IOLogger) Error(message string, args ...any) {
	logger.log(ERROR, message, args...)
}
