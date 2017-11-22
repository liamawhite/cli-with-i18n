package spacequota

import (
	"fmt"

	"github.com/liamawhite/cli-with-i18n/cf/api/spacequotas"
	"github.com/liamawhite/cli-with-i18n/cf/api/spaces"
	"github.com/liamawhite/cli-with-i18n/cf/commandregistry"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/flags"
	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"
	"github.com/liamawhite/cli-with-i18n/cf/terminal"
)

type UnsetSpaceQuota struct {
	ui        terminal.UI
	config    coreconfig.Reader
	quotaRepo spacequotas.SpaceQuotaRepository
	spaceRepo spaces.SpaceRepository
}

func init() {
	commandregistry.Register(&UnsetSpaceQuota{})
}

func (cmd *UnsetSpaceQuota) MetaData() commandregistry.CommandMetadata {
	return commandregistry.CommandMetadata{
		Name:        "unset-space-quota",
		Description: T("Unassign a quota from a space"),
		Usage: []string{
			T("CF_NAME unset-space-quota SPACE QUOTA\n\n"),
		},
	}
}

func (cmd *UnsetSpaceQuota) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	if len(fc.Args()) != 2 {
		cmd.ui.Failed(T("Incorrect Usage. Requires SPACE and QUOTA as arguments\n\n") + commandregistry.Commands.CommandUsage("unset-space-quota"))
		return nil, fmt.Errorf("Incorrect usage: %d arguments of %d required", len(fc.Args()), 2)
	}

	reqs := []requirements.Requirement{
		requirementsFactory.NewLoginRequirement(),
		requirementsFactory.NewTargetedOrgRequirement(),
	}

	return reqs, nil
}

func (cmd *UnsetSpaceQuota) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	cmd.ui = deps.UI
	cmd.config = deps.Config
	cmd.spaceRepo = deps.RepoLocator.GetSpaceRepository()
	cmd.quotaRepo = deps.RepoLocator.GetSpaceQuotaRepository()
	return cmd
}

func (cmd *UnsetSpaceQuota) Execute(c flags.FlagContext) error {
	spaceName := c.Args()[0]
	quotaName := c.Args()[1]

	space, err := cmd.spaceRepo.FindByName(spaceName)
	if err != nil {
		return err
	}

	quota, err := cmd.quotaRepo.FindByName(quotaName)
	if err != nil {
		return err
	}

	cmd.ui.Say(T("Unassigning space quota {{.QuotaName}} from space {{.SpaceName}} as {{.Username}}...",
		map[string]interface{}{
			"QuotaName": terminal.EntityNameColor(quota.Name),
			"SpaceName": terminal.EntityNameColor(space.Name),
			"Username":  terminal.EntityNameColor(cmd.config.Username())}))

	err = cmd.quotaRepo.UnassignQuotaFromSpace(space.GUID, quota.GUID)
	if err != nil {
		return err
	}

	cmd.ui.Ok()
	return nil
}
