package logger

import (
	"io"
	"log"
	"os"
	"path"

	"github.com/Skryvvara/common-services/config"
	"github.com/natefinch/lumberjack"
)

func Initialize() {
	logDir := config.App.Log.LogDir
	logFile := config.App.Log.LogFile
	logPath := path.Join(logDir, logFile)
	permissions := os.O_RDWR | os.O_CREATE | os.O_APPEND

	if _, err := os.Stat(logPath); err != nil {
		if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
			log.Panic(err)
		}
	}

	file, err := os.OpenFile(logPath, permissions, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	mw := io.MultiWriter(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    config.App.Log.MaxSize,
		MaxBackups: config.App.Log.MaxBackups,
		MaxAge:     config.App.Log.MaxAge,
	}, os.Stderr)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(mw)
	log.Println("Successfully initialized logger. Using log file: " + logPath)
}
