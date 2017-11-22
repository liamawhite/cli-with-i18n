package flag

import (
	flags "github.com/jessevdk/go-flags"
	"github.com/liamawhite/cli-with-i18n/types"
)

type Instances struct {
	types.NullInt
}

func (i *Instances) UnmarshalFlag(val string) error {
	err := i.ParseStringValue(val)
	if err != nil || i.Value < 0 {
		return &flags.Error{
			Type:    flags.ErrRequired,
			Message: "invalid argument for flag '-i' (expected int > 0)",
		}
	}
	return nil
}
