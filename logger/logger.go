package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	config LoggerConfig
}

type LoggerConfig struct {
	Prefix        string
	WithTimestamp bool
	Target        LoggerTarget
	KillOnFatal   bool
}

func NewLogger(config LoggerConfig) *Logger {
	return &Logger{config}
}

func (logger *Logger) log(level LogLevel, message string) {
	switch logger.config.Target {
	case Console:
		fmt.Println(logger.format(level, message))
	case File:
		// TODO: implement file logging
	}

	if level == FATAL && logger.config.KillOnFatal {
		// TODO: add log message on kill
		panic(message)
	}
}

func (logger *Logger) format(level LogLevel, message string) string {
	var formattedMessage string

	if logger.config.WithTimestamp {
		formattedMessage += fmt.Sprintf("%s ", time.Now().Format(time.RFC1123))
	}

	if logger.config.Prefix != "" {
		formattedMessage += fmt.Sprintf("[%s] ", logger.config.Prefix)
	}

	formattedMessage += fmt.Sprintf("[%s] %s", level, message)

	return formattedMessage
}

func (logger *Logger) Debug(message string) {
	logger.log(DEBUG, message)
}

func (logger *Logger) Info(message string) {
	logger.log(INFO, message)
}

func (logger *Logger) Warn(message string) {
	logger.log(WARN, message)
}

func (logger *Logger) Error(message string) {
	logger.log(ERROR, message)
}

func (logger *Logger) Fatal(message string) {
	logger.log(FATAL, message)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.log(DEBUG, fmt.Sprintf(format, args...))
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.log(INFO, fmt.Sprintf(format, args...))
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.log(WARN, fmt.Sprintf(format, args...))
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.log(ERROR, fmt.Sprintf(format, args...))
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	logger.log(FATAL, fmt.Sprintf(format, args...))
}
