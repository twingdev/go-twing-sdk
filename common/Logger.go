package common

import "go.uber.org/zap"

type Log *zap.Logger

var logger Log

func init() {
	logger, _ = zap.NewProduction()
}

func GetLogger() *zap.Logger {
	return logger
}
