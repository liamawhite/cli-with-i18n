package application

import (
	"github.com/liamawhite/cli-with-i18n/cf/api"
	"github.com/liamawhite/cli-with-i18n/cf/api/applications"
	"github.com/liamawhite/cli-with-i18n/cf/commandregistry"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/errors"
	"github.com/liamawhite/cli-with-i18n/cf/flags"
	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"
	"github.com/liamawhite/cli-with-i18n/cf/terminal"
)

type DeleteApp struct {
	ui        terminal.UI
	config    coreconfig.Reader
	appRepo   applications.Repository
	routeRepo api.RouteRepository
	appReq    requirements.ApplicationRequirement
}

func init() {
	commandregistry.Register(&DeleteApp{})
}

func (cmd *DeleteApp) MetaData() commandregistry.CommandMetadata {
	fs := make(map[string]flags.FlagSet)
	fs["f"] = &flags.BoolFlag{ShortName: "f", Usage: T("Force deletion without confirmation")}
	fs["r"] = &flags.BoolFlag{ShortName: "r", Usage: T("Also delete any mapped routes")}

	return commandregistry.CommandMetadata{
		Name:        "delete",
		ShortName:   "d",
		Description: T("Delete an app"),
		Usage: []string{
			T("CF_NAME delete APP_NAME [-f -r]"),
		},
		Flags: fs,
	}
}

func (cmd *DeleteApp) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	usageReq := requirementsFactory.NewUsageRequirement(commandregistry.CLICommandUsagePresenter(cmd),
		T("Requires app name as argument"),
		func() bool {
			return len(fc.Args()) != 1
		},
	)

	reqs := []requirements.Requirement{
		usageReq,
		requirementsFactory.NewLoginRequirement(),
		requirementsFactory.NewTargetedSpaceRequirement(),
	}

	return reqs, nil
}

func (cmd *DeleteApp) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	cmd.ui = deps.UI
	cmd.config = deps.Config
	cmd.appRepo = deps.RepoLocator.GetApplicationRepository()
	cmd.routeRepo = deps.RepoLocator.GetRouteRepository()
	return cmd
}

func (cmd *DeleteApp) Execute(c flags.FlagContext) error {
	appName := c.Args()[0]

	if !c.Bool("f") {
		response := cmd.ui.ConfirmDelete(T("app"), appName)
		if !response {
			return nil
		}
	}

	cmd.ui.Say(T("Deleting app {{.AppName}} in org {{.OrgName}} / space {{.SpaceName}} as {{.Username}}...",
		map[string]interface{}{
			"AppName":   terminal.EntityNameColor(appName),
			"OrgName":   terminal.EntityNameColor(cmd.config.OrganizationFields().Name),
			"SpaceName": terminal.EntityNameColor(cmd.config.SpaceFields().Name),
			"Username":  terminal.EntityNameColor(cmd.config.Username())}))

	app, err := cmd.appRepo.Read(appName)

	switch err.(type) {
	case nil: // no error
	case *errors.ModelNotFoundError:
		cmd.ui.Ok()
		cmd.ui.Warn(T("App {{.AppName}} does not exist.", map[string]interface{}{"AppName": appName}))
		return nil
	default:
		return err
	}

	if c.Bool("r") {
		for _, route := range app.Routes {
			err = cmd.routeRepo.Delete(route.GUID)
			if err != nil {
				return err
			}
		}
	}

	err = cmd.appRepo.Delete(app.GUID)
	if err != nil {
		return err
	}

	cmd.ui.Ok()
	return nil
}
