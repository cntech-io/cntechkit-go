package cntechkitgo

import (
	"github.com/sirupsen/logrus"
)

type logger struct {
	Logger *logrus.Logger
	config *LoggerConfig
}

type LoggerConfig struct {
	AppName string
}

func NewLogger(config *LoggerConfig) *logger {
	logrusLogger := logrus.New()
	env := NewServerEnv()
	if env.DebugModeFlag {
		logrusLogger.SetLevel(logrus.DebugLevel)
	} else {
		logrusLogger.SetLevel(logrus.WarnLevel)
	}
	logrusLogger.SetFormatter(&logrus.JSONFormatter{})
	return &logger{
		Logger: logrusLogger,
		config: config,
	}
}

func (l *logger) Warn(msg string) {
	l.Logger.WithFields(logrus.Fields{
		"app_name": l.config.AppName,
	}).Warn(msg)
}

func (l *logger) Info(msg string) {
	l.Logger.WithFields(logrus.Fields{
		"app_name": l.config.AppName,
	}).Info(msg)
}

func (l *logger) Error(msg string) {
	l.Logger.WithFields(logrus.Fields{
		"app_name": l.config.AppName,
	}).Error(msg)
}
