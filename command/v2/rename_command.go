package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type RenameCommand struct {
	RequiredArgs    flag.AppRenameArgs `positional-args:"yes"`
	usage           interface{}        `usage:"CF_NAME rename APP_NAME NEW_APP_NAME"`
	relatedCommands interface{}        `related_commands:"apps, delete"`
}

func (RenameCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (RenameCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
