package main

import (
	"github.com/qdriven/qfluent-go/cmd/base"
	"github.com/qdriven/qfluent-go/cmd/code"
)

func init() {
	base.CmdRoot.AddCommand(code.StarterCmd)
}
func main() {
	_ = base.CmdRoot.Execute()
}
