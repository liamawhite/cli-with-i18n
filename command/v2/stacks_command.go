package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
)

type StacksCommand struct {
	usage           interface{} `usage:"CF_NAME stacks"`
	relatedCommands interface{} `related_commands:"app, push"`
}

func (StacksCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (StacksCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
