package cloudcontroller

import (
	"github.com/blang/semver"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccerror"
)

// MinimumAPIVersionCheck compares `current` to `minimum`.  If `current` is
// older than `minimum` then an error is returned; otherwise, nil is returned.
func MinimumAPIVersionCheck(current string, minimum string) error {
	if minimum == "" {
		return nil
	}

	currentSemvar, err := semver.Make(current)
	if err != nil {
		return err
	}

	minimumSemvar, err := semver.Make(minimum)
	if err != nil {
		return err
	}

	if currentSemvar.Compare(minimumSemvar) == -1 {
		return ccerror.MinimumAPIVersionNotMetError{
			CurrentVersion: current,
			MinimumVersion: minimum,
		}
	}

	return nil
}
