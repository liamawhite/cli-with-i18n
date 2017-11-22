package resources

import "github.com/liamawhite/cli-with-i18n/cf/models"

type FeatureFlagResource struct {
	Entity models.FeatureFlag
}

func (resource FeatureFlagResource) ToFields() (flag models.FeatureFlag) {
	flag.Name = resource.Entity.Name
	flag.Enabled = resource.Entity.Enabled
	flag.ErrorMessage = resource.Entity.ErrorMessage
	return
}
