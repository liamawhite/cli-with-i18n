package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type DisableSSHCommand struct {
	RequiredArgs    flag.AppName `positional-args:"yes"`
	usage           interface{}  `usage:"CF_NAME disable-ssh APP_NAME"`
	relatedCommands interface{}  `related_commands:"disallow-space-ssh, space-ssh-allowed, ssh, ssh-enabled"`
}

func (DisableSSHCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (DisableSSHCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
