package plugin

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type RemovePluginRepoCommand struct {
	RequiredArgs    flag.PluginRepoName `positional-args:"yes"`
	usage           interface{}         `usage:"CF_NAME remove-plugin-repo REPO_NAME\n\nEXAMPLES:\n   CF_NAME remove-plugin-repo PrivateRepo"`
	relatedCommands interface{}         `related_commands:"list-plugin-repos"`
}

func (RemovePluginRepoCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (RemovePluginRepoCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
