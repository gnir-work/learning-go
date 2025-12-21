package logger

type MultiLogger struct {
	loggers []Logger
}

func NewMultiLogger(loggers ...Logger) *MultiLogger {
	return &MultiLogger{loggers: loggers}
}

func (logger *MultiLogger) Info(message string, args ...any) {
	for _, innerLogger := range logger.loggers {
		innerLogger.Info(message, args...)
	}
}

func (logger *MultiLogger) Error(message string, args ...any) {
	for _, innerLogger := range logger.loggers {
		innerLogger.Error(message, args...)
	}
}

func (logger *MultiLogger) Debug(message string, args ...any) {
	for _, innerLogger := range logger.loggers {
		innerLogger.Debug(message, args...)
	}
}
