package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type DeleteSecurityGroupCommand struct {
	RequiredArgs    flag.SecurityGroup `positional-args:"yes"`
	Force           bool               `short:"f" description:"Force deletion without confirmation"`
	usage           interface{}        `usage:"CF_NAME delete-security-group SECURITY_GROUP [-f]"`
	relatedCommands interface{}        `related_commands:"security-groups"`
}

func (DeleteSecurityGroupCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (DeleteSecurityGroupCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
