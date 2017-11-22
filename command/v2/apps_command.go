package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
)

type AppsCommand struct {
	usage           interface{} `usage:"CF_NAME apps"`
	relatedCommands interface{} `related_commands:"events, logs, map-route, push, scale, start, stop, restart"`
}

func (AppsCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (AppsCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
