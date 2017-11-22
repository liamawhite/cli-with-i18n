package stacks

import (
	"fmt"
	"net/url"

	"github.com/liamawhite/cli-with-i18n/cf/api/resources"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/errors"
	"github.com/liamawhite/cli-with-i18n/cf/models"
	"github.com/liamawhite/cli-with-i18n/cf/net"

	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
)

//go:generate counterfeiter . StackRepository

type StackRepository interface {
	FindByName(name string) (stack models.Stack, apiErr error)
	FindByGUID(guid string) (models.Stack, error)
	FindAll() (stacks []models.Stack, apiErr error)
}

type CloudControllerStackRepository struct {
	config  coreconfig.Reader
	gateway net.Gateway
}

func NewCloudControllerStackRepository(config coreconfig.Reader, gateway net.Gateway) (repo CloudControllerStackRepository) {
	repo.config = config
	repo.gateway = gateway
	return
}

func (repo CloudControllerStackRepository) FindByGUID(guid string) (models.Stack, error) {
	stackRequest := resources.StackResource{}
	path := fmt.Sprintf("%s/v2/stacks/%s", repo.config.APIEndpoint(), guid)
	err := repo.gateway.GetResource(path, &stackRequest)
	if err != nil {
		if errNotFound, ok := err.(*errors.HTTPNotFoundError); ok {
			return models.Stack{}, errNotFound
		}

		return models.Stack{}, fmt.Errorf(T("Error retrieving stacks: {{.Error}}", map[string]interface{}{
			"Error": err.Error(),
		}))
	}

	return *stackRequest.ToFields(), nil
}

func (repo CloudControllerStackRepository) FindByName(name string) (stack models.Stack, apiErr error) {
	path := fmt.Sprintf("/v2/stacks?q=%s", url.QueryEscape("name:"+name))
	stacks, apiErr := repo.findAllWithPath(path)
	if apiErr != nil {
		return
	}

	if len(stacks) == 0 {
		apiErr = errors.NewModelNotFoundError("Stack", name)
		return
	}

	stack = stacks[0]
	return
}

func (repo CloudControllerStackRepository) FindAll() (stacks []models.Stack, apiErr error) {
	return repo.findAllWithPath("/v2/stacks")
}

func (repo CloudControllerStackRepository) findAllWithPath(path string) ([]models.Stack, error) {
	var stacks []models.Stack
	apiErr := repo.gateway.ListPaginatedResources(
		repo.config.APIEndpoint(),
		path,
		resources.StackResource{},
		func(resource interface{}) bool {
			if sr, ok := resource.(resources.StackResource); ok {
				stacks = append(stacks, *sr.ToFields())
			}
			return true
		})
	return stacks, apiErr
}
