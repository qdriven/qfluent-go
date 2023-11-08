package config

import (
	"fmt"
	"testing"
)

// `mapstructure:",squash"` is required for embedded structs
type testConfig struct {
	BaseConfig `mapstructure:"core"`
	Name       string
}

type AnotherConfig struct {
	Desc string
	Misc string
	Name string
}

var testConfigInstance = testConfig{
	BaseConfig: BaseConfig{
		DB: DBConfig{
			Driver: "sqlite",
			DSN:    "./test.db",
		},
		HTTP: HTTPConfig{
			Address: ":8080",
		},
		LogLevel: "info",
	},
	Name: "Test Config",
}

var AnotherConfigInstance = &AnotherConfig{
	Desc: "desc",
	Misc: DefaultConfigFile,
	Name: ConfigEvnPrefix,
}

// Test New App Configuration
func TestNewAppConfig(t *testing.T) {
	appconfig, _ := DefaultAppConfig()
	_ = appconfig.AddConfig(testConfigInstance)
	//_ = appconfig.AddConfig(AnotherConfigInstance)
	appconfig.WriteConfig("config.yaml")
}

func TestWriteConfig(t *testing.T) {
	appconfig, _ := NewAppConfig(nil)
	_ = appconfig.AddConfig(testConfigInstance)
	//_ = appconfig.AddConfig(AnotherConfigInstance)
	appconfig.WriteConfig("config.yaml")
}

func TestReadConfig(t *testing.T) {
	appconfig, _ := NewAppConfig("config.yaml")
	config := &testConfig{}
	fmt.Println(appconfig.Viper.Get("db"))
	appconfig.ToStruct(config)

	fmt.Println(config)
}
