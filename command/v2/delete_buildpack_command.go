package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type DeleteBuildpackCommand struct {
	RequiredArgs    flag.BuildpackName `positional-args:"yes"`
	Force           bool               `short:"f" description:"Force deletion without confirmation"`
	usage           interface{}        `usage:"CF_NAME delete-buildpack BUILDPACK [-f]"`
	relatedCommands interface{}        `related_commands:"buildpacks"`
}

func (DeleteBuildpackCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (DeleteBuildpackCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
