package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type QuotaCommand struct {
	RequiredArgs    flag.Quota  `positional-args:"yes"`
	usage           interface{} `usage:"CF_NAME quota QUOTA"`
	relatedCommands interface{} `related_commands:"org, quotas"`
}

func (QuotaCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (QuotaCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
