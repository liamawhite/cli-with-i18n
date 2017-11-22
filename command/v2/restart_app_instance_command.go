package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type RestartAppInstanceCommand struct {
	RequiredArgs    flag.AppInstance `positional-args:"yes"`
	usage           interface{}      `usage:"CF_NAME restart-app-instance APP_NAME INDEX"`
	relatedCommands interface{}      `related_commands:"restart"`
}

func (RestartAppInstanceCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (RestartAppInstanceCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
