package main

import (
	"feira-api/config"
	"feira-api/server"

	"github.com/sirupsen/logrus"
)

func main() {
	config.Start()
	port := "8080"

	if err := server.Run(port); err != nil {
		logrus.Errorf("error running server: %+v", err)
	}
}
