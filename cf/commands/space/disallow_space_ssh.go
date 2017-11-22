package space

import (
	"errors"
	"fmt"

	"github.com/liamawhite/cli-with-i18n/cf/api/spaces"
	"github.com/liamawhite/cli-with-i18n/cf/commandregistry"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/flags"
	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"
	"github.com/liamawhite/cli-with-i18n/cf/terminal"
)

type DisallowSpaceSSH struct {
	ui        terminal.UI
	config    coreconfig.Reader
	spaceReq  requirements.SpaceRequirement
	spaceRepo spaces.SpaceRepository
}

func init() {
	commandregistry.Register(&DisallowSpaceSSH{})
}

func (cmd *DisallowSpaceSSH) MetaData() commandregistry.CommandMetadata {
	return commandregistry.CommandMetadata{
		Name:        "disallow-space-ssh",
		Description: T("Disallow SSH access for the space"),
		Usage: []string{
			T("CF_NAME disallow-space-ssh SPACE_NAME"),
		},
	}
}

func (cmd *DisallowSpaceSSH) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	if len(fc.Args()) != 1 {
		cmd.ui.Failed(T("Incorrect Usage. Requires SPACE_NAME as argument\n\n") + commandregistry.Commands.CommandUsage("disallow-space-ssh"))
		return nil, fmt.Errorf("Incorrect usage: %d arguments of %d required", len(fc.Args()), 1)
	}

	cmd.spaceReq = requirementsFactory.NewSpaceRequirement(fc.Args()[0])

	reqs := []requirements.Requirement{
		requirementsFactory.NewLoginRequirement(),
		requirementsFactory.NewTargetedOrgRequirement(),
		cmd.spaceReq,
	}

	return reqs, nil
}

func (cmd *DisallowSpaceSSH) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	cmd.ui = deps.UI
	cmd.config = deps.Config
	cmd.spaceRepo = deps.RepoLocator.GetSpaceRepository()
	return cmd
}

func (cmd *DisallowSpaceSSH) Execute(fc flags.FlagContext) error {
	space := cmd.spaceReq.GetSpace()

	if !space.AllowSSH {
		cmd.ui.Say(fmt.Sprintf(T("ssh support is already disabled in space ")+"'%s'", space.Name))
		return nil
	}

	cmd.ui.Say(T("Disabling ssh support for space '{{.SpaceName}}'...",
		map[string]interface{}{
			"SpaceName": space.Name,
		},
	))
	cmd.ui.Say("")

	err := cmd.spaceRepo.SetAllowSSH(space.GUID, false)
	if err != nil {
		return errors.New(T("Error disabling ssh support for space ") + space.Name + ": " + err.Error())
	}

	cmd.ui.Ok()
	return nil
}
