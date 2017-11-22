package running

import (
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/models"
	"github.com/liamawhite/cli-with-i18n/cf/net"

	. "github.com/liamawhite/cli-with-i18n/cf/api/securitygroups/defaults"
)

const urlPath = "/v2/config/running_security_groups"

//go:generate counterfeiter . SecurityGroupsRepo

type SecurityGroupsRepo interface {
	BindToRunningSet(string) error
	List() ([]models.SecurityGroupFields, error)
	UnbindFromRunningSet(string) error
}

type cloudControllerRunningSecurityGroupRepo struct {
	repoBase DefaultSecurityGroupsRepoBase
}

func NewSecurityGroupsRepo(configRepo coreconfig.Reader, gateway net.Gateway) SecurityGroupsRepo {
	return &cloudControllerRunningSecurityGroupRepo{
		repoBase: DefaultSecurityGroupsRepoBase{
			ConfigRepo: configRepo,
			Gateway:    gateway,
		},
	}
}

func (repo *cloudControllerRunningSecurityGroupRepo) BindToRunningSet(groupGUID string) error {
	return repo.repoBase.Bind(groupGUID, urlPath)
}

func (repo *cloudControllerRunningSecurityGroupRepo) List() ([]models.SecurityGroupFields, error) {
	return repo.repoBase.List(urlPath)
}

func (repo *cloudControllerRunningSecurityGroupRepo) UnbindFromRunningSet(groupGUID string) error {
	return repo.repoBase.Delete(groupGUID, urlPath)
}
