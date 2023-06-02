package scripts

import (
	"io"
	"log"
	"os"
	"qfluent-go/internal/jsonutil"
	"qfluent-go/internal/tpl"
)

// TODO Get Pipeline Output
type NamedCommand struct {
	Name    string `json:"name "`
	Command string `json:"command"`
}

type NamedCommands struct {
	Commands []NamedCommand `json:"commands"`
}

type CommandRunner interface {
	Run(commandsFile string)
}

type CommandsContext struct {
	Commands     NamedCommands
	errorChecker ErrorChecker
	RunningLogs  map[string]string
	InitData     any
}

func ExecShellCommand(cmd string) (int, error) {
	return Exec(cmd).
		FilterScan(func(s string, writer io.Writer) {
			log.Println(s)
		}).Stdout()
}

func ExecShellCommands(jsonFile string, data any) error {
	jsonBytes, _ := os.ReadFile(jsonFile)
	var commands = NamedCommands{}
	jsonutil.JsonConverter.ToObject(string(jsonBytes), &commands)
	return ExecCommands(commands, data)
}

func ExecCommands(commands NamedCommands, data any) error {
	for _, namedCommand := range commands.Commands {
		log.Printf("start execute command: %s,%s", namedCommand.Name, namedCommand.Command)
		output := tpl.RenderTemplate(namedCommand.Command, data)
		_, err := ExecShellCommand(output)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

func (c CommandsContext) Run(commandFilePath string, initData any) {
	var commands = NamedCommands{}
	jsonutil.JsonConverter.FileContentToObject(commandFilePath, &commands)

	for _, namedCommand := range commands.Commands {
		log.Printf("start execute command: %s,%s", namedCommand.Name, namedCommand.Command)
		realCommand := tpl.RenderTemplate(namedCommand.Command, initData)
		data, err := ExecShellCommand(realCommand)
		if err != nil {
			log.Fatal(err)
			c.RunningLogs[namedCommand.Name] = err.Error()
		}
		c.RunningLogs[namedCommand.Name] = string(data)
	}
}
