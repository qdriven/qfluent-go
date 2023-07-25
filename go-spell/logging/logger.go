package logging

import "go.uber.org/zap"

func Logger() *zap.Logger {
	return zap.L().Named("default")
}

func NamedLogger(name string) *zap.Logger {
	return zap.L().Named(name)
}
