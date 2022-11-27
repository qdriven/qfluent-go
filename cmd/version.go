package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"qfluent-go/internal"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(internal.GetVersion())
	},
	Short: "Show Version Info",
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
