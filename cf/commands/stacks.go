package commands

import (
	"github.com/liamawhite/cli-with-i18n/cf/api/stacks"
	"github.com/liamawhite/cli-with-i18n/cf/commandregistry"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/flags"
	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"
	"github.com/liamawhite/cli-with-i18n/cf/terminal"
)

type ListStacks struct {
	ui         terminal.UI
	config     coreconfig.Reader
	stacksRepo stacks.StackRepository
}

func init() {
	commandregistry.Register(&ListStacks{})
}

func (cmd *ListStacks) MetaData() commandregistry.CommandMetadata {
	return commandregistry.CommandMetadata{
		Name:        "stacks",
		Description: T("List all stacks (a stack is a pre-built file system, including an operating system, that can run apps)"),
		Usage: []string{
			T("CF_NAME stacks"),
		},
	}
}

func (cmd *ListStacks) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
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

func (cmd *ListStacks) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	cmd.ui = deps.UI
	cmd.config = deps.Config
	cmd.stacksRepo = deps.RepoLocator.GetStackRepository()
	return cmd
}

func (cmd *ListStacks) Execute(c flags.FlagContext) error {
	cmd.ui.Say(T("Getting stacks in org {{.OrganizationName}} / space {{.SpaceName}} as {{.Username}}...",
		map[string]interface{}{"OrganizationName": terminal.EntityNameColor(cmd.config.OrganizationFields().Name),
			"SpaceName": terminal.EntityNameColor(cmd.config.SpaceFields().Name),
			"Username":  terminal.EntityNameColor(cmd.config.Username())}))

	stacks, err := cmd.stacksRepo.FindAll()
	if err != nil {
		return err
	}

	cmd.ui.Ok()
	cmd.ui.Say("")

	table := cmd.ui.Table([]string{T("name"), T("description")})

	for _, stack := range stacks {
		table.Add(stack.Name, stack.Description)
	}

	err = table.Print()
	if err != nil {
		return err
	}
	return nil
}
