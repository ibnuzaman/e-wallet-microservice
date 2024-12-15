package helpers

import (
	"github.com/sirupsen/logrus"
)

func SetupLogger() *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	return log
}
