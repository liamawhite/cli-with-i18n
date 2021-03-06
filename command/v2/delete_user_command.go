package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type DeleteUserCommand struct {
	RequiredArgs    flag.Username `positional-args:"yes"`
	Force           bool          `short:"f" description:"Force deletion without confirmation"`
	usage           interface{}   `usage:"CF_NAME delete-user USERNAME [-f]"`
	relatedCommands interface{}   `related_commands:"org-users"`
}

func (DeleteUserCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (DeleteUserCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
