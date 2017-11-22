package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type RenameOrgCommand struct {
	RequiredArgs flag.RenameOrgArgs `positional-args:"yes"`
	usage        interface{}        `usage:"CF_NAME rename-org ORG NEW_ORG"`
}

func (RenameOrgCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (RenameOrgCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
