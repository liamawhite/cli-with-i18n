package resources

import "github.com/liamawhite/cli-with-i18n/cf/models"

type PaginatedStackResources struct {
	Resources []StackResource
}

type StackResource struct {
	Resource
	Entity StackEntity
}

type StackEntity struct {
	Name        string
	Description string
}

func (resource StackResource) ToFields() *models.Stack {
	return &models.Stack{
		GUID:        resource.Metadata.GUID,
		Name:        resource.Entity.Name,
		Description: resource.Entity.Description,
	}
}
