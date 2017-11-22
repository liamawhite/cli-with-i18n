package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
)

type LogoutCommand struct {
	usage interface{} `usage:"CF_NAME logout"`
}

func (LogoutCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (LogoutCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
