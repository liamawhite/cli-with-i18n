package resources

import "github.com/liamawhite/cli-with-i18n/cf/models"

type ServiceBindingResource struct {
	Resource
	Entity ServiceBindingEntity
}

type ServiceBindingEntity struct {
	AppGUID string `json:"app_guid"`
}

func (resource ServiceBindingResource) ToFields() models.ServiceBindingFields {
	return models.ServiceBindingFields{
		URL:     resource.Metadata.URL,
		GUID:    resource.Metadata.GUID,
		AppGUID: resource.Entity.AppGUID,
	}
}
