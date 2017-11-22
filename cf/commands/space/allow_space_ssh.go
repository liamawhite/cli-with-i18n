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

type AllowSpaceSSH struct {
	ui        terminal.UI
	config    coreconfig.Reader
	spaceReq  requirements.SpaceRequirement
	spaceRepo spaces.SpaceRepository
}

func init() {
	commandregistry.Register(&AllowSpaceSSH{})
}

func (cmd *AllowSpaceSSH) MetaData() commandregistry.CommandMetadata {
	return commandregistry.CommandMetadata{
		Name:        "allow-space-ssh",
		Description: T("Allow SSH access for the space"),
		Usage: []string{
			T("CF_NAME allow-space-ssh SPACE_NAME"),
		},
	}
}

func (cmd *AllowSpaceSSH) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	if len(fc.Args()) != 1 {
		cmd.ui.Failed(T("Incorrect Usage. Requires SPACE_NAME as argument\n\n") + commandregistry.Commands.CommandUsage("allow-space-ssh"))
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

func (cmd *AllowSpaceSSH) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	cmd.ui = deps.UI
	cmd.config = deps.Config
	cmd.spaceRepo = deps.RepoLocator.GetSpaceRepository()
	return cmd
}

func (cmd *AllowSpaceSSH) Execute(fc flags.FlagContext) error {
	space := cmd.spaceReq.GetSpace()

	if space.AllowSSH {
		cmd.ui.Say(T("ssh support is already enabled in space '{{.SpaceName}}'",
			map[string]interface{}{
				"SpaceName": space.Name,
			},
		))
		return nil
	}

	cmd.ui.Say(T("Enabling ssh support for space '{{.SpaceName}}'...",
		map[string]interface{}{
			"SpaceName": space.Name,
		},
	))
	cmd.ui.Say("")

	err := cmd.spaceRepo.SetAllowSSH(space.GUID, true)
	if err != nil {
		return errors.New(T("Error enabling ssh support for space ") + space.Name + ": " + err.Error())
	}

	cmd.ui.Ok()
	return nil
}
