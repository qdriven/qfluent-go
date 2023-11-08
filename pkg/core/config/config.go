package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/qdriven/qfluent-go/pkg/utils"
	"github.com/spf13/viper"
	"log/slog"
	"reflect"
)

type AppConfig struct {
	Viper           *viper.Viper
	ConfigFile      string
	SavedConfigFile string
}

func NewAppConfig(filePath any) (*AppConfig, error) {
	appConfig := &AppConfig{
		Viper:           viper.New(),
		SavedConfigFile: "config-backup",
	}
	if filePath != nil && reflect.TypeOf(filePath).Kind() == reflect.String {
		appConfig.Viper.SetConfigType(utils.GetFileExtension(filePath.(string)))
		appConfig.Viper.SetConfigFile(filePath.(string))
		err := appConfig.Viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	}

	return appConfig, nil
}

var DefaultAppConfig, _ = NewAppConfig(DefaultConfigFile)

func (a *AppConfig) ToStruct(toStruct any) {
	err := a.Viper.Unmarshal(toStruct)
	if err != nil {
		panic(err)
	}
}

func (a *AppConfig) AddConfig(config any) error {
	configJson, err := json.Marshal(config)
	if err != nil {
		return err
	}

	a.Viper.SetConfigType("json")
	err = a.Viper.ReadConfig(bytes.NewBuffer(configJson))
	return err
}

func (a *AppConfig) WriteConfig(filePath string) {
	err := a.Viper.SafeWriteConfigAs(filePath)
	if err != nil {
		slog.Error("write configuration failed", err)
		_ = a.Viper.SafeWriteConfigAs(a.SavedConfigFile)
		return
	}
}

// Watched Changed Overtime to update, it is dangerous operations
func (a *AppConfig) WatchConfigChanges() {
	a.Viper.WatchConfig()
	a.Viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := a.Viper.ReadInConfig(); err != nil {
			fmt.Printf("couldn't load config: %s", err)
		}
	})
}
