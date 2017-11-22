package formatters

import (
	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
)

func Allowed(allowed bool) string {
	if allowed {
		return T("allowed")
	}
	return T("disallowed")
}
