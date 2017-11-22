package v2_test

import (
	"errors"

	"github.com/liamawhite/cli-with-i18n/actor/sharedaction"
	"github.com/liamawhite/cli-with-i18n/actor/v2action"
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

var _ = Describe("orgs Command", func() {
	var (
		cmd             OrgsCommand
		testUI          *ui.UI
		fakeConfig      *commandfakes.FakeConfig
		fakeSharedActor *commandfakes.FakeSharedActor
		fakeActor       *v2fakes.FakeOrgsActor
		binaryName      string
		executeErr      error
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(nil, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeSharedActor = new(commandfakes.FakeSharedActor)
		fakeActor = new(v2fakes.FakeOrgsActor)

		cmd = OrgsCommand{
			UI:          testUI,
			Config:      fakeConfig,
			SharedActor: fakeSharedActor,
			Actor:       fakeActor,
		}

		binaryName = "faceman"
		fakeConfig.BinaryNameReturns(binaryName)
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	Context("when an error is encountered checking if the environment is setup correctly", func() {
		BeforeEach(func() {
			fakeSharedActor.CheckTargetReturns(sharedaction.NotLoggedInError{BinaryName: binaryName})
		})

		It("returns an error", func() {
			Expect(executeErr).To(MatchError(translatableerror.NotLoggedInError{BinaryName: binaryName}))

			Expect(fakeSharedActor.CheckTargetCallCount()).To(Equal(1))
			_, checkTargetedOrgArg, checkTargetedSpaceArg := fakeSharedActor.CheckTargetArgsForCall(0)
			Expect(checkTargetedOrgArg).To(BeFalse())
			Expect(checkTargetedSpaceArg).To(BeFalse())
		})
	})

	Context("when the user is logged in and an org is targeted", func() {
		Context("when getting the current user fails", func() {
			BeforeEach(func() {
				fakeConfig.CurrentUserReturns(configv3.User{}, errors.New("get-user-error"))
			})

			It("returns the error", func() {
				Expect(executeErr).To(MatchError("get-user-error"))
			})
		})

		Context("when getting the current user succeeds", func() {
			BeforeEach(func() {
				fakeConfig.CurrentUserReturns(
					configv3.User{Name: "some-user"},
					nil)
			})

			Context("when there are no orgs", func() {
				BeforeEach(func() {
					fakeActor.GetOrganizationsReturns(
						[]v2action.Organization{},
						v2action.Warnings{"get-orgs-warning"},
						nil)
				})

				It("displays that there are no orgs", func() {
					Expect(executeErr).ToNot(HaveOccurred())

					Expect(testUI.Out).To(Say("Getting orgs as some-user\\.\\.\\."))
					Expect(testUI.Out).To(Say(""))
					Expect(testUI.Out).To(Say("No orgs found\\."))

					Expect(testUI.Err).To(Say("get-orgs-warning"))

					Expect(fakeActor.GetOrganizationsCallCount()).To(Equal(1))
				})
			})

			Context("when there are multiple orgs", func() {
				BeforeEach(func() {
					fakeActor.GetOrganizationsReturns(
						[]v2action.Organization{
							{Name: "org-1"},
							{Name: "org-2"},
						},
						v2action.Warnings{"get-orgs-warning"},
						nil)
				})

				It("displays all the orgs in the org", func() {
					Expect(executeErr).ToNot(HaveOccurred())

					Expect(testUI.Out).To(Say("Getting orgs as some-user\\.\\.\\."))
					Expect(testUI.Out).To(Say(""))
					Expect(testUI.Out).To(Say("name"))
					Expect(testUI.Out).To(Say("org-1"))
					Expect(testUI.Out).To(Say("org-2"))

					Expect(testUI.Err).To(Say("get-orgs-warning"))

					Expect(fakeActor.GetOrganizationsCallCount()).To(Equal(1))
				})
			})

			Context("when a translatable error is encountered getting orgs", func() {
				BeforeEach(func() {
					fakeActor.GetOrganizationsReturns(
						nil,
						v2action.Warnings{"get-orgs-warning"},
						v2action.OrganizationNotFoundError{Name: "not-found-org"})
				})

				It("returns a translatable error", func() {
					Expect(executeErr).To(MatchError(translatableerror.OrganizationNotFoundError{Name: "not-found-org"}))

					Expect(testUI.Out).To(Say("Getting orgs as some-user\\.\\.\\."))
					Expect(testUI.Out).To(Say(""))

					Expect(testUI.Err).To(Say("get-orgs-warning"))

					Expect(fakeActor.GetOrganizationsCallCount()).To(Equal(1))
				})
			})
		})
	})
})
