package v3_test

import (
	"github.com/liamawhite/cli-with-i18n/actor/cfnetworkingaction"
	"github.com/liamawhite/cli-with-i18n/actor/sharedaction"
	"github.com/liamawhite/cli-with-i18n/actor/v3action"
	"github.com/liamawhite/cli-with-i18n/command/commandfakes"
	"github.com/liamawhite/cli-with-i18n/command/flag"
	"github.com/liamawhite/cli-with-i18n/command/translatableerror"
	. "github.com/liamawhite/cli-with-i18n/command/v3"
	"github.com/liamawhite/cli-with-i18n/command/v3/v3fakes"
	"github.com/liamawhite/cli-with-i18n/util/configv3"
	"github.com/liamawhite/cli-with-i18n/util/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("remove-network-policy Command", func() {
	var (
		cmd             RemoveNetworkPolicyCommand
		testUI          *ui.UI
		fakeConfig      *commandfakes.FakeConfig
		fakeSharedActor *commandfakes.FakeSharedActor
		fakeActor       *v3fakes.FakeRemoveNetworkPolicyActor
		binaryName      string
		executeErr      error
		srcApp          string
		destApp         string
		protocol        string
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(nil, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeSharedActor = new(commandfakes.FakeSharedActor)
		fakeActor = new(v3fakes.FakeRemoveNetworkPolicyActor)

		srcApp = "some-app"
		destApp = "some-other-app"
		protocol = "tcp"

		cmd = RemoveNetworkPolicyCommand{
			UI:             testUI,
			Config:         fakeConfig,
			SharedActor:    fakeSharedActor,
			Actor:          fakeActor,
			RequiredArgs:   flag.RemoveNetworkPolicyArgs{SourceApp: srcApp},
			DestinationApp: destApp,
			Protocol:       flag.NetworkProtocol{Protocol: protocol},
			Port:           flag.NetworkPort{StartPort: 8080, EndPort: 8081},
		}

		binaryName = "faceman"
		fakeConfig.BinaryNameReturns(binaryName)
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	Context("when checking target fails", func() {
		BeforeEach(func() {
			fakeSharedActor.CheckTargetReturns(sharedaction.NotLoggedInError{BinaryName: binaryName})
		})

		It("returns an error", func() {
			Expect(executeErr).To(MatchError(translatableerror.NotLoggedInError{BinaryName: binaryName}))

			Expect(fakeSharedActor.CheckTargetCallCount()).To(Equal(1))
			passedConfig, checkTargetedOrg, checkTargetedSpace := fakeSharedActor.CheckTargetArgsForCall(0)
			Expect(passedConfig).To(Equal(fakeConfig))
			Expect(checkTargetedOrg).To(BeTrue())
			Expect(checkTargetedSpace).To(BeTrue())
		})
	})

	Context("when the user is logged in", func() {
		BeforeEach(func() {
			fakeConfig.CurrentUserReturns(configv3.User{Name: "some-user"}, nil)
			fakeConfig.TargetedSpaceReturns(configv3.Space{Name: "some-space", GUID: "some-space-guid"})
			fakeConfig.TargetedOrganizationReturns(configv3.Organization{Name: "some-org"})
		})

		It("outputs flavor text", func() {
			Expect(testUI.Out).To(Say(`Removing network policy for app %s in org some-org / space some-space as some-user\.\.\.`, srcApp))
		})

		Context("when the policy deletion is successful", func() {
			BeforeEach(func() {
				fakeActor.RemoveNetworkPolicyReturns(cfnetworkingaction.Warnings{"some-warning-1", "some-warning-2"}, nil)
			})

			It("displays OK when no error occurs", func() {
				Expect(executeErr).ToNot(HaveOccurred())
				Expect(fakeActor.RemoveNetworkPolicyCallCount()).To(Equal(1))
				passedSpaceGuid, passedSrcAppName, passedDestAppName, passedProtocol, passedStartPort, passedEndPort := fakeActor.RemoveNetworkPolicyArgsForCall(0)
				Expect(passedSpaceGuid).To(Equal("some-space-guid"))
				Expect(passedSrcAppName).To(Equal("some-app"))
				Expect(passedDestAppName).To(Equal("some-other-app"))
				Expect(passedProtocol).To(Equal("tcp"))
				Expect(passedStartPort).To(Equal(8080))
				Expect(passedEndPort).To(Equal(8081))

				Expect(testUI.Out).To(Say(`Removing network policy for app %s in org some-org / space some-space as some-user\.\.\.`, srcApp))
				Expect(testUI.Err).To(Say("some-warning-1"))
				Expect(testUI.Err).To(Say("some-warning-2"))
				Expect(testUI.Out).To(Say("OK"))
			})
		})

		Context("when the policy does not exist", func() {
			BeforeEach(func() {
				fakeActor.RemoveNetworkPolicyReturns(cfnetworkingaction.Warnings{"some-warning-1", "some-warning-2"}, cfnetworkingaction.PolicyDoesNotExistError{})
			})

			It("displays OK when no error occurs", func() {
				Expect(executeErr).ToNot(HaveOccurred())
				Expect(fakeActor.RemoveNetworkPolicyCallCount()).To(Equal(1))
				passedSpaceGuid, passedSrcAppName, passedDestAppName, passedProtocol, passedStartPort, passedEndPort := fakeActor.RemoveNetworkPolicyArgsForCall(0)
				Expect(passedSpaceGuid).To(Equal("some-space-guid"))
				Expect(passedSrcAppName).To(Equal("some-app"))
				Expect(passedDestAppName).To(Equal("some-other-app"))
				Expect(passedProtocol).To(Equal("tcp"))
				Expect(passedStartPort).To(Equal(8080))
				Expect(passedEndPort).To(Equal(8081))

				Expect(testUI.Out).To(Say(`Removing network policy for app %s in org some-org / space some-space as some-user\.\.\.`, srcApp))
				Expect(testUI.Err).To(Say("some-warning-1"))
				Expect(testUI.Err).To(Say("some-warning-2"))
				Expect(testUI.Out).To(Say("Policy does not exist."))
				Expect(testUI.Out).To(Say("OK"))
			})
		})

		Context("when the policy deletion is not successful", func() {
			BeforeEach(func() {
				fakeActor.RemoveNetworkPolicyReturns(cfnetworkingaction.Warnings{"some-warning-1", "some-warning-2"}, v3action.ApplicationNotFoundError{Name: srcApp})
			})

			It("does not display OK when an error occurs", func() {
				Expect(executeErr).To(MatchError(translatableerror.ApplicationNotFoundError{Name: srcApp}))

				Expect(testUI.Out).To(Say(`Removing network policy for app %s in org some-org / space some-space as some-user\.\.\.`, srcApp))
				Expect(testUI.Err).To(Say("some-warning-1"))
				Expect(testUI.Err).To(Say("some-warning-2"))
				Expect(testUI.Out).ToNot(Say("OK"))
			})
		})
	})
})
