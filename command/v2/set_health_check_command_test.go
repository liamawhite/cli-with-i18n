package v2_test

import (
	"errors"

	"github.com/liamawhite/cli-with-i18n/actor/sharedaction"
	"github.com/liamawhite/cli-with-i18n/actor/v2action"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccv2"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccversion"
	"github.com/liamawhite/cli-with-i18n/command/commandfakes"
	"github.com/liamawhite/cli-with-i18n/command/translatableerror"
	. "github.com/liamawhite/cli-with-i18n/command/v2"
	"github.com/liamawhite/cli-with-i18n/command/v2/v2fakes"
	"github.com/liamawhite/cli-with-i18n/util/configv3"
	"github.com/liamawhite/cli-with-i18n/util/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("set-health-check Command", func() {
	var (
		cmd             SetHealthCheckCommand
		testUI          *ui.UI
		fakeConfig      *commandfakes.FakeConfig
		fakeSharedActor *commandfakes.FakeSharedActor
		fakeActor       *v2fakes.FakeSetHealthCheckActor
		binaryName      string
		executeErr      error
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(nil, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeSharedActor = new(commandfakes.FakeSharedActor)
		fakeActor = new(v2fakes.FakeSetHealthCheckActor)

		cmd = SetHealthCheckCommand{
			UI:          testUI,
			Config:      fakeConfig,
			SharedActor: fakeSharedActor,
			Actor:       fakeActor,
		}

		binaryName = "faceman"
		fakeConfig.BinaryNameReturns(binaryName)
		fakeConfig.TargetedOrganizationReturns(configv3.Organization{
			Name: "some-org",
		})
		fakeConfig.TargetedSpaceReturns(configv3.Space{
			GUID: "some-space-guid",
			Name: "some-space",
		})

		fakeConfig.CurrentUserReturns(configv3.User{Name: "some-user"}, nil)

		fakeActor.CloudControllerAPIVersionReturns(ccversion.MinVersionHTTPEndpointHealthCheckV2)
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	Context("when checking the target fails", func() {
		BeforeEach(func() {
			fakeSharedActor.CheckTargetReturns(
				sharedaction.NotLoggedInError{BinaryName: binaryName})
		})

		It("returns an error", func() {
			Expect(fakeSharedActor.CheckTargetCallCount()).To(Equal(1))
			config, targetedOrganizationRequired, targetedSpaceRequired := fakeSharedActor.CheckTargetArgsForCall(0)
			Expect(config).To(Equal(fakeConfig))
			Expect(targetedOrganizationRequired).To(Equal(true))
			Expect(targetedSpaceRequired).To(Equal(true))

			Expect(executeErr).To(MatchError(translatableerror.NotLoggedInError{BinaryName: binaryName}))
		})
	})

	Context("when the API version is below 2.47.0", func() {
		BeforeEach(func() {
			fakeActor.CloudControllerAPIVersionReturns("2.46.0")
		})

		Context("when the health-check-type 'process' is specified", func() {
			BeforeEach(func() {
				cmd.RequiredArgs.HealthCheck.Type = "process"
			})

			It("returns the UnsupportedHealthCheckTypeError", func() {
				Expect(executeErr).To(MatchError(translatableerror.HealthCheckTypeUnsupportedError{
					SupportedTypes: []string{"port", "none"},
				}))
			})
		})

		Context("when the health-check-type 'http' is specified", func() {
			BeforeEach(func() {
				cmd.RequiredArgs.HealthCheck.Type = "http"
			})

			It("returns the UnsupportedHealthCheckTypeError", func() {
				Expect(executeErr).To(MatchError(translatableerror.HealthCheckTypeUnsupportedError{
					SupportedTypes: []string{"port", "none"},
				}))
			})
		})

		Context("when a valid health-check-type is specified", func() {
			BeforeEach(func() {
				cmd.RequiredArgs.HealthCheck.Type = "port"
			})

			It("does not error", func() {
				Expect(executeErr).ToNot(HaveOccurred())
			})
		})
	})

	Context("when the API version is below 2.68.0", func() {
		BeforeEach(func() {
			fakeActor.CloudControllerAPIVersionReturns("2.67.0")
		})

		Context("when the health-check-type 'http' is specified", func() {
			BeforeEach(func() {
				cmd.RequiredArgs.HealthCheck.Type = "http"
			})

			It("returns the UnsupportedHealthCheckTypeError", func() {
				Expect(executeErr).To(MatchError(translatableerror.HealthCheckTypeUnsupportedError{
					SupportedTypes: []string{"port", "none", "process"},
				}))
			})
		})

		Context("when a valid health-check-type is specified", func() {
			BeforeEach(func() {
				cmd.RequiredArgs.HealthCheck.Type = "process"
			})

			It("does not error", func() {
				Expect(executeErr).ToNot(HaveOccurred())
			})
		})
	})

	Context("when setting the application health check type returns an error", func() {
		var expectedErr error

		BeforeEach(func() {
			cmd.RequiredArgs.AppName = "some-app"
			cmd.RequiredArgs.HealthCheck.Type = "some-health-check-type"

			expectedErr = errors.New("set health check error")
			fakeActor.SetApplicationHealthCheckTypeByNameAndSpaceReturns(
				v2action.Application{}, v2action.Warnings{"warning-1"}, expectedErr)
		})

		It("displays warnings and returns the error", func() {
			Expect(testUI.Err).To(Say("warning-1"))
			Expect(executeErr).To(MatchError(expectedErr))
		})
	})

	Context("when setting health check is successful", func() {
		BeforeEach(func() {
			cmd.RequiredArgs.AppName = "some-app"
			cmd.RequiredArgs.HealthCheck.Type = "some-health-check-type"
			cmd.HTTPEndpoint = "/"

			fakeActor.SetApplicationHealthCheckTypeByNameAndSpaceReturns(
				v2action.Application{}, v2action.Warnings{"warning-1"}, nil)
		})

		It("informs the user and displays warnings", func() {
			Expect(testUI.Out).To(Say("Updating health check type for app some-app in org some-org / space some-space as some-user..."))
			Expect(testUI.Err).To(Say("warning-1"))
			Expect(testUI.Out).To(Say("OK"))
			Expect(executeErr).ToNot(HaveOccurred())

			Expect(fakeActor.SetApplicationHealthCheckTypeByNameAndSpaceCallCount()).To(Equal(1))
			name, spaceGUID, healthCheckType, healthCheckHTTPEndpoint := fakeActor.SetApplicationHealthCheckTypeByNameAndSpaceArgsForCall(0)
			Expect(name).To(Equal("some-app"))
			Expect(spaceGUID).To(Equal("some-space-guid"))
			Expect(healthCheckType).To(Equal(v2action.ApplicationHealthCheckType("some-health-check-type")))
			Expect(healthCheckHTTPEndpoint).To(Equal("/"))
		})
	})

	Context("when the app is started", func() {
		BeforeEach(func() {
			cmd.RequiredArgs.AppName = "some-app"
			cmd.RequiredArgs.HealthCheck.Type = "some-health-check-type"

			fakeActor.SetApplicationHealthCheckTypeByNameAndSpaceReturns(
				v2action.Application{State: ccv2.ApplicationStarted}, v2action.Warnings{"warning-1"}, nil)
		})

		It("displays a tip to restart the app", func() {
			Expect(testUI.Out).To(Say("TIP: An app restart is required for the change to take affect."))
		})
	})

})
