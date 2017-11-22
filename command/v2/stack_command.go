package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type StackCommand struct {
	RequiredArgs    flag.StackName `positional-args:"yes"`
	GUID            bool           `long:"guid" description:"Retrieve and display the given stack's guid. All other output for the stack is suppressed."`
	usage           interface{}    `usage:"CF_NAME stack STACK_NAME"`
	relatedCommands interface{}    `related_commands:"app, push, stacks"`
}

func (StackCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (StackCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
