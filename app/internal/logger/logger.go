package logger

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

func InitLogger(logDir string) (*slog.Logger, error) {
	var logger *slog.Logger
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.MkdirAll(logDir, 0755); err != nil {
			return nil, err
		}
	}

	logFileName := fmt.Sprintf("%s/log_%s.log", logDir, time.Now().Format("2006-01-02"))
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	handler := slog.NewJSONHandler(logFile, nil)
	logger = slog.New(handler)
	logger.Info("Logger created", "file", logFileName)
	return logger, nil
}
