package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type ServiceKeysCommand struct {
	RequiredArgs    flag.ServiceInstance `positional-args:"yes"`
	usage           interface{}          `usage:"CF_NAME service-keys SERVICE_INSTANCE\n\nEXAMPLES:\n   CF_NAME service-keys mydb"`
	relatedCommands interface{}          `related_commands:"delete-service-key"`
}

func (ServiceKeysCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (ServiceKeysCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
