package v2_test

import (
	"errors"

	"github.com/liamawhite/cli-with-i18n/actor/sharedaction"
	"github.com/liamawhite/cli-with-i18n/command/commandfakes"
	"github.com/liamawhite/cli-with-i18n/command/translatableerror"
	. "github.com/liamawhite/cli-with-i18n/command/v2"
	"github.com/liamawhite/cli-with-i18n/command/v2/v2fakes"
	"github.com/liamawhite/cli-with-i18n/util/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("ssh-code Command", func() {
	var (
		cmd             SSHCodeCommand
		testUI          *ui.UI
		fakeConfig      *commandfakes.FakeConfig
		fakeSharedActor *commandfakes.FakeSharedActor
		fakeActor       *v2fakes.FakeSSHCodeActor
		binaryName      string
		executeErr      error
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(nil, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeSharedActor = new(commandfakes.FakeSharedActor)
		fakeActor = new(v2fakes.FakeSSHCodeActor)

		cmd = SSHCodeCommand{
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

	Context("when checking the target fails", func() {
		BeforeEach(func() {
			fakeSharedActor.CheckTargetReturns(
				sharedaction.NotLoggedInError{BinaryName: binaryName})
		})

		It("returns an error", func() {
			Expect(executeErr).To(MatchError(translatableerror.NotLoggedInError{BinaryName: binaryName}))

			Expect(fakeSharedActor.CheckTargetCallCount()).To(Equal(1))
			config, targetedOrganizationRequired, targetedSpaceRequired := fakeSharedActor.CheckTargetArgsForCall(0)
			Expect(config).To(Equal(fakeConfig))
			Expect(targetedOrganizationRequired).To(Equal(false))
			Expect(targetedSpaceRequired).To(Equal(false))
		})
	})

	Context("when the user is logged in", func() {
		var code string

		BeforeEach(func() {
			code = "s3curep4ss"
			fakeActor.GetSSHPasscodeReturns(code, nil)
		})

		It("displays the ssh code", func() {
			Expect(executeErr).NotTo(HaveOccurred())
			Expect(testUI.Out).To(Say(code))
			Expect(fakeActor.GetSSHPasscodeCallCount()).To(Equal(1))
		})

		Context("when an error is encountered getting the ssh code", func() {
			var expectedErr error

			BeforeEach(func() {
				expectedErr = errors.New("get ssh code error")
				fakeActor.GetSSHPasscodeReturns("", expectedErr)
			})

			It("returns the error", func() {
				Expect(executeErr).To(MatchError(expectedErr))

				Expect(fakeActor.GetSSHPasscodeCallCount()).To(Equal(1))
			})
		})
	})
})
