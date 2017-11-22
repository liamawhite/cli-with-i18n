package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type SpaceSSHAllowedCommand struct {
	RequiredArgs    flag.Space  `positional-args:"yes"`
	usage           interface{} `usage:"CF_NAME space-ssh-allowed SPACE_NAME"`
	relatedCommands interface{} `related_commands:"allow-space-ssh, ssh-enabled, ssh"`
}

func (SpaceSSHAllowedCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (SpaceSSHAllowedCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
