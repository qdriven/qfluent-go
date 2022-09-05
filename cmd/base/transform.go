package cmd

import (
	"github.com/qdriven/qfluent-go/cmd"
	"github.com/qdriven/qfluent-go/pkg/generator"
	"github.com/spf13/cobra"
	"log"
)

// CLI flags
var (
	transformationsFile *string
	source              *string
	destination         *string
)

// transformCmd represents the transform command
var transformCmd = &cobra.Command{
	Use:   "transform",
	Short: "Transform a blueprint to a live project",
	Run: func(cmd *cobra.Command, args []string) {
		err := generator.Generate(*transformationsFile, *source, *destination, args)
		if err != nil {
			log.Fatalf("error generating: %s", err)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(transformCmd)
	transformationsFile = addRequiredStringFlag(transformCmd, "transformations", "",
		"Location of your transformations.yaml file")
	source = addRequiredStringFlag(transformCmd, "source", ".",
		"Location of the source (blueprint) files")
	destination = addRequiredStringFlag(transformCmd, "destination", "",
		"Location of the destination (generated) files")
}

func addRequiredStringFlag(command *cobra.Command, name, value, usage string) *string {
	ref := command.Flags().String(name, value, usage)
	err := command.MarkFlagRequired(name)
	if err != nil {
		panic(err)
	}
	return ref
}
