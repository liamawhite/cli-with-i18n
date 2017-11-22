package defaults

import (
	"fmt"

	"github.com/liamawhite/cli-with-i18n/cf/api/resources"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/models"
	"github.com/liamawhite/cli-with-i18n/cf/net"
)

type DefaultSecurityGroupsRepoBase struct {
	ConfigRepo coreconfig.Reader
	Gateway    net.Gateway
}

func (repo *DefaultSecurityGroupsRepoBase) Bind(groupGUID string, path string) error {
	updatedPath := fmt.Sprintf("%s/%s", path, groupGUID)
	return repo.Gateway.UpdateResourceFromStruct(repo.ConfigRepo.APIEndpoint(), updatedPath, "")
}

func (repo *DefaultSecurityGroupsRepoBase) List(path string) ([]models.SecurityGroupFields, error) {
	groups := []models.SecurityGroupFields{}

	err := repo.Gateway.ListPaginatedResources(
		repo.ConfigRepo.APIEndpoint(),
		path,
		resources.SecurityGroupResource{},
		func(resource interface{}) bool {
			if securityGroupResource, ok := resource.(resources.SecurityGroupResource); ok {
				groups = append(groups, securityGroupResource.ToFields())
			}

			return true
		},
	)

	return groups, err
}

func (repo *DefaultSecurityGroupsRepoBase) Delete(groupGUID string, path string) error {
	updatedPath := fmt.Sprintf("%s/%s", path, groupGUID)
	return repo.Gateway.DeleteResource(repo.ConfigRepo.APIEndpoint(), updatedPath)
}
