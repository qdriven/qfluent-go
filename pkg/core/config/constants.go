package config

import "errors"

const (
	ConfigEnv         = "GVA_CONFIG"
	ConfigDefaultFile = "config.yaml"
	ConfigTestFile    = "config.test.yaml"
	ConfigDebugFile   = "config.debug.yaml"
	ConfigReleaseFile = "config.release.yaml"
)

var ErrConfigNotPtrToStruct = errors.New("config must be a pointer to struct")
