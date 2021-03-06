package v2

import (
	"github.com/liamawhite/cli-with-i18n/actor/sharedaction"
	"github.com/liamawhite/cli-with-i18n/actor/v2action"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/v2/shared"
)

//go:generate counterfeiter . OauthTokenActor

type OauthTokenActor interface {
	RefreshAccessToken(refreshToken string) (string, error)
}

type OauthTokenCommand struct {
	usage           interface{} `usage:"CF_NAME oauth-token"`
	relatedCommands interface{} `related_commands:"curl"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       OauthTokenActor
}

func (cmd *OauthTokenCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config
	cmd.SharedActor = sharedaction.NewActor(config)

	ccClient, uaaClient, err := shared.NewClients(config, ui, true)
	if err != nil {
		return err
	}
	cmd.Actor = v2action.NewActor(ccClient, uaaClient, config)

	return nil
}

func (cmd OauthTokenCommand) Execute(_ []string) error {
	err := cmd.SharedActor.CheckTarget(cmd.Config, false, false)
	if err != nil {
		return shared.HandleError(err)
	}

	accessToken, err := cmd.Actor.RefreshAccessToken(cmd.Config.RefreshToken())
	if err != nil {
		return shared.HandleError(err)
	}

	cmd.UI.DisplayText(accessToken)
	return nil
}
