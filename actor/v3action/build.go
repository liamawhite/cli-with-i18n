package v3action

import (
	"errors"
	"time"

	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccv3"
)

type StagingTimeoutError struct {
	AppName string
	Timeout time.Duration
}

func (StagingTimeoutError) Error() string {
	return "Timed out waiting for package to stage"
}

func (actor Actor) StagePackage(packageGUID string, appName string) (<-chan Droplet, <-chan Warnings, <-chan error) {
	dropletStream := make(chan Droplet)
	warningsStream := make(chan Warnings)
	errorStream := make(chan error)

	go func() {
		defer close(dropletStream)
		defer close(warningsStream)
		defer close(errorStream)

		build := ccv3.Build{PackageGUID: packageGUID}
		build, allWarnings, err := actor.CloudControllerClient.CreateBuild(build)
		warningsStream <- Warnings(allWarnings)

		if err != nil {
			errorStream <- err
			return
		}

		timeout := time.Now().Add(actor.Config.StagingTimeout())

		for time.Now().Before(timeout) {
			var warnings ccv3.Warnings
			build, warnings, err = actor.CloudControllerClient.GetBuild(build.GUID)
			warningsStream <- Warnings(warnings)
			if err != nil {
				errorStream <- err
				return
			}

			switch build.State {
			case ccv3.BuildStateFailed:
				errorStream <- errors.New(build.Error)
				return
			case ccv3.BuildStateStaging:
				time.Sleep(actor.Config.PollingInterval())
			default:

				//TODO: uncommend after #150569020
				// ccv3Droplet, warnings, err := actor.CloudControllerClient.GetDroplet(build.DropletGUID)
				// warningsStream <- Warnings(warnings)
				// if err != nil {
				// 	errorStream <- err
				// 	return
				// }

				ccv3Droplet := ccv3.Droplet{
					GUID:      build.DropletGUID,
					State:     ccv3.DropletState(build.State),
					CreatedAt: build.CreatedAt,
				}

				dropletStream <- actor.convertCCToActorDroplet(ccv3Droplet)
				return
			}
		}

		errorStream <- StagingTimeoutError{AppName: appName, Timeout: actor.Config.StagingTimeout()}
	}()

	return dropletStream, warningsStream, errorStream
}
