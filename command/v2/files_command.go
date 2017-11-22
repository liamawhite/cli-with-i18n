package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type FilesCommand struct {
	RequiredArgs    flag.FilesArgs `positional-args:"yes"`
	Instance        int            `short:"i" description:"Instance"`
	usage           interface{}    `usage:"CF_NAME files APP_NAME [PATH] [-i INSTANCE]\n\nTIP:\n   To list and inspect files of an app running on the Diego backend, use 'CF_NAME ssh'"`
	relatedCommands interface{}    `related_commands:"ssh"`
}

func (FilesCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (FilesCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
