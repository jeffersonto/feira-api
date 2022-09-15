package config

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var (
	LogLevel logrus.Level
)

func Start() {
	Newlogger()
}

func Newlogger() {
	var file, err = os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}
	logrus.SetOutput(file)

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return f.Function + "()", fmt.Sprintf("%s:%d", path.Base(f.File), f.Line)
		},
	})
	logrus.SetLevel(log("TRACE"))

	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func log(input string) logrus.Level {
	switch input {
	case "TRACE":
		return logrus.TraceLevel
	case "INFO":
		return logrus.InfoLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "DEBUG":
		return logrus.DebugLevel
	case "WARNING":
		return logrus.WarnLevel
	}
	return logrus.InfoLevel
}
