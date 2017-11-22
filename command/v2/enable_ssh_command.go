package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type EnableSSHCommand struct {
	RequiredArgs    flag.AppName `positional-args:"yes"`
	usage           interface{}  `usage:"CF_NAME enable-ssh APP_NAME"`
	relatedCommands interface{}  `related_commands:"allow-space-ssh, space-ssh-allowed, ssh, ssh-enabled"`
}

func (EnableSSHCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (EnableSSHCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
