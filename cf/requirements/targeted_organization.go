package requirements

import (
	"errors"
	"fmt"

	"github.com/liamawhite/cli-with-i18n/cf"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
	"github.com/liamawhite/cli-with-i18n/cf/models"
	"github.com/liamawhite/cli-with-i18n/cf/terminal"
)

//go:generate counterfeiter . TargetedOrgRequirement

type TargetedOrgRequirement interface {
	Requirement
	GetOrganizationFields() models.OrganizationFields
}

type targetedOrgAPIRequirement struct {
	config coreconfig.Reader
}

func NewTargetedOrgRequirement(config coreconfig.Reader) TargetedOrgRequirement {
	return targetedOrgAPIRequirement{config}
}

func (req targetedOrgAPIRequirement) Execute() error {
	if !req.config.HasOrganization() {
		message := fmt.Sprintf(T("No org targeted, use '{{.Command}}' to target an org.", map[string]interface{}{"Command": terminal.CommandColor(cf.Name + " target -o ORG")}))
		return errors.New(message)
	}

	return nil
}

func (req targetedOrgAPIRequirement) GetOrganizationFields() (org models.OrganizationFields) {
	return req.config.OrganizationFields()
}
