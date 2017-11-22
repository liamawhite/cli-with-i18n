package requirements_test

import (
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/requirements"

	testconfig "github.com/liamawhite/cli-with-i18n/util/testhelpers/configuration"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RoutingApi", func() {
	var (
		config      coreconfig.Repository
		requirement requirements.RoutingAPIRequirement
	)

	BeforeEach(func() {
		config = testconfig.NewRepositoryWithAccessToken(coreconfig.TokenInfo{Username: "my-user"})
		requirement = requirements.NewRoutingAPIRequirement(config)
	})

	Context("when the config has a zero-length RoutingAPIEndpoint", func() {
		BeforeEach(func() {
			config.SetRoutingAPIEndpoint("")
		})

		It("errors", func() {
			err := requirement.Execute()
			Expect(err.Error()).To(ContainSubstring("This command requires the Routing API. Your targeted endpoint reports it is not enabled."))
		})
	})

	Context("when the config has a RoutingAPIEndpoint", func() {
		BeforeEach(func() {
			config.SetRoutingAPIEndpoint("api.example.com")
		})

		It("does not error", func() {
			err := requirement.Execute()
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
