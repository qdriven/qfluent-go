package base

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	Revision = "dev"
	Version  = "dev"
)

func getVersion() string {
	return fmt.Sprintf(`Version: %s
				Revision: %s
				OS: %s
				Arch: %s`, Version, Revision, runtime.GOOS, runtime.GOARCH)
}

var VersionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getVersion())
	},
	Short: "Show version info",
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}
