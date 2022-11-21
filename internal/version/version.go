package version

import (
	"fmt"
	"runtime"
)

var (
	MAJOR_VERSION = "0.1"
	Revision      = "0.1"
)

func GetVersion() string {
	return fmt.Sprintf(`MAJOR_VERSION:%s
Revision: %s
OS:%s
Arch:%s`, MAJOR_VERSION, Revision, runtime.GOOS, runtime.GOARCH)
}
