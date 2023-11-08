package code

import (
	"github.com/qdriven/qfluent-go/pkg/utils/shell"
	"github.com/spf13/cobra"
)

func createPythonStartProject() {
	command := "cookiecutter https://github.com/fluent-qa/fluentqa-pytpl.git --no-input"
	_, _ = shell.ExecShellCommand(command)

}

func CreateStartProject(category string) {
	if category == "python" {
		createPythonStartProject()
	}
}

var (
	AvailableStarter = []string{
		"python",
	}
	StarterCmd = &cobra.Command{
		Use:       "starter",
		Short:     "fluent start [python,go,java]",
		Long:      "fluent start [python,go,java]",
		Example:   "fluent start python",
		ValidArgs: AvailableStarter,
		Run: func(cmd *cobra.Command, args []string) {
			category := args[0]
			CreateStartProject(category)
		},
	}
)
