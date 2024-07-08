package console

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Entry

func Setup(service string) {

	apiKey := os.Getenv("DATADOG_API_KEY")
	ddEP := DatadogHostUS5
	options := &Options{
		ApiKey:          &apiKey,
		Service:         &service,
		Host:            &service,
		DatadogEndpoint: &ddEP,
	}

	hook, err := New(options)
	if err != nil {
		panic(err.Error())
	}

	session := uuid.New().String()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.AddHook(hook)
	logrus.DeferExitHandler(hook.Close)

	logger = logrus.WithFields(logrus.Fields{
		"hostname": os.Getenv("ENV"),
		"service":  service,
		"session":  session,
	})

}
