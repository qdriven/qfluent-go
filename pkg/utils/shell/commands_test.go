package shell

import (
	"fmt"
	"testing"
)

func TestExecShellCommands(t *testing.T) {
	commands := &Commands{}
	result := ExecShellCommands("commands.json", commands)
	fmt.Println(result)
}
