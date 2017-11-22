package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type SetQuotaCommand struct {
	RequiredArgs    flag.SetOrgQuotaArgs `positional-args:"yes"`
	usage           interface{}          `usage:"CF_NAME set-quota ORG QUOTA\n\nTIP:\n   View allowable quotas with 'CF_NAME quotas'"`
	relatedCommands interface{}          `related_commands:"orgs, quotas"`
}

func (SetQuotaCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (SetQuotaCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
