package utils

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log zerolog.Logger

func InitLogger() {
	logDir := "logs"
	_ = os.MkdirAll(logDir, os.ModePerm)

	// File naming (daily)
	fileName := logDir + "/log_" + time.Now().Format("2006-01-02") + ".log"

	fileWriter := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    50, // MB
		MaxBackups: 10, // keep last 10 files
		MaxAge:     10, // days (auto delete)
		Compress:   true,
	}

	// 👇 Multi-output (file + console)
	multiWriter := zerolog.MultiLevelWriter(fileWriter, os.Stdout)

	Log = zerolog.New(multiWriter).With().Timestamp().Logger()
}
