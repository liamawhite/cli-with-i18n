package v2

import (
	"os"

	"github.com/liamawhite/cli-with-i18n/cf/cmd"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/flag"
)

type EnableFeatureFlagCommand struct {
	RequiredArgs    flag.Feature `positional-args:"yes"`
	usage           interface{}  `usage:"CF_NAME enable-feature-flag FEATURE_NAME"`
	relatedCommands interface{}  `related_commands:"disable-feature-flag, feature-flags"`
}

func (EnableFeatureFlagCommand) Setup(config command.Config, ui command.UI) error {
	return nil
}

func (EnableFeatureFlagCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
