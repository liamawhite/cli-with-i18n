package commands_test

import (
	"errors"

	"github.com/liamawhite/cli-with-i18n/cf/api/authentication/authenticationfakes"
	"github.com/liamawhite/cli-with-i18n/cf/commandregistry"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"
	"github.com/liamawhite/cli-with-i18n/cf/requirements/requirementsfakes"
	"github.com/liamawhite/cli-with-i18n/cf/trace/tracefakes"
	"github.com/liamawhite/cli-with-i18n/plugin/models"
	testcmd "github.com/liamawhite/cli-with-i18n/util/testhelpers/commands"
	testconfig "github.com/liamawhite/cli-with-i18n/util/testhelpers/configuration"
	testterm "github.com/liamawhite/cli-with-i18n/util/testhelpers/terminal"

	"os"

	. "github.com/liamawhite/cli-with-i18n/util/testhelpers/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OauthToken", func() {
	var (
		ui                  *testterm.FakeUI
		authRepo            *authenticationfakes.FakeRepository
		requirementsFactory *requirementsfakes.FakeFactory
		configRepo          coreconfig.Repository
		deps                commandregistry.Dependency
	)

	updateCommandDependency := func(pluginCall bool) {
		deps.UI = ui
		deps.RepoLocator = deps.RepoLocator.SetAuthenticationRepository(authRepo)
		deps.Config = configRepo
		commandregistry.Commands.SetCommand(commandregistry.Commands.FindCommand("oauth-token").SetDependency(deps, pluginCall))
	}

	BeforeEach(func() {
		ui = &testterm.FakeUI{}
		fakeLogger := new(tracefakes.FakePrinter)
		authRepo = new(authenticationfakes.FakeRepository)
		configRepo = testconfig.NewRepositoryWithDefaults()
		requirementsFactory = new(requirementsfakes.FakeFactory)
		deps = commandregistry.NewDependency(os.Stdout, fakeLogger, "")
	})

	runCommand := func() bool {
		return testcmd.RunCLICommand("oauth-token", []string{}, requirementsFactory, updateCommandDependency, false, ui)
	}

	Describe("requirements", func() {
		It("fails when the user is not logged in", func() {
			requirementsFactory.NewLoginRequirementReturns(requirements.Failing{Message: "not logged in"})
			Expect(runCommand()).ToNot(HavePassedRequirements())
		})
	})

	Describe("when logged in", func() {
		BeforeEach(func() {
			requirementsFactory.NewLoginRequirementReturns(requirements.Passing{})
		})

		It("fails if oauth refresh fails", func() {
			authRepo.RefreshAuthTokenReturns("", errors.New("Could not refresh"))
			runCommand()

			Expect(ui.Outputs()).To(ContainSubstrings(
				[]string{"FAILED"},
				[]string{"Could not refresh"},
			))
		})

		It("returns to the user the oauth token after a refresh", func() {
			authRepo.RefreshAuthTokenReturns("1234567890", nil)
			runCommand()

			Expect(ui.Outputs()).To(ContainSubstrings(
				[]string{"1234567890"},
			))
		})

		Context("when invoked by a plugin", func() {
			var (
				pluginModel plugin_models.GetOauthToken_Model
			)

			BeforeEach(func() {
				pluginModel = plugin_models.GetOauthToken_Model{}
				deps.PluginModels.OauthToken = &pluginModel
			})

			It("populates the plugin model upon execution", func() {
				authRepo.RefreshAuthTokenReturns("911999111", nil)
				testcmd.RunCLICommand("oauth-token", []string{}, requirementsFactory, updateCommandDependency, true, ui)
				Expect(pluginModel.Token).To(Equal("911999111"))
			})
		})
	})
})
