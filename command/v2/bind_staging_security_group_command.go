package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type BindStagingSecurityGroupCommand struct {
	RequiredArgs    flag.SecurityGroup `positional-args:"yes"`
	usage           interface{}        `usage:"CF_NAME bind-staging-security-group SECURITY_GROUP"`
	relatedCommands interface{}        `related_commands:"apps, bind-running-security-group, bind-security-group, restart, security-groups, staging-security-groups"`
}

func (BindStagingSecurityGroupCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (BindStagingSecurityGroupCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
