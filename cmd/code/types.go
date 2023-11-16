package code

import "github.com/qdriven/qfluent-go/pkg/utils/shell"

var actions = shell.LoadCommands("starters.json")

type StarterSet struct {
	Starters []Command `yaml:"starters" mapstructure:"starters" json:"starters"`
}
type Command struct {
	Name     string   `yaml:"name" mapstructure:"name" json:"name"`
	Desc     string   `yaml:"desc" mapstructure:"desc" json:"desc"`
	Commands []string `yaml:"commands" mapstructure:"commands" json:"commands"`
}
