package service_test

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

var _ = Describe("unbind-service command", func() {
	var (
		app                 models.Application
		ui                  *testterm.FakeUI
		config              coreconfig.Repository
		serviceInstance     models.ServiceInstance
		requirementsFactory *requirementsfakes.FakeFactory
		serviceBindingRepo  *apifakes.FakeServiceBindingRepository
		deps                commandregistry.Dependency
	)

	updateCommandDependency := func(pluginCall bool) {
		deps.UI = ui
		deps.RepoLocator = deps.RepoLocator.SetServiceBindingRepository(serviceBindingRepo)
		deps.Config = config
		commandregistry.Commands.SetCommand(commandregistry.Commands.FindCommand("unbind-service").SetDependency(deps, pluginCall))
	}

	BeforeEach(func() {
		ui = new(testterm.FakeUI)
		serviceBindingRepo = new(apifakes.FakeServiceBindingRepository)

		app = models.Application{
			ApplicationFields: models.ApplicationFields{
				Name: "my-app",
				GUID: "my-app-guid",
			},
		}

		serviceInstance = models.ServiceInstance{
			ServiceInstanceFields: models.ServiceInstanceFields{
				Name: "my-service",
				GUID: "my-service-guid",
			},
		}

		config = testconfig.NewRepositoryWithDefaults()
		requirementsFactory = new(requirementsfakes.FakeFactory)
		applicationReq := new(requirementsfakes.FakeApplicationRequirement)
		applicationReq.GetApplicationReturns(app)
		requirementsFactory.NewApplicationRequirementReturns(applicationReq)
		serviceInstanceReq := new(requirementsfakes.FakeServiceInstanceRequirement)
		serviceInstanceReq.GetServiceInstanceReturns(serviceInstance)
		requirementsFactory.NewServiceInstanceRequirementReturns(serviceInstanceReq)
	})

	callUnbindService := func(args []string) bool {
		return testcmd.RunCLICommand("unbind-service", args, requirementsFactory, updateCommandDependency, false, ui)
	}

	Context("when not logged in", func() {
		It("fails requirements when not logged in", func() {
			requirementsFactory.NewLoginRequirementReturns(requirements.Failing{Message: "not logged in"})
			Expect(testcmd.RunCLICommand("unbind-service", []string{"my-service", "my-app"}, requirementsFactory, updateCommandDependency, false, ui)).To(BeFalse())
		})
	})

	Context("when logged in", func() {
		BeforeEach(func() {
			requirementsFactory.NewLoginRequirementReturns(requirements.Passing{})
		})

		Context("when the service instance exists", func() {
			It("unbinds a service from an app", func() {
				callUnbindService([]string{"my-app", "my-service"})

				Expect(ui.Outputs()).To(ContainSubstrings(
					[]string{"Unbinding app", "my-service", "my-app", "my-org", "my-space", "my-user"},
					[]string{"OK"},
				))

				Expect(serviceBindingRepo.DeleteCallCount()).To(Equal(1))
				serviceInstance, applicationGUID := serviceBindingRepo.DeleteArgsForCall(0)
				Expect(serviceInstance).To(Equal(serviceInstance))
				Expect(applicationGUID).To(Equal("my-app-guid"))
			})
		})

		Context("when the service instance does not exist", func() {
			BeforeEach(func() {
				serviceBindingRepo.DeleteReturns(false, nil)
			})

			It("warns the user the the service instance does not exist", func() {
				callUnbindService([]string{"my-app", "my-service"})

				Expect(ui.Outputs()).To(ContainSubstrings(
					[]string{"Unbinding app", "my-service", "my-app"},
					[]string{"OK"},
					[]string{"my-service", "my-app", "did not exist"},
				))

				Expect(serviceBindingRepo.DeleteCallCount()).To(Equal(1))
				serviceInstance, applicationGUID := serviceBindingRepo.DeleteArgsForCall(0)
				Expect(serviceInstance).To(Equal(serviceInstance))
				Expect(applicationGUID).To(Equal("my-app-guid"))
			})
		})

		It("when no parameters are given the command fails with usage", func() {
			callUnbindService([]string{"my-service"})
			Expect(ui.Outputs()).To(ContainSubstrings(
				[]string{"Incorrect Usage", "Requires", "argument"},
			))

			ui = &testterm.FakeUI{}
			callUnbindService([]string{"my-app"})
			Expect(ui.Outputs()).To(ContainSubstrings(
				[]string{"Incorrect Usage", "Requires", "argument"},
			))

			ui = &testterm.FakeUI{}
			callUnbindService([]string{"my-app", "my-service"})
			Expect(ui.Outputs()).ToNot(ContainSubstrings(
				[]string{"Incorrect Usage", "Requires", "argument"},
			))
		})
	})
})
