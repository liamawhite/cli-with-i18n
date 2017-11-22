package v3action_test

import (
	"errors"
	"net/url"

	. "github.com/liamawhite/cli-with-i18n/actor/v3action"
	"github.com/liamawhite/cli-with-i18n/actor/v3action/v3actionfakes"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccerror"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccv3"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Droplet Actions", func() {
	var (
		actor                     *Actor
		fakeCloudControllerClient *v3actionfakes.FakeCloudControllerClient
	)

	BeforeEach(func() {
		fakeCloudControllerClient = new(v3actionfakes.FakeCloudControllerClient)
		actor = NewActor(nil, fakeCloudControllerClient, nil)
	})

	Describe("SetApplicationDroplet", func() {
		Context("when there are no client errors", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.GetApplicationsReturns(
					[]ccv3.Application{
						{GUID: "some-app-guid"},
					},
					ccv3.Warnings{"get-applications-warning"},
					nil,
				)

				fakeCloudControllerClient.SetApplicationDropletReturns(
					ccv3.Relationship{GUID: "some-droplet-guid"},
					ccv3.Warnings{"set-application-droplet-warning"},
					nil,
				)
			})

			It("sets the app's droplet", func() {
				warnings, err := actor.SetApplicationDroplet("some-app-name", "some-space-guid", "some-droplet-guid")

				Expect(err).ToNot(HaveOccurred())
				Expect(warnings).To(ConsistOf("get-applications-warning", "set-application-droplet-warning"))

				Expect(fakeCloudControllerClient.GetApplicationsCallCount()).To(Equal(1))
				queryURL := fakeCloudControllerClient.GetApplicationsArgsForCall(0)
				query := url.Values{"names": []string{"some-app-name"}, "space_guids": []string{"some-space-guid"}}
				Expect(queryURL).To(Equal(query))

				Expect(fakeCloudControllerClient.SetApplicationDropletCallCount()).To(Equal(1))
				appGUID, dropletGUID := fakeCloudControllerClient.SetApplicationDropletArgsForCall(0)
				Expect(appGUID).To(Equal("some-app-guid"))
				Expect(dropletGUID).To(Equal("some-droplet-guid"))
			})
		})

		Context("when getting the application fails", func() {
			var expectedErr error

			BeforeEach(func() {
				expectedErr = errors.New("some get application error")

				fakeCloudControllerClient.GetApplicationsReturns(
					[]ccv3.Application{},
					ccv3.Warnings{"get-applications-warning"},
					expectedErr,
				)
			})

			It("returns the error", func() {
				warnings, err := actor.SetApplicationDroplet("some-app-name", "some-space-guid", "some-droplet-guid")

				Expect(err).To(Equal(expectedErr))
				Expect(warnings).To(ConsistOf("get-applications-warning"))
			})
		})

		Context("when setting the droplet fails", func() {
			var expectedErr error
			BeforeEach(func() {
				expectedErr = errors.New("some set application-droplet error")
				fakeCloudControllerClient.GetApplicationsReturns(
					[]ccv3.Application{
						{GUID: "some-app-guid"},
					},
					ccv3.Warnings{"get-applications-warning"},
					nil,
				)

				fakeCloudControllerClient.SetApplicationDropletReturns(
					ccv3.Relationship{},
					ccv3.Warnings{"set-application-droplet-warning"},
					expectedErr,
				)
			})

			It("returns the error", func() {
				warnings, err := actor.SetApplicationDroplet("some-app-name", "some-space-guid", "some-droplet-guid")

				Expect(err).To(Equal(expectedErr))
				Expect(warnings).To(ConsistOf("get-applications-warning", "set-application-droplet-warning"))
			})

			Context("when the cc client response contains an UnprocessableEntityError", func() {
				BeforeEach(func() {
					fakeCloudControllerClient.SetApplicationDropletReturns(
						ccv3.Relationship{},
						ccv3.Warnings{"set-application-droplet-warning"},
						ccerror.UnprocessableEntityError{Message: "some-message"},
					)
				})

				It("raises the error as AssignDropletError and returns warnings", func() {
					warnings, err := actor.SetApplicationDroplet("some-app-name", "some-space-guid", "some-droplet-guid")

					Expect(err).To(MatchError("some-message"))
					Expect(warnings).To(ConsistOf("get-applications-warning", "set-application-droplet-warning"))
				})
			})

		})
	})

	Describe("GetApplicationDroplets", func() {
		Context("when there are no client errors", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.GetApplicationsReturns(
					[]ccv3.Application{
						{GUID: "some-app-guid"},
					},
					ccv3.Warnings{"get-applications-warning"},
					nil,
				)

				fakeCloudControllerClient.GetApplicationDropletsReturns(
					[]ccv3.Droplet{
						{
							GUID:      "some-droplet-guid-1",
							State:     ccv3.DropletStateStaged,
							CreatedAt: "2017-08-14T21:16:42Z",
							Buildpacks: []ccv3.DropletBuildpack{
								{Name: "ruby"},
								{Name: "nodejs"},
							},
							Image: "docker/some-image",
							Stack: "penguin",
						},
						{
							GUID:      "some-droplet-guid-2",
							State:     ccv3.DropletStateFailed,
							CreatedAt: "2017-08-16T00:18:24Z",
							Buildpacks: []ccv3.DropletBuildpack{
								{Name: "java"},
							},
							Stack: "windows",
						},
					},
					ccv3.Warnings{"get-application-droplets-warning"},
					nil,
				)
			})

			It("gets the app's droplets", func() {
				droplets, warnings, err := actor.GetApplicationDroplets("some-app-name", "some-space-guid")

				Expect(err).ToNot(HaveOccurred())
				Expect(warnings).To(ConsistOf("get-applications-warning", "get-application-droplets-warning"))
				Expect(droplets).To(Equal([]Droplet{
					{
						GUID:      "some-droplet-guid-1",
						State:     DropletStateStaged,
						CreatedAt: "2017-08-14T21:16:42Z",
						Buildpacks: []Buildpack{
							{Name: "ruby"},
							{Name: "nodejs"},
						},
						Image: "docker/some-image",
						Stack: "penguin",
					},
					{
						GUID:      "some-droplet-guid-2",
						State:     DropletStateFailed,
						CreatedAt: "2017-08-16T00:18:24Z",
						Buildpacks: []Buildpack{
							{Name: "java"},
						},
						Stack: "windows",
					},
				}))

				Expect(fakeCloudControllerClient.GetApplicationsCallCount()).To(Equal(1))
				queryURL := fakeCloudControllerClient.GetApplicationsArgsForCall(0)
				query := url.Values{"names": []string{"some-app-name"}, "space_guids": []string{"some-space-guid"}}
				Expect(queryURL).To(Equal(query))

				Expect(fakeCloudControllerClient.GetApplicationDropletsCallCount()).To(Equal(1))
				appGUID, query := fakeCloudControllerClient.GetApplicationDropletsArgsForCall(0)
				Expect(appGUID).To(Equal("some-app-guid"))
				Expect(query).To(Equal(url.Values{}))
			})
		})

		Context("when getting the application fails", func() {
			var expectedErr error

			BeforeEach(func() {
				expectedErr = errors.New("some get application error")

				fakeCloudControllerClient.GetApplicationsReturns(
					[]ccv3.Application{},
					ccv3.Warnings{"get-applications-warning"},
					expectedErr,
				)
			})

			It("returns the error", func() {
				_, warnings, err := actor.GetApplicationDroplets("some-app-name", "some-space-guid")

				Expect(err).To(Equal(expectedErr))
				Expect(warnings).To(ConsistOf("get-applications-warning"))
			})
		})

		Context("when getting the application droplets fails", func() {
			var expectedErr error

			BeforeEach(func() {
				expectedErr = errors.New("some get application error")

				fakeCloudControllerClient.GetApplicationsReturns(
					[]ccv3.Application{
						{GUID: "some-app-guid"},
					},
					ccv3.Warnings{"get-applications-warning"},
					nil,
				)

				fakeCloudControllerClient.GetApplicationDropletsReturns(
					[]ccv3.Droplet{},
					ccv3.Warnings{"get-application-droplets-warning"},
					expectedErr,
				)
			})

			It("returns the error", func() {
				_, warnings, err := actor.GetApplicationDroplets("some-app-name", "some-space-guid")

				Expect(err).To(Equal(expectedErr))
				Expect(warnings).To(ConsistOf("get-applications-warning", "get-application-droplets-warning"))
			})
		})
	})
})
