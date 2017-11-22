package staging

import (
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/models"
	"github.com/liamawhite/cli-with-i18n/cf/net"

	. "github.com/liamawhite/cli-with-i18n/cf/api/securitygroups/defaults"
)

const urlPath = "/v2/config/staging_security_groups"

//go:generate counterfeiter . SecurityGroupsRepo

type SecurityGroupsRepo interface {
	BindToStagingSet(string) error
	List() ([]models.SecurityGroupFields, error)
	UnbindFromStagingSet(string) error
}

type cloudControllerStagingSecurityGroupRepo struct {
	repoBase DefaultSecurityGroupsRepoBase
}

func NewSecurityGroupsRepo(configRepo coreconfig.Reader, gateway net.Gateway) SecurityGroupsRepo {
	return &cloudControllerStagingSecurityGroupRepo{
		repoBase: DefaultSecurityGroupsRepoBase{
			ConfigRepo: configRepo,
			Gateway:    gateway,
		},
	}
}

func (repo *cloudControllerStagingSecurityGroupRepo) BindToStagingSet(groupGUID string) error {
	return repo.repoBase.Bind(groupGUID, urlPath)
}

func (repo *cloudControllerStagingSecurityGroupRepo) List() ([]models.SecurityGroupFields, error) {
	return repo.repoBase.List(urlPath)
}

func (repo *cloudControllerStagingSecurityGroupRepo) UnbindFromStagingSet(groupGUID string) error {
	return repo.repoBase.Delete(groupGUID, urlPath)
}
