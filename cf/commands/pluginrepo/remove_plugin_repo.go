package pluginrepo

import (
	"errors"
	"fmt"
	"strings"

	"github.com/liamawhite/cli-with-i18n/cf/commandregistry"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/flags"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"
	"github.com/liamawhite/cli-with-i18n/cf/terminal"

	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
)

type RemovePluginRepo struct {
	ui     terminal.UI
	config coreconfig.ReadWriter
}

func init() {
	commandregistry.Register(&RemovePluginRepo{})
}

func (cmd *RemovePluginRepo) MetaData() commandregistry.CommandMetadata {
	return commandregistry.CommandMetadata{
		Name:        "remove-plugin-repo",
		Description: T("Remove a plugin repository"),
		Usage: []string{
			T("CF_NAME remove-plugin-repo REPO_NAME"),
		},
		Examples: []string{
			"CF_NAME remove-plugin-repo PrivateRepo",
		},
		TotalArgs: 1,
	}
}

func (cmd *RemovePluginRepo) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	if len(fc.Args()) != 1 {
		cmd.ui.Failed(T("Incorrect Usage. Requires an argument\n\n") + commandregistry.Commands.CommandUsage("remove-plugin-repo"))
		return nil, fmt.Errorf("Incorrect usage: %d arguments of %d required", len(fc.Args()), 1)
	}

	reqs := []requirements.Requirement{}
	return reqs, nil
}

func (cmd *RemovePluginRepo) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	cmd.ui = deps.UI
	cmd.config = deps.Config
	return cmd
}

func (cmd *RemovePluginRepo) Execute(c flags.FlagContext) error {
	cmd.ui.Say("")
	repoName := strings.Trim(c.Args()[0], " ")

	if i := cmd.findRepoIndex(repoName); i != -1 {
		cmd.config.UnSetPluginRepo(i)
		cmd.ui.Ok()
		cmd.ui.Say(repoName + T(" removed from list of repositories"))
		cmd.ui.Say("")
	} else {
		return errors.New(repoName + T(" does not exist as a repo"))
	}
	return nil
}

func (cmd RemovePluginRepo) findRepoIndex(repoName string) int {
	repos := cmd.config.PluginRepos()
	for i, repo := range repos {
		if strings.ToLower(repo.Name) == strings.ToLower(repoName) {
			return i
		}
	}
	return -1
}
