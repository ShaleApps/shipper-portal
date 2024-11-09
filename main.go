package main

import (
	"github.com/ShaleApps/go-service-utils/helpers"
	"github.com/ShaleApps/{{SERVICE_NAME}}/app"
	"github.com/ShaleApps/{{SERVICE_NAME}}/internal/config"
	"github.com/sirupsen/logrus"
	"strings"
)

var Version = "0.1.0"

func init() {
	// Configure logger
	logLevel, err := logrus.ParseLevel(strings.ToLower(helpers.GetEnv("LOG_LEVEL", "INFO")))
	if err != nil {
		panic(err)
	}
	logrus.SetLevel(logLevel)

	fieldMap := logrus.FieldMap{
		logrus.FieldKeyTime:  "@log_timestamp",
		logrus.FieldKeyLevel: "@severity",
		logrus.FieldKeyMsg:   "@message",
		logrus.FieldKeyFunc:  "@caller",
	}

	var formatter = &logrus.JSONFormatter{
		FieldMap: fieldMap,
	}
	logrus.SetFormatter(formatter)
	logrus.SetReportCaller(true)
}

func main() {

	logrus.Infof("Starting {{SERVICE_NAME}} server v%s", Version)

	cnf := config.LoadConfig()

	app.StartApp(cnf)
}
