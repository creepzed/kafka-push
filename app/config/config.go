package config

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func OpenLogFile() *os.File {
	logFilePath := os.Getenv("KAFKAPUSH_LOGS_FILE_PATH")
	f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatalf("error opening logfile: %v error: %v", logFilePath, err)
	}
	log.SetFormatter(&log.JSONFormatter{})
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
	return f
}
