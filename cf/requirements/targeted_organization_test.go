package requirements_test

import (
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/models"

	testconfig "github.com/liamawhite/cli-with-i18n/util/testhelpers/configuration"

	. "github.com/liamawhite/cli-with-i18n/cf/requirements"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TargetedOrganizationRequirement", func() {
	var (
		config coreconfig.ReadWriter
	)

	BeforeEach(func() {
		config = testconfig.NewRepositoryWithDefaults()
	})

	Context("when the user has an org targeted", func() {
		It("succeeds", func() {
			req := NewTargetedOrgRequirement(config)
			err := req.Execute()
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("when the user does not have an org targeted", func() {
		It("errors", func() {
			config.SetOrganizationFields(models.OrganizationFields{})

			err := NewTargetedOrgRequirement(config).Execute()

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("No org targeted"))
		})
	})
})
