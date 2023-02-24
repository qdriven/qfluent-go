package gosh

import (
	"fmt"
	"testing"
)

func TestExecutor(t *testing.T) {
	result := ExecCommand("/Users/patrick/workspace/personal/qdriven/daily-qa/daily-go/qfluent-go/pkg/gosh/git_commands.sh")
	fmt.Println(result)
}
