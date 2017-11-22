package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type UpdateServiceBrokerCommand struct {
	RequiredArgs    flag.ServiceBrokerArgs `positional-args:"yes"`
	usage           interface{}            `usage:"CF_NAME update-service-broker SERVICE_BROKER USERNAME PASSWORD URL"`
	relatedCommands interface{}            `related_commands:"rename-service-broker, service-brokers"`
}

func (UpdateServiceBrokerCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (UpdateServiceBrokerCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
