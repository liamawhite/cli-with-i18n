package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type DeleteSpaceQuotaCommand struct {
	RequiredArgs    flag.SpaceQuota `positional-args:"yes"`
	Force           bool            `short:"f" description:"Force deletion without confirmation"`
	usage           interface{}     `usage:"CF_NAME delete-space-quota SPACE_QUOTA_NAME [-f]"`
	relatedCommands interface{}     `related_commands:"space-quotas"`
}

func (DeleteSpaceQuotaCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (DeleteSpaceQuotaCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
