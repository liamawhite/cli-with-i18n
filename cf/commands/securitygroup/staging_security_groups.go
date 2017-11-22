package securitygroup

import (
	"github.com/liamawhite/cli-with-i18n/cf/api/securitygroups/defaults/staging"
	"github.com/liamawhite/cli-with-i18n/cf/commandregistry"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/flags"
	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"
	"github.com/liamawhite/cli-with-i18n/cf/terminal"
)

type listStagingSecurityGroups struct {
	ui                       terminal.UI
	stagingSecurityGroupRepo staging.SecurityGroupsRepo
	configRepo               coreconfig.Reader
}

func init() {
	commandregistry.Register(&listStagingSecurityGroups{})
}

func (cmd *listStagingSecurityGroups) MetaData() commandregistry.CommandMetadata {
	return commandregistry.CommandMetadata{
		Name:        "staging-security-groups",
		Description: T("List security groups in the staging set for applications"),
		Usage: []string{
			"CF_NAME staging-security-groups",
		},
	}
}

func (cmd *listStagingSecurityGroups) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
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

func (cmd *listStagingSecurityGroups) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	cmd.ui = deps.UI
	cmd.configRepo = deps.Config
	cmd.stagingSecurityGroupRepo = deps.RepoLocator.GetStagingSecurityGroupsRepository()
	return cmd
}

func (cmd *listStagingSecurityGroups) Execute(context flags.FlagContext) error {
	cmd.ui.Say(T("Acquiring staging security group as {{.username}}",
		map[string]interface{}{
			"username": terminal.EntityNameColor(cmd.configRepo.Username()),
		}))

	SecurityGroupsFields, err := cmd.stagingSecurityGroupRepo.List()
	if err != nil {
		return err
	}

	cmd.ui.Ok()
	cmd.ui.Say("")

	if len(SecurityGroupsFields) > 0 {
		for _, value := range SecurityGroupsFields {
			cmd.ui.Say(value.Name)
		}
	} else {
		cmd.ui.Say(T("No staging security group set"))
	}
	return nil
}
