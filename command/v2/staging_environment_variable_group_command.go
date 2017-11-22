package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
)

type StagingEnvironmentVariableGroupCommand struct {
	usage           interface{} `usage:"CF_NAME staging-environment-variable-group"`
	relatedCommands interface{} `related_commands:"env, running-environment-variable-group"`
}

func (StagingEnvironmentVariableGroupCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (StagingEnvironmentVariableGroupCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
