package v2action

import (
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccv2"
	"github.com/liamawhite/cli-with-i18n/util/manifest"
)

func (actor Actor) CreateApplicationManifestByNameAndSpace(appName string, spaceGUID string, pathToFile string) (Warnings, error) {

	var allWarnings Warnings
	applicationSummary, appSummaryWarnings, err := actor.GetApplicationSummaryByNameAndSpace(appName, spaceGUID)
	allWarnings = append(allWarnings, appSummaryWarnings...)
	if err != nil {
		return allWarnings, err
	}

	serviceInstances, serviceWarnings, err := actor.GetServiceInstancesByApplication(applicationSummary.GUID)
	allWarnings = append(allWarnings, serviceWarnings...)
	if err != nil {
		return allWarnings, err
	}

	var routes []string
	for _, route := range applicationSummary.Routes {
		routes = append(routes, route.String())
	}

	var services []string
	for _, serviceInstace := range serviceInstances {
		services = append(services, serviceInstace.Name)
	}

	manifestApp := manifest.Application{
		Buildpack:            applicationSummary.Buildpack,
		Command:              applicationSummary.Command,
		DockerImage:          applicationSummary.DockerImage,
		DockerUsername:       applicationSummary.DockerCredentials.Username,
		EnvironmentVariables: applicationSummary.EnvironmentVariables,
		HealthCheckTimeout:   applicationSummary.HealthCheckTimeout,
		Instances:            applicationSummary.Instances,
		Name:                 applicationSummary.Name,
		Routes:               routes,
		Services:             services,
		StackName:            applicationSummary.Stack.Name,
	}
	manifestApp.DiskQuota.ParseUint64Value(&applicationSummary.DiskQuota)
	manifestApp.Memory.ParseUint64Value(&applicationSummary.Memory)
	if len(routes) < 1 {
		manifestApp.NoRoute = true
	}

	if applicationSummary.HealthCheckType != ccv2.ApplicationHealthCheckPort {
		manifestApp.HealthCheckType = string(applicationSummary.HealthCheckType)

		if applicationSummary.HealthCheckType == ccv2.ApplicationHealthCheckHTTP &&
			applicationSummary.HealthCheckHTTPEndpoint != "/" {
			manifestApp.HealthCheckHTTPEndpoint = applicationSummary.HealthCheckHTTPEndpoint
		}
	}

	err = manifest.WriteApplicationManifest(manifestApp, pathToFile)
	return allWarnings, err
}
