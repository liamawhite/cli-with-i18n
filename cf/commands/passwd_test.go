package commands_test

import (
	"github.com/liamawhite/cli-with-i18n/cf/api/apifakes"
	"github.com/liamawhite/cli-with-i18n/cf/commandregistry"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/models"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"
	"github.com/liamawhite/cli-with-i18n/cf/requirements/requirementsfakes"
	testcmd "github.com/liamawhite/cli-with-i18n/util/testhelpers/commands"
	testconfig "github.com/liamawhite/cli-with-i18n/util/testhelpers/configuration"
	testterm "github.com/liamawhite/cli-with-i18n/util/testhelpers/terminal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/liamawhite/cli-with-i18n/util/testhelpers/matchers"
)

var _ = Describe("password command", func() {
	var (
		pwDeps passwordDeps
		ui     *testterm.FakeUI
		deps   commandregistry.Dependency
	)

	updateCommandDependency := func(pluginCall bool) {
		deps.UI = ui
		deps.Config = pwDeps.Config
		deps.RepoLocator = deps.RepoLocator.SetPasswordRepository(pwDeps.PwdRepo)
		commandregistry.Commands.SetCommand(commandregistry.Commands.FindCommand("passwd").SetDependency(deps, pluginCall))
	}

	callPassword := func(inputs []string, pwDeps passwordDeps) (*testterm.FakeUI, bool) {
		ui = &testterm.FakeUI{Inputs: inputs}
		passed := testcmd.RunCLICommand("passwd", []string{}, pwDeps.ReqFactory, updateCommandDependency, false, ui)
		return ui, passed
	}

	BeforeEach(func() {
		pwDeps = getPasswordDeps()
	})

	It("does not pass requirements if you are not logged in", func() {
		pwDeps.ReqFactory.NewLoginRequirementReturns(requirements.Failing{Message: "not logged in"})
		_, passed := callPassword([]string{}, pwDeps)
		Expect(passed).To(BeFalse())
	})

	Context("when logged in successfully", func() {
		BeforeEach(func() {
			pwDeps.ReqFactory.NewLoginRequirementReturns(requirements.Passing{})
			pwDeps.PwdRepo.UpdateUnauthorized = false
		})

		It("passes requirements", func() {
			_, passed := callPassword([]string{"", "", ""}, pwDeps)
			Expect(passed).To(BeTrue())
		})

		It("changes your password when given a new password", func() {
			pwDeps.PwdRepo.UpdateUnauthorized = false
			ui, _ := callPassword([]string{"old-password", "new-password", "new-password"}, pwDeps)

			Expect(ui.PasswordPrompts).To(ContainSubstrings(
				[]string{"Current Password"},
				[]string{"New Password"},
				[]string{"Verify Password"},
			))

			Expect(ui.Outputs()).To(ContainSubstrings(
				[]string{"Changing password..."},
				[]string{"OK"},
				[]string{"Please log in again"},
			))

			Expect(pwDeps.PwdRepo.UpdateNewPassword).To(Equal("new-password"))
			Expect(pwDeps.PwdRepo.UpdateOldPassword).To(Equal("old-password"))

			Expect(pwDeps.Config.AccessToken()).To(Equal(""))
			Expect(pwDeps.Config.OrganizationFields()).To(Equal(models.OrganizationFields{}))
			Expect(pwDeps.Config.SpaceFields()).To(Equal(models.SpaceFields{}))
		})

		It("fails when the password verification does not match", func() {
			ui, _ := callPassword([]string{"old-password", "new-password", "new-password-with-error"}, pwDeps)

			Expect(ui.PasswordPrompts).To(ContainSubstrings(
				[]string{"Current Password"},
				[]string{"New Password"},
				[]string{"Verify Password"},
			))

			Expect(ui.Outputs()).To(ContainSubstrings(
				[]string{"FAILED"},
				[]string{"Password verification does not match"},
			))

			Expect(pwDeps.PwdRepo.UpdateNewPassword).To(Equal(""))
		})

		It("fails when the current password does not match", func() {
			pwDeps.PwdRepo.UpdateUnauthorized = true
			ui, _ := callPassword([]string{"old-password", "new-password", "new-password"}, pwDeps)

			Expect(ui.PasswordPrompts).To(ContainSubstrings(
				[]string{"Current Password"},
				[]string{"New Password"},
				[]string{"Verify Password"},
			))

			Expect(ui.Outputs()).To(ContainSubstrings(
				[]string{"Changing password..."},
				[]string{"FAILED"},
				[]string{"Current password did not match"},
			))

			Expect(pwDeps.PwdRepo.UpdateNewPassword).To(Equal("new-password"))
			Expect(pwDeps.PwdRepo.UpdateOldPassword).To(Equal("old-password"))
		})
	})
})

type passwordDeps struct {
	ReqFactory *requirementsfakes.FakeFactory
	PwdRepo    *apifakes.OldFakePasswordRepo
	Config     coreconfig.Repository
}

func getPasswordDeps() passwordDeps {
	return passwordDeps{
		ReqFactory: new(requirementsfakes.FakeFactory),
		PwdRepo:    &apifakes.OldFakePasswordRepo{UpdateUnauthorized: true},
		Config:     testconfig.NewRepository(),
	}
}
