package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
)

type RunningEnvironmentVariableGroupCommand struct {
	usage           interface{} `usage:"CF_NAME running-environment-variable-group"`
	relatedCommands interface{} `related_commands:"env, staging-environment-variable-group"`
}

func (RunningEnvironmentVariableGroupCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (RunningEnvironmentVariableGroupCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
