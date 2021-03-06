package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type DeleteServiceAuthTokenCommand struct {
	RequiredArgs flag.DeleteServiceAuthTokenArgs `positional-args:"yes"`
	Force        bool                            `short:"f" description:"Force deletion without confirmation"`
	usage        interface{}                     `usage:"CF_NAME delete-service-auth-token LABEL PROVIDER [-f]"`
}

func (DeleteServiceAuthTokenCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (DeleteServiceAuthTokenCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
