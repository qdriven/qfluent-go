package code

import (
	"fmt"
	"github.com/qdriven/qfluent-go/pkg/utils/shell"
	"os"
	"path"
	"testing"
)

func TestCreateStartProject(t *testing.T) {
	var CurrentDir, _ = os.Getwd()
	var actions = shell.LoadCommands(path.Join(CurrentDir, "starters.json"))
	fmt.Println(actions)
}
