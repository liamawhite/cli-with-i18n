package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
)

type RunningSecurityGroupsCommand struct {
	usage           interface{} `usage:"CF_NAME running-security-groups"`
	relatedCommands interface{} `related_commands:"bind-running-security-group, security-group, unbind-running-security-group"`
}

func (RunningSecurityGroupsCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (RunningSecurityGroupsCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
