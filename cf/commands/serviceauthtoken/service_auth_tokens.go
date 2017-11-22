package serviceauthtoken

import (
	"github.com/liamawhite/cli-with-i18n/cf"
	"github.com/liamawhite/cli-with-i18n/cf/api"
	"github.com/liamawhite/cli-with-i18n/cf/commandregistry"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/flags"
	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"
	"github.com/liamawhite/cli-with-i18n/cf/terminal"
)

type ListServiceAuthTokens struct {
	ui            terminal.UI
	config        coreconfig.Reader
	authTokenRepo api.ServiceAuthTokenRepository
}

func init() {
	commandregistry.Register(&ListServiceAuthTokens{})
}

func (cmd *ListServiceAuthTokens) MetaData() commandregistry.CommandMetadata {
	return commandregistry.CommandMetadata{
		Name:        "service-auth-tokens",
		Description: T("List service auth tokens"),
		Usage: []string{
			T("CF_NAME service-auth-tokens"),
		},
	}
}

func (cmd *ListServiceAuthTokens) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	usageReq := requirements.NewUsageRequirement(commandregistry.CLICommandUsagePresenter(cmd),
		T("No argument required"),
		func() bool {
			return len(fc.Args()) != 0
		},
	)

	reqs := []requirements.Requirement{
		usageReq,
		requirementsFactory.NewLoginRequirement(),
		requirementsFactory.NewMaxAPIVersionRequirement(
			"service-auth-tokens",
			cf.ServiceAuthTokenMaximumAPIVersion,
		),
	}

	return reqs, nil
}

func (cmd *ListServiceAuthTokens) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	cmd.ui = deps.UI
	cmd.config = deps.Config
	cmd.authTokenRepo = deps.RepoLocator.GetServiceAuthTokenRepository()
	return cmd
}

func (cmd *ListServiceAuthTokens) Execute(c flags.FlagContext) error {
	cmd.ui.Say(T("Getting service auth tokens as {{.CurrentUser}}...",
		map[string]interface{}{
			"CurrentUser": terminal.EntityNameColor(cmd.config.Username()),
		}))
	authTokens, err := cmd.authTokenRepo.FindAll()
	if err != nil {
		return err
	}
	cmd.ui.Ok()
	cmd.ui.Say("")

	table := cmd.ui.Table([]string{T("label"), T("provider")})

	for _, authToken := range authTokens {
		table.Add(authToken.Label, authToken.Provider)
	}

	err = table.Print()
	if err != nil {
		return err
	}
	return nil
}
