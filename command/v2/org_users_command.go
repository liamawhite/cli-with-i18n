package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type OrgUsersCommand struct {
	RequiredArgs    flag.Organization `positional-args:"yes"`
	AllUsers        bool              `short:"a" description:"List all users in the org"`
	usage           interface{}       `usage:"CF_NAME org-users ORG"`
	relatedCommands interface{}       `related_commands:"orgs"`
}

func (OrgUsersCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (OrgUsersCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
