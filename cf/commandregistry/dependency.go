package commandregistry

import (
	"fmt"
	"io"
	"os"
	"time"

	"path/filepath"

	"github.com/liamawhite/cli-with-i18n/cf/actors"
	"github.com/liamawhite/cli-with-i18n/cf/actors/brokerbuilder"
	"github.com/liamawhite/cli-with-i18n/cf/actors/planbuilder"
	"github.com/liamawhite/cli-with-i18n/cf/actors/pluginrepo"
	"github.com/liamawhite/cli-with-i18n/cf/actors/servicebuilder"
	"github.com/liamawhite/cli-with-i18n/cf/api"
	"github.com/liamawhite/cli-with-i18n/cf/appfiles"
	"github.com/liamawhite/cli-with-i18n/cf/configuration"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/confighelpers"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/pluginconfig"
	"github.com/liamawhite/cli-with-i18n/cf/manifest"
	"github.com/liamawhite/cli-with-i18n/cf/net"
	"github.com/liamawhite/cli-with-i18n/cf/terminal"
	"github.com/liamawhite/cli-with-i18n/cf/trace"
	"github.com/liamawhite/cli-with-i18n/plugin/models"
	"github.com/liamawhite/cli-with-i18n/util"
	"github.com/liamawhite/cli-with-i18n/util/words/generator"
)

type Dependency struct {
	UI                 terminal.UI
	Config             coreconfig.Repository
	RepoLocator        api.RepositoryLocator
	PluginConfig       pluginconfig.PluginConfiguration
	ManifestRepo       manifest.Repository
	AppManifest        manifest.App
	Gateways           map[string]net.Gateway
	TeePrinter         *terminal.TeePrinter
	PluginRepo         pluginrepo.PluginRepo
	PluginModels       *PluginModels
	ServiceBuilder     servicebuilder.ServiceBuilder
	BrokerBuilder      brokerbuilder.Builder
	PlanBuilder        planbuilder.PlanBuilder
	ServiceHandler     actors.ServiceActor
	ServicePlanHandler actors.ServicePlanActor
	WordGenerator      generator.WordGenerator
	AppZipper          appfiles.Zipper
	AppFiles           appfiles.AppFiles
	PushActor          actors.PushActor
	RouteActor         actors.RouteActor
	ChecksumUtil       util.Sha1Checksum
	WildcardDependency interface{} //use for injecting fakes
	Logger             trace.Printer
}

type PluginModels struct {
	Application   *plugin_models.GetAppModel
	AppsSummary   *[]plugin_models.GetAppsModel
	Organizations *[]plugin_models.GetOrgs_Model
	Organization  *plugin_models.GetOrg_Model
	Spaces        *[]plugin_models.GetSpaces_Model
	Space         *plugin_models.GetSpace_Model
	OrgUsers      *[]plugin_models.GetOrgUsers_Model
	SpaceUsers    *[]plugin_models.GetSpaceUsers_Model
	Services      *[]plugin_models.GetServices_Model
	Service       *plugin_models.GetService_Model
	OauthToken    *plugin_models.GetOauthToken_Model
}

func NewDependency(writer io.Writer, logger trace.Printer, envDialTimeout string) Dependency {
	deps := Dependency{}
	deps.TeePrinter = terminal.NewTeePrinter(writer)
	deps.UI = terminal.NewUI(os.Stdin, writer, deps.TeePrinter, logger)

	errorHandler := func(err error) {
		if err != nil {
			deps.UI.Failed(fmt.Sprintf("Config error: %s", err))
		}
	}

	configPath, err := confighelpers.DefaultFilePath()
	if err != nil {
		errorHandler(err)
	}
	deps.Config = coreconfig.NewRepositoryFromFilepath(configPath, errorHandler)

	deps.ManifestRepo = manifest.NewDiskRepository()
	deps.AppManifest = manifest.NewGenerator()

	pluginPath := filepath.Join(confighelpers.PluginRepoDir(), ".cf", "plugins")
	deps.PluginConfig = pluginconfig.NewPluginConfig(
		errorHandler,
		configuration.NewDiskPersistor(filepath.Join(pluginPath, "config.json")),
		pluginPath,
	)

	terminal.UserAskedForColors = deps.Config.ColorEnabled()
	terminal.InitColorSupport()

	deps.Gateways = map[string]net.Gateway{
		"cloud-controller": net.NewCloudControllerGateway(deps.Config, time.Now, deps.UI, logger, envDialTimeout),
		"uaa":              net.NewUAAGateway(deps.Config, deps.UI, logger, envDialTimeout),
		"routing-api":      net.NewRoutingAPIGateway(deps.Config, time.Now, deps.UI, logger, envDialTimeout),
	}
	deps.RepoLocator = api.NewRepositoryLocator(deps.Config, deps.Gateways, logger, envDialTimeout)

	deps.PluginModels = &PluginModels{Application: nil}

	deps.PlanBuilder = planbuilder.NewBuilder(
		deps.RepoLocator.GetServicePlanRepository(),
		deps.RepoLocator.GetServicePlanVisibilityRepository(),
		deps.RepoLocator.GetOrganizationRepository(),
	)

	deps.ServiceBuilder = servicebuilder.NewBuilder(
		deps.RepoLocator.GetServiceRepository(),
		deps.PlanBuilder,
	)

	deps.BrokerBuilder = brokerbuilder.NewBuilder(
		deps.RepoLocator.GetServiceBrokerRepository(),
		deps.ServiceBuilder,
	)

	deps.PluginRepo = pluginrepo.NewPluginRepo()

	deps.ServiceHandler = actors.NewServiceHandler(
		deps.RepoLocator.GetOrganizationRepository(),
		deps.BrokerBuilder,
		deps.ServiceBuilder,
	)

	deps.ServicePlanHandler = actors.NewServicePlanHandler(
		deps.RepoLocator.GetServicePlanRepository(),
		deps.RepoLocator.GetServicePlanVisibilityRepository(),
		deps.RepoLocator.GetOrganizationRepository(),
		deps.PlanBuilder,
		deps.ServiceBuilder,
	)

	deps.WordGenerator = generator.NewWordGenerator()

	deps.AppZipper = appfiles.ApplicationZipper{}
	deps.AppFiles = appfiles.ApplicationFiles{}

	deps.RouteActor = actors.NewRouteActor(deps.UI, deps.RepoLocator.GetRouteRepository(), deps.RepoLocator.GetDomainRepository())
	deps.PushActor = actors.NewPushActor(deps.RepoLocator.GetApplicationBitsRepository(), deps.AppZipper, deps.AppFiles, deps.RouteActor)

	deps.ChecksumUtil = util.NewSha1Checksum("")

	deps.Logger = logger

	return deps
}
