package isolated

import (
	"github.com/liamawhite/cli-with-i18n/integration/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
	. "github.com/onsi/gomega/ghttp"
)

var _ = Describe("enable-org-isolation command", func() {
	var organizationName string
	var isolationSegmentName string

	BeforeEach(func() {
		organizationName = helpers.NewOrgName()
		isolationSegmentName = helpers.NewIsolationSegmentName()
	})

	Describe("help", func() {
		Context("when --help flag is set", func() {
			It("Displays command usage to output", func() {
				session := helpers.CF("enable-org-isolation", "--help")
				Eventually(session).Should(Say("NAME:"))
				Eventually(session).Should(Say("enable-org-isolation - Entitle an organization to an isolation segment"))
				Eventually(session).Should(Say("USAGE:"))
				Eventually(session).Should(Say("cf enable-org-isolation ORG_NAME SEGMENT_NAME"))
				Eventually(session).Should(Say("SEE ALSO:"))
				Eventually(session).Should(Say("create-isolation-segment, isolation-segments, set-org-default-isolation-segment, set-space-isolation-segment"))
				Eventually(session).Should(Exit(0))
			})
		})
	})

	Context("when the environment is not setup correctly", func() {
		Context("when no API endpoint is set", func() {
			BeforeEach(func() {
				helpers.UnsetAPI()
			})

			It("fails with no API endpoint set message", func() {
				session := helpers.CF("enable-org-isolation", organizationName, isolationSegmentName)
				Eventually(session).Should(Say("FAILED"))
				Eventually(session.Err).Should(Say("No API endpoint set. Use 'cf login' or 'cf api' to target an endpoint."))
				Eventually(session).Should(Exit(1))
			})
		})

		Context("when the v3 api does not exist", func() {
			var server *Server

			BeforeEach(func() {
				server = helpers.StartAndTargetServerWithoutV3API()
			})

			AfterEach(func() {
				server.Close()
			})

			It("fails with error message that the minimum version is not met", func() {
				session := helpers.CF("enable-org-isolation", organizationName, isolationSegmentName)
				Eventually(session).Should(Say("FAILED"))
				Eventually(session.Err).Should(Say("This command requires CF API version 3\\.11\\.0 or higher\\."))
				Eventually(session).Should(Exit(1))
			})
		})

		Context("when the v3 api version is lower than the minimum version", func() {
			var server *Server

			BeforeEach(func() {
				server = helpers.StartAndTargetServerWithV3Version("3.0.0")
			})

			AfterEach(func() {
				server.Close()
			})

			It("fails with error message that the minimum version is not met", func() {
				session := helpers.CF("enable-org-isolation", organizationName, isolationSegmentName)
				Eventually(session).Should(Say("FAILED"))
				Eventually(session.Err).Should(Say("This command requires CF API version 3\\.11\\.0 or higher\\."))
				Eventually(session).Should(Exit(1))
			})
		})

		Context("when not logged in", func() {
			BeforeEach(func() {
				helpers.LogoutCF()
			})

			It("fails with not logged in message", func() {
				session := helpers.CF("enable-org-isolation", organizationName, isolationSegmentName)
				Eventually(session).Should(Say("FAILED"))
				Eventually(session.Err).Should(Say("Not logged in. Use 'cf login' to log in."))
				Eventually(session).Should(Exit(1))
			})
		})
	})

	Context("when the environment is set up correctly", func() {
		var userName string

		BeforeEach(func() {
			helpers.LoginCF()
			userName, _ = helpers.GetCredentials()
		})

		Context("when the isolation segment does not exist", func() {
			It("fails with isolation segment not found message", func() {
				session := helpers.CF("enable-org-isolation", organizationName, isolationSegmentName)
				Eventually(session).Should(Say("Enabling isolation segment %s for org %s as %s...", isolationSegmentName, organizationName, userName))
				Eventually(session).Should(Say("FAILED"))
				Eventually(session.Err).Should(Say("Isolation segment '%s' not found.", isolationSegmentName))
				Eventually(session).Should(Exit(1))
			})
		})

		Context("when the isolation segment exists", func() {
			BeforeEach(func() {
				Eventually(helpers.CF("create-isolation-segment", isolationSegmentName)).Should(Exit(0))
			})

			Context("when the organization does not exist", func() {
				It("fails with organization not found message", func() {
					session := helpers.CF("enable-org-isolation", organizationName, isolationSegmentName)
					Eventually(session).Should(Say("Enabling isolation segment %s for org %s as %s...", isolationSegmentName, organizationName, userName))
					Eventually(session).Should(Say("FAILED"))
					Eventually(session.Err).Should(Say("Organization '%s' not found.", organizationName))
					Eventually(session).Should(Exit(1))
				})
			})

			Context("when the organization exists", func() {
				BeforeEach(func() {
					helpers.CreateOrg(organizationName)
					helpers.TargetOrg(organizationName)
				})

				AfterEach(func() {
					helpers.QuickDeleteOrg(organizationName)
				})

				It("displays OK", func() {
					session := helpers.CF("enable-org-isolation", organizationName, isolationSegmentName)
					Eventually(session).Should(Say("Enabling isolation segment %s for org %s as %s...", isolationSegmentName, organizationName, userName))
					Eventually(session).Should(Say("OK"))
					Eventually(session).Should(Exit(0))
				})

				Context("when the isolation is already enabled", func() {
					BeforeEach(func() {
						Eventually(helpers.CF("enable-org-isolation", organizationName, isolationSegmentName)).Should(Exit(0))
					})

					It("displays OK", func() {
						session := helpers.CF("enable-org-isolation", organizationName, isolationSegmentName)
						Eventually(session).Should(Say("Enabling isolation segment %s for org %s as %s...", isolationSegmentName, organizationName, userName))
						Eventually(session).Should(Say("OK"))
						Eventually(session).Should(Exit(0))
					})
				})
			})
		})
	})
})
