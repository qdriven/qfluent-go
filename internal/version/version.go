package version

import (
	"fmt"
	"runtime"
)

var (
	Version  = "0.1"
	Revision = "0.1"
)

func GetVersion() string {
	return fmt.Sprintf(`Version:%s
Revision: %s
OS:%s
Arch:%s`, Version, Revision, runtime.GOOS, runtime.GOARCH)
}
