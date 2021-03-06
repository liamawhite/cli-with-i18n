package requirements

import (
	"github.com/liamawhite/cli-with-i18n/cf/api"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/models"
)

//go:generate counterfeiter . DomainRequirement

type DomainRequirement interface {
	Requirement
	GetDomain() models.DomainFields
}

type domainAPIRequirement struct {
	name       string
	config     coreconfig.Reader
	domainRepo api.DomainRepository
	domain     models.DomainFields
}

func NewDomainRequirement(name string, config coreconfig.Reader, domainRepo api.DomainRepository) (req *domainAPIRequirement) {
	req = new(domainAPIRequirement)
	req.name = name
	req.config = config
	req.domainRepo = domainRepo
	return
}

func (req *domainAPIRequirement) Execute() error {
	var apiErr error
	req.domain, apiErr = req.domainRepo.FindByNameInOrg(req.name, req.config.OrganizationFields().GUID)

	if apiErr != nil {
		return apiErr
	}

	return nil
}

func (req *domainAPIRequirement) GetDomain() models.DomainFields {
	return req.domain
}
