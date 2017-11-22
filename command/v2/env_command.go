package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type EnvCommand struct {
	RequiredArgs    flag.AppName `positional-args:"yes"`
	usage           interface{}  `usage:"CF_NAME env APP_NAME"`
	relatedCommands interface{}  `related_commands:"app, apps, set-env, unset-env, running-environment-variable-group, staging-environment-variable-group"`
}

func (EnvCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (EnvCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
