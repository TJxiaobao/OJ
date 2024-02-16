package logger

import (
	"log"
	"os"
	"path/filepath"
)

var logger *log.Logger

func init() {
	logFilePath := "./home/log/logfile.log"

	// 获取日志文件所在的目录路径
	logDir := filepath.Dir(logFilePath)

	// 创建日志文件所在的目录（如果不存在）
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		log.Fatal("Failed to create log directory:", err)
	}

	// 创建日志文件（如果不存在）
	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogToFile(message string) {
	logger.Println(message)
}
