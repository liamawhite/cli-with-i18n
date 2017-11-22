package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type CreateServiceAuthTokenCommand struct {
	RequiredArgs flag.ServiceAuthTokenArgs `positional-args:"yes"`
	usage        interface{}               `usage:"CF_NAME create-service-auth-token LABEL PROVIDER TOKEN"`
}

func (CreateServiceAuthTokenCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (CreateServiceAuthTokenCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
