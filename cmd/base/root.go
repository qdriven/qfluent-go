package base

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var RootCmd = &cobra.Command{
	Use: "spell-cli",
}

func exitError(msg interface{}) {
	_, err := fmt.Fprintln(os.Stderr, msg)
	if err != nil {
		log.Fatal(err)
		return
	}
	os.Exit(1)
}

func Execute() {
	RootCmd.Run = func(cmd *cobra.Command, args []string) {
		_ = RootCmd.Help()
	}

	if err := RootCmd.Execute(); err != nil {
		exitError(err)
	}
}
