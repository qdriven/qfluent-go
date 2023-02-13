package configs

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func TestInitConfig(t *testing.T) {
	type args struct {
		cfgFile string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test", args: args{cfgFile: "./goqa.toml"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitConfig(tt.args.cfgFile)
			fmt.Println(viper.Get("default"))
		})
	}
}
