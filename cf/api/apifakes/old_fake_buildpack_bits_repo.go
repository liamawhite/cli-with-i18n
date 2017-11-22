package apifakes

import (
	"github.com/liamawhite/cli-with-i18n/cf/errors"
	"github.com/liamawhite/cli-with-i18n/cf/models"
)

type OldFakeBuildpackBitsRepository struct {
	UploadBuildpackErr         bool
	UploadBuildpackAPIResponse error
	UploadBuildpackPath        string
}

func (repo *OldFakeBuildpackBitsRepository) UploadBuildpack(buildpack models.Buildpack, dir string) error {
	if repo.UploadBuildpackErr {
		return errors.New("Invalid buildpack")
	}

	repo.UploadBuildpackPath = dir
	return repo.UploadBuildpackAPIResponse
}
