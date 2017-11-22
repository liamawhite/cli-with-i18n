package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
)

type BuildpacksCommand struct {
	usage           interface{} `usage:"CF_NAME buildpacks"`
	relatedCommands interface{} `related_commands:"push"`
}

func (BuildpacksCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (BuildpacksCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
