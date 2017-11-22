package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type SetStagingEnvironmentVariableGroupCommand struct {
	RequiredArgs    flag.ParamsAsJSON `positional-args:"yes"`
	usage           interface{}       `usage:"CF_NAME set-staging-environment-variable-group '{\"name\":\"value\",\"name\":\"value\"}'"`
	relatedCommands interface{}       `related_commands:"set-env, staging-environment-variable-group"`
}

func (SetStagingEnvironmentVariableGroupCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (SetStagingEnvironmentVariableGroupCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
