package organization

import (
	"errors"

	"github.com/liamawhite/cli-with-i18n/cf/api/organizations"
	"github.com/liamawhite/cli-with-i18n/cf/commandregistry"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/flags"
	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
	"github.com/liamawhite/cli-with-i18n/cf/models"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"
	"github.com/liamawhite/cli-with-i18n/cf/terminal"
	"github.com/liamawhite/cli-with-i18n/plugin/models"
)

const orgLimit = 0

type ListOrgs struct {
	ui              terminal.UI
	config          coreconfig.Reader
	orgRepo         organizations.OrganizationRepository
	pluginOrgsModel *[]plugin_models.GetOrgs_Model
	pluginCall      bool
}

func init() {
	commandregistry.Register(&ListOrgs{})
}

func (cmd *ListOrgs) MetaData() commandregistry.CommandMetadata {
	return commandregistry.CommandMetadata{
		Name:        "orgs",
		ShortName:   "o",
		Description: T("List all orgs"),
		Usage: []string{
			"CF_NAME orgs",
		},
	}
}

func (cmd *ListOrgs) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	usageReq := requirements.NewUsageRequirement(commandregistry.CLICommandUsagePresenter(cmd),
		T("No argument required"),
		func() bool {
			return len(fc.Args()) != 0
		},
	)

	reqs := []requirements.Requirement{
		usageReq,
		requirementsFactory.NewLoginRequirement(),
	}

	return reqs, nil
}

func (cmd *ListOrgs) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	cmd.ui = deps.UI
	cmd.config = deps.Config
	cmd.orgRepo = deps.RepoLocator.GetOrganizationRepository()
	cmd.pluginOrgsModel = deps.PluginModels.Organizations
	cmd.pluginCall = pluginCall
	return cmd
}

func (cmd ListOrgs) Execute(fc flags.FlagContext) error {
	cmd.ui.Say(T("Getting orgs as {{.Username}}...\n",
		map[string]interface{}{"Username": terminal.EntityNameColor(cmd.config.Username())}))

	noOrgs := true
	table := cmd.ui.Table([]string{T("name")})

	orgs, err := cmd.orgRepo.ListOrgs(orgLimit)
	if err != nil {
		return err
	}
	for _, org := range orgs {
		table.Add(org.Name)
		noOrgs = false
	}

	err = table.Print()
	if err != nil {
		return err
	}

	if err != nil {
		return errors.New(T("Failed fetching orgs.\n{{.APIErr}}",
			map[string]interface{}{"APIErr": err}))
	}

	if noOrgs {
		cmd.ui.Say(T("No orgs found"))
	}

	if cmd.pluginCall {
		cmd.populatePluginModel(orgs)
	}
	return nil
}

func (cmd *ListOrgs) populatePluginModel(orgs []models.Organization) {
	for _, org := range orgs {
		orgModel := plugin_models.GetOrgs_Model{}
		orgModel.Name = org.Name
		orgModel.Guid = org.GUID
		*(cmd.pluginOrgsModel) = append(*(cmd.pluginOrgsModel), orgModel)
	}
}
