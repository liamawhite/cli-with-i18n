package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type RenameSpaceCommand struct {
	RequiredArgs flag.RenameSpaceArgs `positional-args:"yes"`
	usage        interface{}          `usage:"CF_NAME rename-space SPACE NEW_SPACE"`
}

func (RenameSpaceCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (RenameSpaceCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
