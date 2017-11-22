package shared

import (
	"github.com/cloudfoundry/bytefmt"
	"github.com/liamawhite/cli-with-i18n/actor/pushaction"
	"github.com/liamawhite/cli-with-i18n/util/ui"
)

func GetApplicationChanges(appConfig pushaction.ApplicationConfig) []ui.Change {
	changes := []ui.Change{
		{
			Header:       "name:",
			CurrentValue: appConfig.CurrentApplication.Name,
			NewValue:     appConfig.DesiredApplication.Name,
		},
	}

	if appConfig.DesiredApplication.DockerImage != "" {
		changes = append(changes,
			ui.Change{
				Header:       "docker image:",
				CurrentValue: appConfig.CurrentApplication.DockerImage,
				NewValue:     appConfig.DesiredApplication.DockerImage,
			})

		if appConfig.CurrentApplication.DockerCredentials.Username != "" || appConfig.DesiredApplication.DockerCredentials.Username != "" {
			changes = append(changes,
				ui.Change{
					Header:       "docker username:",
					CurrentValue: appConfig.CurrentApplication.DockerCredentials.Username,
					NewValue:     appConfig.DesiredApplication.DockerCredentials.Username,
				})
		}
	} else {
		changes = append(changes,
			ui.Change{
				Header:       "path:",
				CurrentValue: appConfig.Path,
				NewValue:     appConfig.Path,
			})
	}

	// Existing buildpack and existing detected buildpack are mutually exclusive
	oldBuildpack := appConfig.CurrentApplication.CalculatedBuildpack()
	newBuildpack := appConfig.DesiredApplication.CalculatedBuildpack()
	if oldBuildpack != "" || newBuildpack != "" {
		changes = append(changes,
			ui.Change{
				Header:       "buildpack:",
				CurrentValue: oldBuildpack,
				NewValue:     newBuildpack,
			})
	}

	// Existing command and existing detected start command are mutually exclusive
	oldCommand := appConfig.CurrentApplication.CalculatedCommand()
	newCommand := appConfig.DesiredApplication.CalculatedCommand()
	if oldCommand != "" || newCommand != "" {
		changes = append(changes,
			ui.Change{
				Header:       "command:",
				CurrentValue: oldCommand,
				NewValue:     newCommand,
			})
	}

	if appConfig.CurrentApplication.DiskQuota != 0 || appConfig.DesiredApplication.DiskQuota != 0 {
		var currentDiskQuota string
		if appConfig.CurrentApplication.DiskQuota != 0 {
			currentDiskQuota = MegabytesToString(appConfig.CurrentApplication.DiskQuota)
		}
		changes = append(changes,
			ui.Change{
				Header:       "disk quota:",
				CurrentValue: currentDiskQuota,
				NewValue:     MegabytesToString(appConfig.DesiredApplication.DiskQuota),
			})
	}

	if appConfig.CurrentApplication.HealthCheckHTTPEndpoint != "" || appConfig.DesiredApplication.HealthCheckHTTPEndpoint != "" {
		changes = append(changes,
			ui.Change{
				Header:       "health check http endpoint:",
				CurrentValue: appConfig.CurrentApplication.HealthCheckHTTPEndpoint,
				NewValue:     appConfig.DesiredApplication.HealthCheckHTTPEndpoint,
			})
	}

	if appConfig.CurrentApplication.HealthCheckTimeout != 0 || appConfig.DesiredApplication.HealthCheckTimeout != 0 {
		changes = append(changes,
			ui.Change{
				Header:       "health check timeout:",
				CurrentValue: appConfig.CurrentApplication.HealthCheckTimeout,
				NewValue:     appConfig.DesiredApplication.HealthCheckTimeout,
			})
	}

	if appConfig.CurrentApplication.HealthCheckType != "" || appConfig.DesiredApplication.HealthCheckType != "" {
		changes = append(changes,
			ui.Change{
				Header:       "health check type:",
				CurrentValue: string(appConfig.CurrentApplication.HealthCheckType),
				NewValue:     string(appConfig.DesiredApplication.HealthCheckType),
			})
	}

	if appConfig.CurrentApplication.Instances.IsSet || appConfig.DesiredApplication.Instances.IsSet {
		changes = append(changes,
			ui.Change{
				Header:       "instances:",
				CurrentValue: appConfig.CurrentApplication.Instances,
				NewValue:     appConfig.DesiredApplication.Instances,
			})
	}

	if appConfig.CurrentApplication.Memory != 0 || appConfig.DesiredApplication.Memory != 0 {
		var currentMemory string
		if appConfig.CurrentApplication.Memory != 0 {
			currentMemory = MegabytesToString(appConfig.CurrentApplication.Memory)
		}
		changes = append(changes,
			ui.Change{
				Header:       "memory:",
				CurrentValue: currentMemory,
				NewValue:     MegabytesToString(appConfig.DesiredApplication.Memory),
			})
	}

	if appConfig.CurrentApplication.Stack.Name != "" || appConfig.DesiredApplication.Stack.Name != "" {
		changes = append(changes,
			ui.Change{
				Header:       "stack:",
				CurrentValue: appConfig.CurrentApplication.Stack.Name,
				NewValue:     appConfig.DesiredApplication.Stack.Name,
			})
	}

	var oldServices []string
	for name := range appConfig.CurrentServices {
		oldServices = append(oldServices, name)
	}

	var newServices []string
	for name := range appConfig.DesiredServices {
		newServices = append(newServices, name)
	}

	changes = append(changes,
		ui.Change{
			Header:       "services:",
			CurrentValue: oldServices,
			NewValue:     newServices,
		})

	changes = append(changes,
		ui.Change{
			Header:       "env:",
			CurrentValue: appConfig.CurrentApplication.EnvironmentVariables,
			NewValue:     appConfig.DesiredApplication.EnvironmentVariables,
		})

	var currentRoutes []string
	for _, route := range appConfig.CurrentRoutes {
		currentRoutes = append(currentRoutes, route.String())
	}

	var desiredRotues []string
	for _, route := range appConfig.DesiredRoutes {
		desiredRotues = append(desiredRotues, route.String())
	}

	changes = append(changes,
		ui.Change{
			Header:       "routes:",
			CurrentValue: currentRoutes,
			NewValue:     desiredRotues,
		})

	return changes
}

func SelectNonBlankValue(str ...string) string {
	for _, s := range str {
		if s != "" {
			return s
		}
	}
	return ""
}

func MegabytesToString(value uint64) string {
	return bytefmt.ByteSize(bytefmt.MEGABYTE * uint64(value))
}
