package main

import (
	"github.com/jeffersonto/feira-api/config"
	"github.com/jeffersonto/feira-api/server"

	"github.com/sirupsen/logrus"
)

func main() {
	config.Start()
	port := "8080"

	if err := server.Run(port); err != nil {
		logrus.Errorf("error running server: %+v", err)
	}
}
