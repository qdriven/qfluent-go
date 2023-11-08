package shell

import (
	"github.com/qdriven/qfluent-go/pkg/utils"
	"github.com/qdriven/qfluent-go/pkg/utils/jsonutil"
	"io"
	"log"
	"log/slog"
	"os"
)

type NamedCommand struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	Help    string `json: "help"`
}
type Commands struct {
	Commands []NamedCommand `json:"commands"`
}

func ExecShellCommand(cmd string) (int, error) {
	return Exec(cmd).
		FilterScan(func(s string, writer io.Writer) {
			log.Println(s)
		}).Stdout()
}

func ExecShellCommands(jsonFile string, data any) error {
	jsonBytes, err := os.ReadFile(jsonFile)
	if err != nil {
		slog.Error("error when read file", err)
	}

	var commands = &Commands{}
	jsonutil.ToStruct(string(jsonBytes), &commands)
	for _, command := range commands.Commands {
		output := utils.RenderTemplate(command.Command, data)
		_, err := ExecShellCommand(output)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}
