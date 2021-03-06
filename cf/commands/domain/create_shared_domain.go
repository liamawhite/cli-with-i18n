package domain

import (
	"errors"
	"fmt"

	"github.com/liamawhite/cli-with-i18n/cf"
	"github.com/liamawhite/cli-with-i18n/cf/api"
	"github.com/liamawhite/cli-with-i18n/cf/commandregistry"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/flags"
	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
	"github.com/liamawhite/cli-with-i18n/cf/models"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"
	"github.com/liamawhite/cli-with-i18n/cf/terminal"
)

type CreateSharedDomain struct {
	ui             terminal.UI
	config         coreconfig.Reader
	domainRepo     api.DomainRepository
	routingAPIRepo api.RoutingAPIRepository
}

func init() {
	commandregistry.Register(&CreateSharedDomain{})
}

func (cmd *CreateSharedDomain) MetaData() commandregistry.CommandMetadata {
	fs := make(map[string]flags.FlagSet)
	fs["router-group"] = &flags.StringFlag{Name: "router-group", Usage: T("Routes for this domain will be configured only on the specified router group")}
	return commandregistry.CommandMetadata{
		Name:        "create-shared-domain",
		Description: T("Create a domain that can be used by all orgs (admin-only)"),
		Usage: []string{
			T("CF_NAME create-shared-domain DOMAIN [--router-group ROUTER_GROUP]"),
		},
		Flags: fs,
	}
}

func (cmd *CreateSharedDomain) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	if len(fc.Args()) != 1 {
		cmd.ui.Failed(T("Incorrect Usage. Requires DOMAIN as an argument\n\n") + commandregistry.Commands.CommandUsage("create-shared-domain"))
		return nil, fmt.Errorf("Incorrect usage: %d arguments of %d required", len(fc.Args()), 1)
	}

	reqs := []requirements.Requirement{
		requirementsFactory.NewLoginRequirement(),
	}

	if fc.String("router-group") != "" {
		reqs = append(reqs, []requirements.Requirement{
			requirementsFactory.NewMinAPIVersionRequirement("Option '--router-group'", cf.RoutePathMinimumAPIVersion),
			requirementsFactory.NewRoutingAPIRequirement(),
		}...)
	}

	return reqs, nil
}

func (cmd *CreateSharedDomain) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	cmd.ui = deps.UI
	cmd.config = deps.Config
	cmd.domainRepo = deps.RepoLocator.GetDomainRepository()
	cmd.routingAPIRepo = deps.RepoLocator.GetRoutingAPIRepository()
	return cmd
}

func (cmd *CreateSharedDomain) Execute(c flags.FlagContext) error {
	var routerGroup models.RouterGroup
	domainName := c.Args()[0]
	routerGroupName := c.String("router-group")

	if routerGroupName != "" {
		var routerGroupFound bool
		err := cmd.routingAPIRepo.ListRouterGroups(func(group models.RouterGroup) bool {
			if group.Name == routerGroupName {
				routerGroup = group
				routerGroupFound = true
				return false
			}

			return true
		})

		if err != nil {
			return err
		}
		if !routerGroupFound {
			return errors.New(T("Router group {{.RouterGroup}} not found", map[string]interface{}{
				"RouterGroup": routerGroupName,
			}))
		}
	}

	cmd.ui.Say(T("Creating shared domain {{.DomainName}} as {{.Username}}...",
		map[string]interface{}{
			"DomainName": terminal.EntityNameColor(domainName),
			"Username":   terminal.EntityNameColor(cmd.config.Username())}))

	err := cmd.domainRepo.CreateSharedDomain(domainName, routerGroup.GUID)
	if err != nil {
		return err
	}

	cmd.ui.Ok()
	return nil
}
