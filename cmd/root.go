package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"qfluent-go/config"
)

var RootCmd = &cobra.Command{
	Use:   "goqa",
	Short: "fluentqa-go is a qa daily cli application",
	Long:  `A QA tools for daily testing usage`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		println(args)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type CliBaseInfo struct {
	CfgFile     string `default: goqa.toml`
	ProjectBase string `default: ./`
	Author      string `default: Patrick`
	License     string `default: NDV`
	UseViper    bool   `default: true`
}

var appInfo = &CliBaseInfo{}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&appInfo.CfgFile, "config", "", "config file (default is ./goqa.yaml)")
	RootCmd.PersistentFlags().StringVarP(&appInfo.ProjectBase, "workspace", "b", "", "base project directory eg. github.com/spf13/")
	RootCmd.PersistentFlags().StringP("author", "a", appInfo.Author, "Author name for copyright attribution")
	RootCmd.PersistentFlags().StringVarP(&appInfo.License, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	RootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	_ = viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
	_ = viper.BindPFlag("workspace", RootCmd.PersistentFlags().Lookup("workspace"))
	_ = viper.BindPFlag("useViper", RootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", appInfo.Author)
	viper.SetDefault("license", appInfo.License)
}

func initConfig() {
	config.InitConfig(appInfo.CfgFile)
}
