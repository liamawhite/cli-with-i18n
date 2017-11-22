package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
)

type ServiceBrokersCommand struct {
	usage           interface{} `usage:"CF_NAME service-brokers"`
	relatedCommands interface{} `related_commands:"delete-service-broker, disable-service-access, enable-service-access"`
}

func (ServiceBrokersCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (ServiceBrokersCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
