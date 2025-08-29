package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitLogger() *logrus.Logger {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)
	return log
}

func GetLogger() *logrus.Logger {
	return log
}
