package v3action_test

import (
	"errors"
	"net/url"

	. "github.com/liamawhite/cli-with-i18n/actor/v3action"
	"github.com/liamawhite/cli-with-i18n/actor/v3action/v3actionfakes"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccv3"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Application with ProcessSummary Actions", func() {
	var (
		actor                     *Actor
		fakeCloudControllerClient *v3actionfakes.FakeCloudControllerClient
	)

	BeforeEach(func() {
		fakeCloudControllerClient = new(v3actionfakes.FakeCloudControllerClient)
		actor = NewActor(nil, fakeCloudControllerClient, nil)
	})

	Describe("GetApplicationsWithProcessesBySpace", func() {
		Context("when there are apps", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.GetApplicationsReturns(
					[]ccv3.Application{
						{
							Name:  "some-app-name-1",
							GUID:  "some-app-guid-1",
							State: "RUNNING",
						},
						{
							Name:  "some-app-name-2",
							GUID:  "some-app-guid-2",
							State: "STOPPED",
						},
					},
					ccv3.Warnings{"some-warning"},
					nil,
				)

				fakeCloudControllerClient.GetApplicationProcessesReturnsOnCall(
					0,
					[]ccv3.Process{
						{
							GUID: "some-process-guid-1",
							Type: "some-process-type-1",
						},
						{
							GUID: "some-process-guid-2",
							Type: "some-process-type-2",
						},
					},
					ccv3.Warnings{"some-process-warning-1"},
					nil,
				)
				fakeCloudControllerClient.GetApplicationProcessesReturnsOnCall(
					1,
					[]ccv3.Process{
						{
							GUID: "some-process-guid-3",
							Type: "some-process-type-3",
						},
					},
					ccv3.Warnings{"some-process-warning-2"},
					nil,
				)

				fakeCloudControllerClient.GetProcessInstancesReturnsOnCall(
					0,
					[]ccv3.Instance{{State: "RUNNING"}, {State: "DOWN"}, {State: "RUNNING"}},
					ccv3.Warnings{"some-process-stats-warning-1"},
					nil,
				)
				fakeCloudControllerClient.GetProcessInstancesReturnsOnCall(
					1,
					[]ccv3.Instance{{State: "RUNNING"}, {State: "RUNNING"}},
					ccv3.Warnings{"some-process-stats-warning-2"},
					nil,
				)
				fakeCloudControllerClient.GetProcessInstancesReturnsOnCall(
					2,
					[]ccv3.Instance{{State: "DOWN"}},
					ccv3.Warnings{"some-process-stats-warning-3"},
					nil,
				)
			})

			It("returns app summaries and warnings", func() {
				summaries, warnings, err := actor.GetApplicationsWithProcessesBySpace("some-space-guid")
				Expect(err).ToNot(HaveOccurred())
				Expect(summaries).To(Equal([]ApplicationWithProcessSummary{
					{
						Application: Application{
							Name:  "some-app-name-1",
							GUID:  "some-app-guid-1",
							State: "RUNNING",
						},
						ProcessSummaries: []ProcessSummary{
							{
								Process:         Process{GUID: "some-process-guid-1", Type: "some-process-type-1"},
								InstanceDetails: []Instance{{State: "RUNNING"}, {State: "DOWN"}, {State: "RUNNING"}},
							},
							{
								Process:         Process{GUID: "some-process-guid-2", Type: "some-process-type-2"},
								InstanceDetails: []Instance{{State: "RUNNING"}, {State: "RUNNING"}},
							},
						},
					},
					{
						Application: Application{
							Name:  "some-app-name-2",
							GUID:  "some-app-guid-2",
							State: "STOPPED",
						},
						ProcessSummaries: []ProcessSummary{
							{
								Process:         Process{GUID: "some-process-guid-3", Type: "some-process-type-3"},
								InstanceDetails: []Instance{{State: "DOWN"}},
							},
						},
					},
				}))
				Expect(warnings).To(Equal(Warnings{"some-warning", "some-process-warning-1", "some-process-stats-warning-1", "some-process-stats-warning-2", "some-process-warning-2", "some-process-stats-warning-3"}))

				Expect(fakeCloudControllerClient.GetApplicationsCallCount()).To(Equal(1))
				expectedQuery := url.Values{
					"space_guids": []string{"some-space-guid"},
					"order_by":    []string{"name"},
				}
				query := fakeCloudControllerClient.GetApplicationsArgsForCall(0)
				Expect(query).To(Equal(expectedQuery))

				Expect(fakeCloudControllerClient.GetApplicationProcessesCallCount()).To(Equal(2))
				appGUID := fakeCloudControllerClient.GetApplicationProcessesArgsForCall(0)
				Expect(appGUID).To(Equal("some-app-guid-1"))
				appGUID = fakeCloudControllerClient.GetApplicationProcessesArgsForCall(1)
				Expect(appGUID).To(Equal("some-app-guid-2"))

				Expect(fakeCloudControllerClient.GetProcessInstancesCallCount()).To(Equal(3))
				processGUID := fakeCloudControllerClient.GetProcessInstancesArgsForCall(0)
				Expect(processGUID).To(Equal("some-process-guid-1"))
				processGUID = fakeCloudControllerClient.GetProcessInstancesArgsForCall(1)
				Expect(processGUID).To(Equal("some-process-guid-2"))
				processGUID = fakeCloudControllerClient.GetProcessInstancesArgsForCall(2)
				Expect(processGUID).To(Equal("some-process-guid-3"))
			})
		})

		Context("when getting the app processes returns an error", func() {
			var expectedErr error

			BeforeEach(func() {
				fakeCloudControllerClient.GetApplicationsReturns(
					[]ccv3.Application{
						{
							Name:  "some-app-name",
							GUID:  "some-app-guid",
							State: "RUNNING",
						},
					},
					ccv3.Warnings{"some-warning"},
					nil,
				)

				expectedErr = errors.New("some error")
				fakeCloudControllerClient.GetApplicationProcessesReturns(
					[]ccv3.Process{},
					ccv3.Warnings{"some-process-warning"},
					expectedErr,
				)
			})

			It("returns the error", func() {
				_, warnings, err := actor.GetApplicationsWithProcessesBySpace("some-space-guid")
				Expect(err).To(Equal(expectedErr))
				Expect(warnings).To(Equal(Warnings{"some-warning", "some-process-warning"}))
			})
		})

		Context("when getting the app process instances returns an error", func() {
			var expectedErr error

			BeforeEach(func() {
				fakeCloudControllerClient.GetApplicationsReturns(
					[]ccv3.Application{
						{
							Name:  "some-app-name",
							GUID:  "some-app-guid",
							State: "RUNNING",
						},
					},
					ccv3.Warnings{"some-warning"},
					nil,
				)

				fakeCloudControllerClient.GetApplicationProcessesReturns(
					[]ccv3.Process{
						{
							GUID: "some-process-guid",
							Type: "some-type",
						},
					},
					ccv3.Warnings{"some-process-warning"},
					nil,
				)

				expectedErr = errors.New("some error")
				fakeCloudControllerClient.GetProcessInstancesReturns(
					[]ccv3.Instance{},
					ccv3.Warnings{"some-process-stats-warning"},
					expectedErr,
				)
			})

			It("returns the error", func() {
				_, warnings, err := actor.GetApplicationsWithProcessesBySpace("some-space-guid")
				Expect(err).To(Equal(expectedErr))
				Expect(warnings).To(Equal(Warnings{"some-warning", "some-process-warning", "some-process-stats-warning"}))
			})
		})
	})
})
