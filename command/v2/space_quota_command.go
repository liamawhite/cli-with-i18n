package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type SpaceQuotaCommand struct {
	RequiredArgs flag.SpaceQuota `positional-args:"yes"`
	usage        interface{}     `usage:"CF_NAME space-quota SPACE_QUOTA_NAME"`
}

func (SpaceQuotaCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (SpaceQuotaCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
