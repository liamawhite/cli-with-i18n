package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
)

type RouterGroupsCommand struct {
	usage           interface{} `usage:"CF_NAME router-groups"`
	relatedCommands interface{} `related_commands:"create-domain, domains"`
}

func (RouterGroupsCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (RouterGroupsCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
