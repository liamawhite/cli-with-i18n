package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type SpaceUsersCommand struct {
	RequiredArgs    flag.OrgSpace `positional-args:"yes"`
	usage           interface{}   `usage:"CF_NAME space-users ORG SPACE"`
	relatedCommands interface{}   `related_commands:"org-users, set-space-role, unset-space-role, orgs, spaces"`
}

func (SpaceUsersCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (SpaceUsersCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
