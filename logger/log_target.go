package logger

type LoggerTarget string

const (
	Console LoggerTarget = "console"
	File    LoggerTarget = "file"
	Both    LoggerTarget = "both"
)
