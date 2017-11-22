package flag

import (
	"github.com/liamawhite/cli-with-i18n/types"
)

type Command struct {
	types.FilteredString
}

func (b *Command) UnmarshalFlag(val string) error {
	b.ParseValue(val)
	return nil
}
