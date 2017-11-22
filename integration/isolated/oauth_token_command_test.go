package isolated

import (
	"github.com/liamawhite/cli-with-i18n/integration/helpers"
	"github.com/liamawhite/cli-with-i18n/util/configv3"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("oauth-token command", func() {
	Context("help", func() {
		It("displays the help information", func() {
			session := helpers.CF("oauth-token", "--help")

			Eventually(session.Out).Should(Say("NAME:"))
			Eventually(session.Out).Should(Say("oauth-token - Retrieve and display the OAuth token for the current session"))
			Eventually(session.Out).Should(Say("USAGE:"))
			Eventually(session.Out).Should(Say("cf oauth-token"))
			Eventually(session.Out).Should(Say("SEE ALSO:"))
			Eventually(session.Out).Should(Say("curl"))
			Eventually(session).Should(Exit(0))
		})
	})

	Context("when the environment is not setup correctly", func() {
		Context("when no API endpoint is set", func() {
			BeforeEach(func() {
				helpers.UnsetAPI()
			})

			It("fails with no API endpoint set message", func() {
				session := helpers.CF("oauth-token")

				Eventually(session.Out).Should(Say("FAILED"))
				Eventually(session.Err).Should(Say("No API endpoint set\\. Use 'cf login' or 'cf api' to target an endpoint\\."))
				Eventually(session).Should(Exit(1))
			})
		})

		Context("when not logged in", func() {
			BeforeEach(func() {
				helpers.LogoutCF()
			})

			It("fails with not logged in message", func() {
				session := helpers.CF("oauth-token")

				Eventually(session.Out).Should(Say("FAILED"))
				Eventually(session.Err).Should(Say("Not logged in\\. Use 'cf login' to log in\\."))
				Eventually(session).Should(Exit(1))
			})
		})
	})

	Context("when the environment is setup correctly", func() {
		BeforeEach(func() {
			helpers.LoginCF()
		})

		Context("when the refresh token is invalid", func() {
			BeforeEach(func() {
				helpers.SetConfig(func(conf *configv3.Config) {
					conf.ConfigFile.RefreshToken = "invalid-refresh-token"
				})
			})

			It("displays an error and exits 1", func() {
				session := helpers.CF("oauth-token")

				Eventually(session.Out).Should(Say("FAILED"))
				Eventually(session.Err).Should(Say("The token expired, was revoked, or the token ID is incorrect\\. Please log back in to re-authenticate\\."))
				Eventually(session).Should(Exit(1))
			})
		})

		Context("when the oauth client ID and secret combination is invalid", func() {
			BeforeEach(func() {
				helpers.SetConfig(func(conf *configv3.Config) {
					conf.ConfigFile.UAAOAuthClient = "foo"
					conf.ConfigFile.UAAOAuthClientSecret = "bar"
				})
			})

			It("displays an error and exits 1", func() {
				session := helpers.CF("oauth-token")

				Eventually(session.Out).Should(Say("FAILED"))
				Eventually(session.Err).Should(Say("Credentials were rejected, please try again\\."))
				Eventually(session).Should(Exit(1))
			})
		})

		Context("when the refresh token and oauth creds are valid", func() {
			It("refreshes the access token and displays it", func() {
				session := helpers.CF("oauth-token")

				Eventually(session.Out).Should(Say("bearer .+"))
				Eventually(session).Should(Exit(0))
			})
		})
	})
})
