package myLog

import (
	"fmt"
	myLogger "github.com/aqualang97/logger/v4"
	"os"
)

func LogInit() *myLogger.Logger {
	//logFile, err := os.Create("/home/NIX/student/delivery/logs/logfile")

	logFile, err := os.OpenFile("/home/NIX/student/delivery/logs/newLogFile", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		fmt.Println(err)
		fmt.Println(err)
	}
	defer logFile.Close()

	logger := myLogger.NewLogger(logFile)
	logger.Info("---START---")
	logger.InfoConsole("---START---")
	return &logger
}
