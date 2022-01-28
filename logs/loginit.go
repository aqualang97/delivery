package myLog

import (
	"fmt"
	myLogger "github.com/aqualang97/logger/v4"
	"os"
)

func LogInit() myLogger.Logger {
	logFile, err := os.OpenFile("logs/logfile", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer logFile.Close()

	logger := myLogger.NewLogger(logFile)
	logger.Info("---START---")
	logger.InfoConsole("---START---")
	return logger
}
