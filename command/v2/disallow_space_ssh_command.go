package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type DisallowSpaceSSHCommand struct {
	RequiredArgs    flag.Space  `positional-args:"yes"`
	usage           interface{} `usage:"CF_NAME disallow-space-ssh SPACE_NAME"`
	relatedCommands interface{} `related_commands:"disable-ssh, space-ssh-allowed, ssh, ssh-enabled"`
}

func (DisallowSpaceSSHCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (DisallowSpaceSSHCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
