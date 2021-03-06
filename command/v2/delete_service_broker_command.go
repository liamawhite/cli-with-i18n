package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type DeleteServiceBrokerCommand struct {
	RequiredArgs    flag.ServiceBroker `positional-args:"yes"`
	Force           bool               `short:"f" description:"Force deletion without confirmation"`
	usage           interface{}        `usage:"CF_NAME delete-service-broker SERVICE_BROKER [-f]"`
	relatedCommands interface{}        `related_commands:"delete-service, purge-service-offering, service-brokers"`
}

func (DeleteServiceBrokerCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (DeleteServiceBrokerCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
