package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type CreateOrgCommand struct {
	RequiredArgs    flag.Organization `positional-args:"yes"`
	Quota           string            `short:"q" description:"Quota to assign to the newly created org (excluding this option results in assignment of default quota)"`
	usage           interface{}       `usage:"CF_NAME create-org ORG"`
	relatedCommands interface{}       `related_commands:"create-space, orgs, quotas, set-org-role"`
}

func (CreateOrgCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (CreateOrgCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
