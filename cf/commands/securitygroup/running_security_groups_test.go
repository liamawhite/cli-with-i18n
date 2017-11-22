package securitygroup_test

import (
	"github.com/liamawhite/cli-with-i18n/cf/commandregistry"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/errors"
	"github.com/liamawhite/cli-with-i18n/cf/models"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"
	"github.com/liamawhite/cli-with-i18n/cf/requirements/requirementsfakes"

	"github.com/liamawhite/cli-with-i18n/cf/api/securitygroups/defaults/running/runningfakes"
	testcmd "github.com/liamawhite/cli-with-i18n/util/testhelpers/commands"
	testconfig "github.com/liamawhite/cli-with-i18n/util/testhelpers/configuration"
	testterm "github.com/liamawhite/cli-with-i18n/util/testhelpers/terminal"

	. "github.com/liamawhite/cli-with-i18n/util/testhelpers/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Running-security-groups command", func() {
	var (
		ui                           *testterm.FakeUI
		configRepo                   coreconfig.Repository
		fakeRunningSecurityGroupRepo *runningfakes.FakeSecurityGroupsRepo
		requirementsFactory          *requirementsfakes.FakeFactory
		deps                         commandregistry.Dependency
	)

	updateCommandDependency := func(pluginCall bool) {
		deps.UI = ui
		deps.RepoLocator = deps.RepoLocator.SetRunningSecurityGroupRepository(fakeRunningSecurityGroupRepo)
		deps.Config = configRepo
		commandregistry.Commands.SetCommand(commandregistry.Commands.FindCommand("running-security-groups").SetDependency(deps, pluginCall))
	}

	BeforeEach(func() {
		ui = &testterm.FakeUI{}
		configRepo = testconfig.NewRepositoryWithDefaults()
		fakeRunningSecurityGroupRepo = new(runningfakes.FakeSecurityGroupsRepo)
		requirementsFactory = new(requirementsfakes.FakeFactory)
	})

	runCommand := func(args ...string) bool {
		return testcmd.RunCLICommand("running-security-groups", args, requirementsFactory, updateCommandDependency, false, ui)
	}

	Describe("requirements", func() {
		It("should fail when not logged in", func() {
			requirementsFactory.NewLoginRequirementReturns(requirements.Failing{Message: "not logged in"})
			Expect(runCommand()).ToNot(HavePassedRequirements())
		})
	})

	Context("when the user is logged in", func() {
		BeforeEach(func() {
			requirementsFactory.NewLoginRequirementReturns(requirements.Passing{})
		})

		Context("when there are some security groups set in the Running group", func() {
			BeforeEach(func() {
				fakeRunningSecurityGroupRepo.ListReturns([]models.SecurityGroupFields{
					{Name: "hiphopopotamus"},
					{Name: "my lyrics are bottomless"},
					{Name: "steve"},
				}, nil)
			})

			It("shows the user the name of the security groups of the Running set", func() {
				Expect(runCommand()).To(BeTrue())
				Expect(ui.Outputs()).To(ContainSubstrings(
					[]string{"Acquiring", "security groups", "my-user"},
					[]string{"hiphopopotamus"},
					[]string{"my lyrics are bottomless"},
					[]string{"steve"},
				))
			})
		})

		Context("when the API returns an error", func() {
			BeforeEach(func() {
				fakeRunningSecurityGroupRepo.ListReturns(nil, errors.New("uh oh"))
			})

			It("fails loudly", func() {
				runCommand()
				Expect(ui.Outputs()).To(ContainSubstrings([]string{"FAILED"}))
			})
		})

		Context("when there are no security groups set in the Running group", func() {
			It("tells the user that there are none", func() {
				runCommand()
				Expect(ui.Outputs()).To(ContainSubstrings(
					[]string{"No", "security groups", "set"},
				))
			})
		})
	})
})
