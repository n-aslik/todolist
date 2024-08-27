package logger

import (
	"Todolist/configs"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

// SetLogger Установка Logger-а
var (
	Info  *log.Logger
	Error *log.Logger
	Warn  *log.Logger
	Debug *log.Logger
)

func Init() error {
	logParams := configs.AppSettings.LogParams
	if _, err := os.Stat(logParams.LogDirectory); os.IsNotExist(err) {
		err = os.Mkdir(logParams.LogDirectory, 0755)
		if err != nil {
			return err
		}
	}

	lumberLogInfo := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.LogDirectory, logParams.LogInfo),
		MaxSize:    logParams.MaxSizeMegabytes, // megabytes
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAge,   //days
		Compress:   logParams.Compress, // disabled by default
		LocalTime:  logParams.LocalTime,
	}

	lumberLogError := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.LogDirectory, logParams.LogError),
		MaxSize:    logParams.MaxSizeMegabytes, // megabytes
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAge,   //days
		Compress:   logParams.Compress, // disabled by default
		LocalTime:  logParams.LocalTime,
	}

	lumberLogWarn := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.LogDirectory, logParams.LogWarn),
		MaxSize:    logParams.MaxSizeMegabytes, // megabytes
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAge,   //days
		Compress:   logParams.Compress, // disabled by default
		LocalTime:  logParams.LocalTime,
	}

	lumberLogDebug := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", logParams.LogDirectory, logParams.LogDebug),
		MaxSize:    logParams.MaxSizeMegabytes, // megabytes
		MaxBackups: logParams.MaxBackups,
		MaxAge:     logParams.MaxAge,   //days
		Compress:   logParams.Compress, // disabled by default
		LocalTime:  logParams.LocalTime,
	}
	Info = log.New(gin.DefaultWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(lumberLogError, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(lumberLogWarn, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(lumberLogDebug, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	gin.DefaultWriter = io.MultiWriter(os.Stdout, lumberLogInfo)

	return nil
}
