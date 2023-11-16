package code

import (
	"github.com/qdriven/qfluent-go/pkg/utils/shell"
	"github.com/spf13/cobra"
	"strings"
)

var actions = shell.LoadCommands("starters.json")

var (
	AvailableStarter = actions.GetAvailableActionNames()
	StarterCmd       = &cobra.Command{
		Use:       "starter",
		Short:     "fluent start " + strings.Join(AvailableStarter, ","),
		Long:      "fluent start " + strings.Join(AvailableStarter, ","),
		Example:   "fluent start " + AvailableStarter[0],
		ValidArgs: AvailableStarter,
		Run: func(cmd *cobra.Command, args []string) {
			category := args[0]
			actions.ExecuteActionByName(category)
		},
	}
)
