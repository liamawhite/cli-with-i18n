package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type RenameBuildpackCommand struct {
	RequiredArgs    flag.RenameBuildpackArgs `positional-args:"yes"`
	usage           interface{}              `usage:"CF_NAME rename-buildpack BUILDPACK_NAME NEW_BUILDPACK_NAME"`
	relatedCommands interface{}              `related_commands:"update-buildpack"`
}

func (RenameBuildpackCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (RenameBuildpackCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
