package log

import (
	logrus "github.com/sirupsen/logrus"
)

var logger = logrus.New()

func LogInit() {
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})
}

func InfoError(e error, s string) {
	logger.WithFields(logrus.Fields{
		"origin_error": e,
	}).Info(s)
}
