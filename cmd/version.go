package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"qfluent-go/internal/version"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.GetVersion())
	},
	Short: "Show Version Info",
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
