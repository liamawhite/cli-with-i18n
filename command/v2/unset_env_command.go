package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
)

type UnsetEnvCommand struct {
	usage           interface{} `usage:"CF_NAME unset-env APP_NAME ENV_VAR_NAME"`
	relatedCommands interface{} `related_commands:"apps, env, restart, set-staging-environment-variable-group, set-running-environment-variable-group"`
}

func (UnsetEnvCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (UnsetEnvCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
