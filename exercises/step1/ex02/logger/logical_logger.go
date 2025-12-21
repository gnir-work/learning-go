package logger

import (
	"fmt"
	"time"
)

type LogicalLogger struct {
	nowFunc func() time.Time
}

type LogLevel int

const (
	ERROR LogLevel = iota
	INFO
	DEBUG
)

func logLevelToDisplay(level LogLevel) string {
	switch level {
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

func (logger *LogicalLogger) now() time.Time {
	if logger.nowFunc == nil {
		return time.Now()
	}
	return logger.nowFunc()
}

func (logger *LogicalLogger) log(logLevel LogLevel, message string, args ...any) string {
	templatedMessage := fmt.Sprintf(message, args...)
	return fmt.Sprintf("[%v] [%v] %v", logLevelToDisplay(logLevel), logger.now().Format(time.RFC3339), templatedMessage)
}

func (logger *LogicalLogger) Info(message string, args ...any) string {
	return logger.log(INFO, message, args...)
}

func (logger *LogicalLogger) Debug(message string, args ...any) string {
	return logger.log(DEBUG, message, args...)
}

func (logger *LogicalLogger) Error(message string, args ...any) string {
	return logger.log(ERROR, message, args...)
}
