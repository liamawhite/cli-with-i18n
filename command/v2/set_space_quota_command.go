package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type SetSpaceQuotaCommand struct {
	RequiredArgs    flag.SetSpaceQuotaArgs `positional-args:"yes"`
	usage           interface{}            `usage:"CF_NAME set-space-quota SPACE_NAME SPACE_QUOTA_NAME"`
	relatedCommands interface{}            `related_commands:"space, space-quotas, spaces"`
}

func (SetSpaceQuotaCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (SetSpaceQuotaCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
