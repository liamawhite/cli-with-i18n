package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
)

type PasswdCommand struct {
	usage interface{} `usage:"CF_NAME passwd"`
}

func (PasswdCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (PasswdCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
