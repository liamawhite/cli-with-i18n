package v3_test

import (
	"errors"
	"time"

	"github.com/liamawhite/cli-with-i18n/actor/sharedaction"
	"github.com/liamawhite/cli-with-i18n/actor/v3action"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccerror"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccversion"
	"github.com/liamawhite/cli-with-i18n/command/commandfakes"
	"github.com/liamawhite/cli-with-i18n/command/flag"
	"github.com/liamawhite/cli-with-i18n/command/translatableerror"
	"github.com/liamawhite/cli-with-i18n/command/v3"
	"github.com/liamawhite/cli-with-i18n/command/v3/v3fakes"
	"github.com/liamawhite/cli-with-i18n/util/configv3"
	"github.com/liamawhite/cli-with-i18n/util/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("v3-packages Command", func() {
	var (
		cmd             v3.V3PackagesCommand
		testUI          *ui.UI
		fakeConfig      *commandfakes.FakeConfig
		fakeSharedActor *commandfakes.FakeSharedActor
		fakeActor       *v3fakes.FakeV3PackagesActor
		binaryName      string
		executeErr      error
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(nil, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeSharedActor = new(commandfakes.FakeSharedActor)
		fakeActor = new(v3fakes.FakeV3PackagesActor)

		binaryName = "faceman"
		fakeConfig.BinaryNameReturns(binaryName)

		cmd = v3.V3PackagesCommand{
			RequiredArgs: flag.AppName{AppName: "some-app"},
			UI:           testUI,
			Config:       fakeConfig,
			Actor:        fakeActor,
			SharedActor:  fakeSharedActor,
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

	Context("when getting the application packages returns an error", func() {
		var expectedErr error

		BeforeEach(func() {
			expectedErr = ccerror.RequestError{}
			fakeActor.GetApplicationPackagesReturns([]v3action.Package{}, v3action.Warnings{"warning-1", "warning-2"}, expectedErr)
		})

		It("returns the error and prints warnings", func() {
			Expect(executeErr).To(Equal(translatableerror.APIRequestError{}))

			Expect(testUI.Out).To(Say("Listing packages of app some-app in org some-org / space some-space as steve\\.\\.\\."))

			Expect(testUI.Err).To(Say("warning-1"))
			Expect(testUI.Err).To(Say("warning-2"))
		})
	})

	Context("when getting the application packages returns some packages", func() {
		var package1UTC, package2UTC string

		BeforeEach(func() {
			package1UTC = "2017-08-14T21:16:42Z"
			package2UTC = "2017-08-16T00:18:24Z"

			packages := []v3action.Package{
				{
					GUID:      "some-package-guid-1",
					State:     "READY",
					CreatedAt: package1UTC,
				},
				{
					GUID:      "some-package-guid-2",
					State:     "FAILED",
					CreatedAt: package2UTC,
				},
			}
			fakeActor.GetApplicationPackagesReturns(packages, v3action.Warnings{"warning-1", "warning-2"}, nil)
		})

		It("prints the application packages and outputs warnings", func() {
			Expect(executeErr).ToNot(HaveOccurred())

			Expect(testUI.Out).To(Say("Listing packages of app some-app in org some-org / space some-space as steve\\.\\.\\."))

			Expect(testUI.Out).To(Say("guid\\s+state\\s+created"))
			package1UTCTime, err := time.Parse(time.RFC3339, package1UTC)
			Expect(err).ToNot(HaveOccurred())
			package2UTCTime, err := time.Parse(time.RFC3339, package2UTC)
			Expect(err).ToNot(HaveOccurred())
			Expect(testUI.Out).To(Say("some-package-guid-1\\s+ready\\s+%s", testUI.UserFriendlyDate(package1UTCTime)))
			Expect(testUI.Out).To(Say("some-package-guid-2\\s+failed\\s+%s", testUI.UserFriendlyDate(package2UTCTime)))

			Expect(testUI.Err).To(Say("warning-1"))
			Expect(testUI.Err).To(Say("warning-2"))

			Expect(fakeActor.GetApplicationPackagesCallCount()).To(Equal(1))
			appName, spaceGUID := fakeActor.GetApplicationPackagesArgsForCall(0)
			Expect(appName).To(Equal("some-app"))
			Expect(spaceGUID).To(Equal("some-space-guid"))
		})
	})

	Context("when getting the application packages returns no packages", func() {
		BeforeEach(func() {
			fakeActor.GetApplicationPackagesReturns([]v3action.Package{}, v3action.Warnings{"warning-1", "warning-2"}, nil)
		})

		It("displays there are no packages", func() {
			Expect(executeErr).ToNot(HaveOccurred())

			Expect(testUI.Out).To(Say("Listing packages of app some-app in org some-org / space some-space as steve\\.\\.\\."))
			Expect(testUI.Out).To(Say("No packages found"))

			Expect(testUI.Err).To(Say("warning-1"))
			Expect(testUI.Err).To(Say("warning-2"))
		})
	})
})
