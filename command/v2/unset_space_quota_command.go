package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type UnsetSpaceQuotaCommand struct {
	RequiredArgs    flag.SetSpaceQuotaArgs `positional-args:"yes"`
	usage           interface{}            `usage:"CF_NAME unset-space-quota SPACE SPACE_QUOTA"`
	relatedCommands interface{}            `related_commands:"space"`
}

func (UnsetSpaceQuotaCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (UnsetSpaceQuotaCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
