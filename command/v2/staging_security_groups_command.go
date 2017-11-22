package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
)

type StagingSecurityGroupsCommand struct {
	usage           interface{} `usage:"CF_NAME staging-security-groups"`
	relatedCommands interface{} `related_commands:"bind-staging-security-group, security-group, unbind-staging-security-group"`
}

func (StagingSecurityGroupsCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (StagingSecurityGroupsCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
