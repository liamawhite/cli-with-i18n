package v3action

import (
	"net/url"

	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccv3"
)

type ApplicationWithProcessSummary struct {
	Application
	ProcessSummaries ProcessSummaries
}

func (actor Actor) GetApplicationsWithProcessesBySpace(spaceGUID string) ([]ApplicationWithProcessSummary, Warnings, error) {
	var allWarnings Warnings

	apps, warnings, err := actor.CloudControllerClient.GetApplications(url.Values{
		ccv3.SpaceGUIDFilter: []string{spaceGUID},
		ccv3.OrderBy:         []string{ccv3.NameOrder},
	})
	allWarnings = Warnings(warnings)
	if err != nil {
		return nil, allWarnings, err
	}

	var appSummaries []ApplicationWithProcessSummary

	for _, app := range apps {
		processSummaries, processWarnings, err := actor.getProcessSummariesForApp(app.GUID)
		allWarnings = append(allWarnings, processWarnings...)
		if err != nil {
			return nil, allWarnings, err
		}

		appSummaries = append(appSummaries, ApplicationWithProcessSummary{
			Application: Application{
				Name:  app.Name,
				GUID:  app.GUID,
				State: app.State,
				Lifecycle: AppLifecycle{
					Type: AppLifecycleType(app.Lifecycle.Type),
					Data: AppLifecycleData(app.Lifecycle.Data),
				},
			},
			ProcessSummaries: processSummaries,
		})
	}

	return appSummaries, allWarnings, nil
}
