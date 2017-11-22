package isolated

import (
	"github.com/liamawhite/cli-with-i18n/integration/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
	. "github.com/onsi/gomega/ghttp"
)

var _ = Describe("reset-org-default-isolation-segment command", func() {
	var orgName string

	BeforeEach(func() {
		orgName = helpers.NewOrgName()
	})

	Describe("help", func() {
		Context("when --help flag is set", func() {
			It("displays command usage to output", func() {
				session := helpers.CF("reset-org-default-isolation-segment", "--help")
				Eventually(session).Should(Say("NAME:"))
				Eventually(session).Should(Say("reset-org-default-isolation-segment - Reset the default isolation segment used for apps in spaces of an org"))
				Eventually(session).Should(Say("USAGE:"))
				Eventually(session).Should(Say("cf reset-org-default-isolation-segment ORG_NAME"))
				Eventually(session).Should(Say("SEE ALSO:"))
				Eventually(session).Should(Say("org, restart"))
				Eventually(session).Should(Exit(0))
			})
		})
	})

	Context("when the environment is not set-up correctly", func() {
		Context("when no API endpoint is set", func() {
			BeforeEach(func() {
				helpers.UnsetAPI()
			})

			It("fails with no API endpoint set message", func() {
				session := helpers.CF("reset-org-default-isolation-segment", orgName)
				Eventually(session).Should(Say("FAILED"))
				Eventually(session.Err).Should(Say("No API endpoint set\\. Use 'cf login' or 'cf api' to target an endpoint\\."))
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
				session := helpers.CF("reset-org-default-isolation-segment", orgName)
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
				session := helpers.CF("reset-org-default-isolation-segment", orgName)
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
				session := helpers.CF("reset-org-default-isolation-segment", orgName)
				Eventually(session).Should(Say("FAILED"))
				Eventually(session.Err).Should(Say("Not logged in\\. Use 'cf login' to log in\\."))
				Eventually(session).Should(Exit(1))
			})
		})
	})

	Context("when the environment is set-up correctly", func() {
		var userName string
		var userOrgName string

		BeforeEach(func() {
			helpers.LoginCF()
			userName, _ = helpers.GetCredentials()
			userOrgName = helpers.NewOrgName()
			helpers.CreateOrg(userOrgName)
			helpers.TargetOrg(userOrgName)
		})

		AfterEach(func() {
			helpers.QuickDeleteOrg(userOrgName)
		})

		Context("when the org does not exist", func() {
			It("fails with org not found message", func() {
				session := helpers.CF("reset-org-default-isolation-segment", orgName)
				Eventually(session).Should(Say("Resetting default isolation segment of org %s as %s\\.\\.\\.", orgName, userName))
				Eventually(session).Should(Say("FAILED"))
				Eventually(session.Err).Should(Say("Organization '%s' not found\\.", orgName))
				Eventually(session).Should(Exit(1))
			})
		})

		Context("when the org exists", func() {
			BeforeEach(func() {
				helpers.CreateOrg(orgName)
			})

			AfterEach(func() {
				helpers.QuickDeleteOrg(orgName)
			})

			Context("when the isolation segment is set as the org's default", func() {
				BeforeEach(func() {
					isolationSegmentName := helpers.NewIsolationSegmentName()
					Eventually(helpers.CF("create-isolation-segment", isolationSegmentName)).Should(Exit(0))
					Eventually(helpers.CF("enable-org-isolation", orgName, isolationSegmentName)).Should(Exit(0))
					Eventually(helpers.CF("set-org-default-isolation-segment", orgName, isolationSegmentName)).Should(Exit(0))
				})

				It("displays OK", func() {
					session := helpers.CF("reset-org-default-isolation-segment", orgName)
					Eventually(session).Should(Say("Resetting default isolation segment of org %s as %s\\.\\.\\.", orgName, userName))
					Eventually(session).Should(Say("OK"))
					Eventually(session).Should(Say("Applications in spaces of this org that have no isolation segment assigned will be placed in the platform default isolation segment\\."))
					Eventually(session).Should(Say("Running applications need a restart to be moved there\\."))
					Eventually(session).Should(Exit(0))
				})
			})

			Context("when the org has no default isolation segment", func() {
				It("displays OK", func() {
					session := helpers.CF("reset-org-default-isolation-segment", orgName)
					Eventually(session).Should(Say("Resetting default isolation segment of org %s as %s\\.\\.\\.", orgName, userName))
					Eventually(session).Should(Say("OK"))
					Eventually(session).Should(Say("Applications in spaces of this org that have no isolation segment assigned will be placed in the platform default isolation segment\\."))
					Eventually(session).Should(Say("Running applications need a restart to be moved there\\."))
					Eventually(session).Should(Exit(0))
				})
			})
		})
	})
})
