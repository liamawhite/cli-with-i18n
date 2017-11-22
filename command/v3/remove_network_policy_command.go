package v3

import (
	"net/http"

	"github.com/liamawhite/cli-with-i18n/actor/cfnetworkingaction"
	"github.com/liamawhite/cli-with-i18n/actor/sharedaction"
	"github.com/liamawhite/cli-with-i18n/actor/v3action"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccerror"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
	"github.com/liamawhite/cli-with-i18n/command/translatableerror"
	"github.com/liamawhite/cli-with-i18n/command/v3/shared"
)

//go:generate counterfeiter . RemoveNetworkPolicyActor

type RemoveNetworkPolicyActor interface {
	RemoveNetworkPolicy(spaceGUID string, srcAppName string, destAppName string, protocol string, startPort int, endPort int) (cfnetworkingaction.Warnings, error)
}

type RemoveNetworkPolicyCommand struct {
	RequiredArgs   flag.RemoveNetworkPolicyArgs `positional-args:"yes"`
	DestinationApp string                       `long:"destination-app" required:"true" description:"Name of app to connect to"`
	Port           flag.NetworkPort             `long:"port" required:"true" description:"Port or range of ports that destination app is connected with"`
	Protocol       flag.NetworkProtocol         `long:"protocol" required:"true" description:"Protocol that apps are connected with"`

	usage           interface{} `usage:"CF_NAME remove-network-policy SOURCE_APP --destination-app DESTINATION_APP --protocol (tcp | udp) --port RANGE\n\nEXAMPLES:\n   CF_NAME remove-network-policy frontend --destination-app backend --protocol tcp --port 8081\n   CF_NAME remove-network-policy frontend --destination-app backend --protocol tcp --port 8080-8090"`
	relatedCommands interface{} `related_commands:"apps, network-policies"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       RemoveNetworkPolicyActor
}

func (cmd *RemoveNetworkPolicyCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config
	cmd.SharedActor = sharedaction.NewActor(config)

	client, uaa, err := shared.NewClients(config, ui, true)
	if err != nil {
		if v3Err, ok := err.(ccerror.V3UnexpectedResponseError); ok && v3Err.ResponseCode == http.StatusNotFound {
			return translatableerror.CFNetworkingEndpointNotFoundError{}
		}

		return err
	}

	v3Actor := v3action.NewActor(nil, client, config)
	networkingClient, err := shared.NewNetworkingClient(client.NetworkPolicyV1(), config, uaa, ui)
	if err != nil {
		return err
	}
	cmd.Actor = cfnetworkingaction.NewActor(networkingClient, v3Actor)

	return nil
}

func (cmd RemoveNetworkPolicyCommand) Execute(args []string) error {
	err := cmd.SharedActor.CheckTarget(cmd.Config, true, true)
	if err != nil {
		return shared.HandleError(err)
	}

	user, err := cmd.Config.CurrentUser()
	if err != nil {
		return err
	}
	cmd.UI.DisplayTextWithFlavor("Removing network policy for app {{.SrcAppName}} in org {{.Org}} / space {{.Space}} as {{.User}}...", map[string]interface{}{
		"SrcAppName": cmd.RequiredArgs.SourceApp,
		"Org":        cmd.Config.TargetedOrganization().Name,
		"Space":      cmd.Config.TargetedSpace().Name,
		"User":       user.Name,
	})

	warnings, err := cmd.Actor.RemoveNetworkPolicy(cmd.Config.TargetedSpace().GUID, cmd.RequiredArgs.SourceApp, cmd.DestinationApp, cmd.Protocol.Protocol, cmd.Port.StartPort, cmd.Port.EndPort)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		switch err.(type) {
		case cfnetworkingaction.PolicyDoesNotExistError:
			cmd.UI.DisplayText("Policy does not exist.")
		default:
			return shared.HandleError(err)
		}
	}
	cmd.UI.DisplayOK()

	return nil
}
