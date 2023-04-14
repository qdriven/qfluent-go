package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfig(t *testing.T) {
	config := NewDefaultConfig()
	fmt.Println(config)
	assert.Equal(t, config.Get("free_form"), "test")
	assert.Equal(t, config.Get("ws.sh_home"), "/bin/sh")
}

func TestNewNamedConfig(t *testing.T) {
	config := &AppConfig{}
	NewNamedConfig("app.toml", "named", config)
	fmt.Println(config)
	assert.Equal(t, config.Ws.ShHome, "/bin/sh")
	v := GetViperByName("named")
	//ws := v.Get("ws").(map[string]interface{})
	ws := v.Get("ws").(map[string]interface{})
	assert.Equal(t, ws["sh_home"], "/bin/sh")

}
