package flag

import (
	"github.com/liamawhite/cli-with-i18n/types"
)

type Buildpack struct {
	types.FilteredString
}

func (b *Buildpack) UnmarshalFlag(val string) error {
	b.ParseValue(val)
	return nil
}
