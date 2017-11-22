package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
)

type ServicesCommand struct {
	usage           interface{} `usage:"CF_NAME services"`
	relatedCommands interface{} `related_commands:"create-service, marketplace"`
}

func (ServicesCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (ServicesCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
