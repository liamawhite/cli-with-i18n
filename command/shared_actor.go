package command

import "github.com/liamawhite/cli-with-i18n/actor/sharedaction"

//go:generate counterfeiter . SharedActor

type SharedActor interface {
	CheckTarget(config sharedaction.Config, targetedOrganizationRequired bool, targetedSpaceRequired bool) error
}
