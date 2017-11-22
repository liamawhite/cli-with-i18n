package v3_test

import (
	"errors"

	"github.com/liamawhite/cli-with-i18n/actor/sharedaction"
	"github.com/liamawhite/cli-with-i18n/actor/v2action"
	"github.com/liamawhite/cli-with-i18n/actor/v3action"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccerror"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccv3/constant"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccversion"
	"github.com/liamawhite/cli-with-i18n/command/commandfakes"
	"github.com/liamawhite/cli-with-i18n/command/translatableerror"
	"github.com/liamawhite/cli-with-i18n/command/v3"
	"github.com/liamawhite/cli-with-i18n/command/v3/shared/sharedfakes"
	"github.com/liamawhite/cli-with-i18n/command/v3/v3fakes"
	"github.com/liamawhite/cli-with-i18n/util/configv3"
	"github.com/liamawhite/cli-with-i18n/util/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("v3-apps Command", func() {
	var (
		cmd             v3.V3AppsCommand
		testUI          *ui.UI
		fakeConfig      *commandfakes.FakeConfig
		fakeSharedActor *commandfakes.FakeSharedActor
		fakeActor       *v3fakes.FakeV3AppsActor
		fakeV2Actor     *sharedfakes.FakeV2AppRouteActor
		binaryName      string
		executeErr      error
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(nil, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeSharedActor = new(commandfakes.FakeSharedActor)
		fakeActor = new(v3fakes.FakeV3AppsActor)
		fakeV2Actor = new(sharedfakes.FakeV2AppRouteActor)

		binaryName = "faceman"
		fakeConfig.BinaryNameReturns(binaryName)

		cmd = v3.V3AppsCommand{
			UI:              testUI,
			Config:          fakeConfig,
			Actor:           fakeActor,
			V2AppRouteActor: fakeV2Actor,
			SharedActor:     fakeSharedActor,
		}

		fakeConfig.TargetedOrganizationReturns(configv3.Organization{
			Name: "some-org",
			GUID: "some-org-guid",
		})
		fakeConfig.TargetedSpaceReturns(configv3.Space{
			Name: "some-space",
			GUID: "some-space-guid",
		})

		fakeConfig.CurrentUserReturns(configv3.User{Name: "steve"}, nil)
		fakeActor.CloudControllerAPIVersionReturns(ccversion.MinVersionV3)
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	Context("when the API version is below the minimum", func() {
		BeforeEach(func() {
			fakeActor.CloudControllerAPIVersionReturns("0.0.0")
		})

		It("returns a MinimumAPIVersionNotMetError", func() {
			Expect(executeErr).To(MatchError(translatableerror.MinimumAPIVersionNotMetError{
				CurrentVersion: "0.0.0",
				MinimumVersion: ccversion.MinVersionV3,
			}))
		})
	})

	Context("when checking target fails", func() {
		BeforeEach(func() {
			fakeSharedActor.CheckTargetReturns(sharedaction.NoOrganizationTargetedError{BinaryName: binaryName})
		})

		It("returns an error", func() {
			Expect(executeErr).To(MatchError(translatableerror.NoOrganizationTargetedError{BinaryName: binaryName}))

			Expect(fakeSharedActor.CheckTargetCallCount()).To(Equal(1))
			_, checkTargetedOrg, checkTargetedSpace := fakeSharedActor.CheckTargetArgsForCall(0)
			Expect(checkTargetedOrg).To(BeTrue())
			Expect(checkTargetedSpace).To(BeTrue())
		})
	})

	Context("when the user is not logged in", func() {
		var expectedErr error

		BeforeEach(func() {
			expectedErr = errors.New("some current user error")
			fakeConfig.CurrentUserReturns(configv3.User{}, expectedErr)
		})

		It("return an error", func() {
			Expect(executeErr).To(Equal(expectedErr))
		})
	})

	Context("when getting the applications returns an error", func() {
		var expectedErr error

		BeforeEach(func() {
			expectedErr = ccerror.RequestError{}
			fakeActor.GetApplicationsWithProcessesBySpaceReturns([]v3action.ApplicationWithProcessSummary{}, v3action.Warnings{"warning-1", "warning-2"}, expectedErr)
		})

		It("returns the error and prints warnings", func() {
			Expect(executeErr).To(Equal(translatableerror.APIRequestError{}))

			Expect(testUI.Out).To(Say("Getting apps in org some-org / space some-space as steve\\.\\.\\."))

			Expect(testUI.Err).To(Say("warning-1"))
			Expect(testUI.Err).To(Say("warning-2"))
		})
	})

	Context("when getting routes returns an error", func() {
		var expectedErr error

		BeforeEach(func() {
			expectedErr = ccerror.RequestError{}
			fakeActor.GetApplicationsWithProcessesBySpaceReturns([]v3action.ApplicationWithProcessSummary{
				{
					Application: v3action.Application{
						GUID:  "app-guid",
						Name:  "some-app",
						State: "STARTED",
					},
					ProcessSummaries: []v3action.ProcessSummary{{Process: v3action.Process{Type: "process-type"}}},
				},
			}, v3action.Warnings{"warning-1", "warning-2"}, nil)

			fakeV2Actor.GetApplicationRoutesReturns([]v2action.Route{}, v2action.Warnings{"route-warning-1", "route-warning-2"}, expectedErr)
		})

		It("returns the error and prints warnings", func() {
			Expect(executeErr).To(Equal(translatableerror.APIRequestError{}))

			Expect(testUI.Out).To(Say("Getting apps in org some-org / space some-space as steve\\.\\.\\."))

			Expect(testUI.Err).To(Say("warning-1"))
			Expect(testUI.Err).To(Say("warning-2"))
			Expect(testUI.Err).To(Say("route-warning-1"))
			Expect(testUI.Err).To(Say("route-warning-2"))
		})
	})

	Context("when the route actor does not return any errors", func() {
		BeforeEach(func() {
			fakeV2Actor.GetApplicationRoutesStub = func(appGUID string) (v2action.Routes, v2action.Warnings, error) {
				switch appGUID {
				case "app-guid-1":
					return []v2action.Route{
							{
								Host:   "some-app-1",
								Domain: v2action.Domain{Name: "some-other-domain"},
							},
							{
								Host:   "some-app-1",
								Domain: v2action.Domain{Name: "some-domain"},
							},
						},
						v2action.Warnings{"route-warning-1", "route-warning-2"},
						nil
				case "app-guid-2":
					return []v2action.Route{
							{
								Host:   "some-app-2",
								Domain: v2action.Domain{Name: "some-domain"},
							},
						},
						v2action.Warnings{"route-warning-3", "route-warning-4"},
						nil
				default:
					panic("unknown app guid")
				}
			}
		})

		Context("with existing apps", func() {
			BeforeEach(func() {
				appSummaries := []v3action.ApplicationWithProcessSummary{
					{
						Application: v3action.Application{
							GUID:  "app-guid-1",
							Name:  "some-app-1",
							State: "STARTED",
						},
						ProcessSummaries: []v3action.ProcessSummary{
							{
								Process: v3action.Process{
									Type: "console",
								},
								InstanceDetails: []v3action.Instance{},
							},
							{
								Process: v3action.Process{
									Type: "worker",
								},
								InstanceDetails: []v3action.Instance{
									{
										Index: 0,
										State: "DOWN",
									},
								},
							},
							{
								Process: v3action.Process{
									Type: constant.ProcessTypeWeb,
								},
								InstanceDetails: []v3action.Instance{
									v3action.Instance{
										Index: 0,
										State: "RUNNING",
									},
									v3action.Instance{
										Index: 1,
										State: "RUNNING",
									},
								},
							},
						},
					},
					{
						Application: v3action.Application{
							GUID:  "app-guid-2",
							Name:  "some-app-2",
							State: "STOPPED",
						},
						ProcessSummaries: []v3action.ProcessSummary{
							{
								Process: v3action.Process{
									Type: constant.ProcessTypeWeb,
								},
								InstanceDetails: []v3action.Instance{
									v3action.Instance{
										Index: 0,
										State: "DOWN",
									},
									v3action.Instance{
										Index: 1,
										State: "DOWN",
									},
								},
							},
						},
					},
				}
				fakeActor.GetApplicationsWithProcessesBySpaceReturns(appSummaries, v3action.Warnings{"warning-1", "warning-2"}, nil)
			})

			It("prints the application summary and outputs warnings", func() {
				Expect(executeErr).ToNot(HaveOccurred())

				Expect(testUI.Out).To(Say("Getting apps in org some-org / space some-space as steve\\.\\.\\."))

				Expect(testUI.Out).To(Say("name\\s+requested state\\s+processes\\s+routes"))
				Expect(testUI.Out).To(Say("some-app-1\\s+started\\s+web:2/2, console:0/0, worker:0/1\\s+some-app-1.some-other-domain, some-app-1.some-domain"))
				Expect(testUI.Out).To(Say("some-app-2\\s+stopped\\s+web:0/2\\s+some-app-2.some-domain"))

				Expect(testUI.Err).To(Say("warning-1"))
				Expect(testUI.Err).To(Say("warning-2"))

				Expect(testUI.Err).To(Say("route-warning-1"))
				Expect(testUI.Err).To(Say("route-warning-2"))
				Expect(testUI.Err).To(Say("route-warning-3"))
				Expect(testUI.Err).To(Say("route-warning-4"))

				Expect(fakeActor.GetApplicationsWithProcessesBySpaceCallCount()).To(Equal(1))
				spaceGUID := fakeActor.GetApplicationsWithProcessesBySpaceArgsForCall(0)
				Expect(spaceGUID).To(Equal("some-space-guid"))

				Expect(fakeV2Actor.GetApplicationRoutesCallCount()).To(Equal(2))
				appGUID := fakeV2Actor.GetApplicationRoutesArgsForCall(0)
				Expect(appGUID).To(Equal("app-guid-1"))
				appGUID = fakeV2Actor.GetApplicationRoutesArgsForCall(1)
				Expect(appGUID).To(Equal("app-guid-2"))
			})
		})

		Context("when app does not have processes", func() {
			BeforeEach(func() {
				appSummaries := []v3action.ApplicationWithProcessSummary{
					{
						Application: v3action.Application{
							GUID:  "app-guid",
							Name:  "some-app",
							State: "STARTED",
						},
						ProcessSummaries: []v3action.ProcessSummary{},
					},
				}
				fakeActor.GetApplicationsWithProcessesBySpaceReturns(appSummaries, v3action.Warnings{"warning"}, nil)
			})

			It("it does not request or display routes information for app", func() {
				Expect(executeErr).ToNot(HaveOccurred())

				Expect(testUI.Out).To(Say("Getting apps in org some-org / space some-space as steve\\.\\.\\."))

				Expect(testUI.Out).To(Say("name\\s+requested state\\s+processes\\s+routes"))
				Expect(testUI.Out).To(Say("some-app\\s+started\\s+$"))
				Expect(testUI.Err).To(Say("warning"))

				Expect(fakeActor.GetApplicationsWithProcessesBySpaceCallCount()).To(Equal(1))
				spaceGUID := fakeActor.GetApplicationsWithProcessesBySpaceArgsForCall(0)
				Expect(spaceGUID).To(Equal("some-space-guid"))

				Expect(fakeV2Actor.GetApplicationRoutesCallCount()).To(Equal(0))
			})
		})

		Context("with no apps", func() {
			BeforeEach(func() {
				fakeActor.GetApplicationsWithProcessesBySpaceReturns([]v3action.ApplicationWithProcessSummary{}, v3action.Warnings{"warning-1", "warning-2"}, nil)
			})

			It("displays there are no apps", func() {
				Expect(executeErr).ToNot(HaveOccurred())

				Expect(testUI.Out).To(Say("Getting apps in org some-org / space some-space as steve\\.\\.\\."))
				Expect(testUI.Out).To(Say("No apps found"))
			})
		})
	})
})
