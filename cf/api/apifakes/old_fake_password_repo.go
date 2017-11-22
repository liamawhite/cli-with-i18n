package apifakes

import "github.com/liamawhite/cli-with-i18n/cf/errors"

type OldFakePasswordRepo struct {
	Score          string
	ScoredPassword string

	UpdateUnauthorized bool
	UpdateNewPassword  string
	UpdateOldPassword  string
}

func (repo *OldFakePasswordRepo) UpdatePassword(old string, new string) (apiErr error) {
	repo.UpdateOldPassword = old
	repo.UpdateNewPassword = new

	if repo.UpdateUnauthorized {
		apiErr = errors.NewHTTPError(401, "unauthorized", "Authorization Failed")
	}

	return
}
