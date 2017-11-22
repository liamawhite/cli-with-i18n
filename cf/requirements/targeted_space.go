package requirements

import (
	"fmt"

	"errors"

	"github.com/liamawhite/cli-with-i18n/cf"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
	"github.com/liamawhite/cli-with-i18n/cf/terminal"
)

type TargetedSpaceRequirement struct {
	config coreconfig.Reader
}

func NewTargetedSpaceRequirement(config coreconfig.Reader) TargetedSpaceRequirement {
	return TargetedSpaceRequirement{config}
}

func (req TargetedSpaceRequirement) Execute() error {
	if !req.config.HasOrganization() {
		message := fmt.Sprintf(T("No org and space targeted, use '{{.Command}}' to target an org and space", map[string]interface{}{"Command": terminal.CommandColor(cf.Name + " target -o ORG -s SPACE")}))
		return errors.New(message)
	}

	if !req.config.HasSpace() {
		message := fmt.Sprintf(T("No space targeted, use '{{.Command}}' to target a space.", map[string]interface{}{"Command": terminal.CommandColor("cf target -s")}))
		return errors.New(message)
	}

	return nil
}
