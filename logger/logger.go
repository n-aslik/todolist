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

	fileInfo, err := os.OpenFile(configs.AppSettings.LogParams.LogInfo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	fileError, err := os.OpenFile(configs.AppSettings.LogParams.LogError, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	fileWarn, err := os.OpenFile(configs.AppSettings.LogParams.LogWarn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	fileDebug, err := os.OpenFile(configs.AppSettings.LogParams.LogDebug, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	Info = log.New(fileInfo, "", log.Ldate|log.Lmicroseconds)
	Error = log.New(fileError, "", log.Ldate|log.Lmicroseconds)
	Warn = log.New(fileWarn, "", log.Ldate|log.Lmicroseconds)
	Debug = log.New(fileDebug, "", log.Ldate|log.Lmicroseconds)

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
