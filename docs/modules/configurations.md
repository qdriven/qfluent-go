# Configuration

- Read Configuration From File
- Read Configuration From Model
- Write Configuration To File
- Write Configuration To Model

## Usage

```go

// Test Default Configuration
func TestReadConfigurationFromModelAndWriteToFile(t *testing.T) {
	appconfig := DefaultAppConfig
	_ = appconfig.AddConfig(testConfigInstance)
	//_ = appconfig.AddConfig(AnotherConfigInstance)
	appconfig.WriteConfig("config.yaml")
}

func TestReadConfigurationFromFileAndWriteToStruct(t *testing.T) {
	appconfig, _ := NewAppConfig("config.yaml")
	appconfig.ToStruct(AnotherConfigInstance)
	assert.Equal(t, "FLUENT", AnotherConfigInstance.Name)
	assert.Equal(t, DefaultConfigFile, AnotherConfigInstance.Misc)
}

func TestReadConfigurationFromFileAndWriteToFile(t *testing.T) {
	appconfig, _ := NewAppConfig("config.yaml")
	config := &testConfig{}
	fmt.Println(appconfig.Viper.Get("db"))
	appconfig.ToStruct(config)

	fmt.Println(config)
}

```