package commandsloader

import (
	"github.com/liamawhite/cli-with-i18n/cf/commands"
	"github.com/liamawhite/cli-with-i18n/cf/commands/application"
	"github.com/liamawhite/cli-with-i18n/cf/commands/buildpack"
	"github.com/liamawhite/cli-with-i18n/cf/commands/domain"
	"github.com/liamawhite/cli-with-i18n/cf/commands/environmentvariablegroup"
	"github.com/liamawhite/cli-with-i18n/cf/commands/featureflag"
	"github.com/liamawhite/cli-with-i18n/cf/commands/organization"
	"github.com/liamawhite/cli-with-i18n/cf/commands/plugin"
	"github.com/liamawhite/cli-with-i18n/cf/commands/pluginrepo"
	"github.com/liamawhite/cli-with-i18n/cf/commands/quota"
	"github.com/liamawhite/cli-with-i18n/cf/commands/route"
	"github.com/liamawhite/cli-with-i18n/cf/commands/routergroups"
	"github.com/liamawhite/cli-with-i18n/cf/commands/securitygroup"
	"github.com/liamawhite/cli-with-i18n/cf/commands/service"
	"github.com/liamawhite/cli-with-i18n/cf/commands/serviceaccess"
	"github.com/liamawhite/cli-with-i18n/cf/commands/serviceauthtoken"
	"github.com/liamawhite/cli-with-i18n/cf/commands/servicebroker"
	"github.com/liamawhite/cli-with-i18n/cf/commands/servicekey"
	"github.com/liamawhite/cli-with-i18n/cf/commands/space"
	"github.com/liamawhite/cli-with-i18n/cf/commands/spacequota"
	"github.com/liamawhite/cli-with-i18n/cf/commands/user"
)

/*******************
This package make a reference to all the command packages
in cf/commands/..., so all init() in the directories will
get initialized

* Any new command packages must be included here for init() to get called
********************/

func Load() {
	_ = commands.API{}
	_ = application.ListApps{}
	_ = buildpack.ListBuildpacks{}
	_ = domain.CreateDomain{}
	_ = environmentvariablegroup.RunningEnvironmentVariableGroup{}
	_ = featureflag.ShowFeatureFlag{}
	_ = organization.ListOrgs{}
	_ = plugin.Plugins{}
	_ = pluginrepo.RepoPlugins{}
	_ = quota.CreateQuota{}
	_ = route.CreateRoute{}
	_ = routergroups.RouterGroups{}
	_ = securitygroup.ShowSecurityGroup{}
	_ = service.ShowService{}
	_ = serviceauthtoken.ListServiceAuthTokens{}
	_ = serviceaccess.ServiceAccess{}
	_ = servicebroker.ListServiceBrokers{}
	_ = servicekey.ServiceKey{}
	_ = space.CreateSpace{}
	_ = spacequota.SpaceQuota{}
	_ = user.CreateUser{}
}
