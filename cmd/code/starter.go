package code

import (
	"github.com/spf13/cobra"
)

var (
	AvailableStarter = actions.GetAvailableActionNames()
	StarterCmd       = &cobra.Command{
		Use:       "starter",
		Short:     "fluent start [python,go,java]",
		Long:      "fluent start [python,go,java]",
		Example:   "fluent start python",
		ValidArgs: AvailableStarter,
		Run: func(cmd *cobra.Command, args []string) {
			category := args[0]
			actions.GetNamedActionByName(category)
		},
	}
)
