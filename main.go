package main

import "github.com/r-evolve/logger/logger"

func main() {
	logger := logger.NewLogger(logger.LoggerConfig{
		Prefix:        "main",
		WithTimestamp: false,
		Target:        logger.Console,
		KillOnFatal:   true,
	})

	logger.Debug("Hello, playground")
	logger.Info("Hello, playground")
	// logger.Fatal("Hello, playground")
	logger.Warn("Hello, playground")
	logger.Error("Hello, playground")
	// logger.Fatal("Hello, playground")

	logger.Debugf("Hello, %s", "playground")
}
