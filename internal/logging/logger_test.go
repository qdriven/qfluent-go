package log

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger(DefaultLoggerOptions()...)
	logger.Infof("testing infor")
}

func TestNewZoneLogger(t *testing.T) {
	logger := ZoneLogger("log-test")
	logger.Infof("testing infor")
}

func TestDefaultLogger(t *testing.T) {
	Logger.Warn("test")
	Logger.Info("test")
	Logger.Error("test")

}
