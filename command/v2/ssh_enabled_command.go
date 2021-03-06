package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type SSHEnabledCommand struct {
	RequiredArgs    flag.AppName `positional-args:"yes"`
	usage           interface{}  `usage:"CF_NAME ssh-enabled APP_NAME"`
	relatedCommands interface{}  `related_commands:"enable-ssh, space-ssh-allowed, ssh"`
}

func (SSHEnabledCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (SSHEnabledCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
