package main

import (
	"github.com/kosovo21/go-rest-mongo/pkg/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		DisableColors: true,
	})
	logrus.Info("server starting")

	err := server.RunServer()
	if err != nil {
		logrus.Error("error running server", err)
	}
}
