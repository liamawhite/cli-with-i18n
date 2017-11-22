package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
)

type SpaceQuotasCommand struct {
	usage           interface{} `usage:"CF_NAME space-quotas"`
	relatedCommands interface{} `related_commands:"set-space-quota"`
}

func (SpaceQuotasCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (SpaceQuotasCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
