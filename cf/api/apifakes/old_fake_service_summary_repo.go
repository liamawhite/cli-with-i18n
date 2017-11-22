package apifakes

import "github.com/liamawhite/cli-with-i18n/cf/models"

type OldFakeServiceSummaryRepo struct {
	GetSummariesInCurrentSpaceInstances []models.ServiceInstance
}

func (repo *OldFakeServiceSummaryRepo) GetSummariesInCurrentSpace() (instances []models.ServiceInstance, apiErr error) {
	instances = repo.GetSummariesInCurrentSpaceInstances
	return
}
