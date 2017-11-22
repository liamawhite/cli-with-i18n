package push

import (
	"path/filepath"

	"github.com/liamawhite/cli-with-i18n/integration/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("push with different start command values", func() {
	var (
		appName string
	)

	BeforeEach(func() {
		appName = helpers.NewAppName()
	})

	Context("when the start command flag is provided", func() {
		It("sets the start command correctly for the pushed app", func() {
			helpers.WithHelloWorldApp(func(dir string) {
				By("pushing the app with no provided start command uses detected command")
				session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
				Eventually(session).Should(Say("start command:\\s+\\$HOME/boot.sh"))
				Eventually(session).Should(Exit(0))

				By("pushing the app with a start command uses provided start command")
				session = helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir},
					PushCommandName, appName,
					"-c", "$HOME/boot.sh && echo hello")
				Eventually(session).Should(Say("start command:\\s+\\$HOME/boot.sh && echo hello"))
				Eventually(session).Should(Exit(0))

				By("pushing the app with no provided start command again uses previously set command")
				session = helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
				Eventually(session).Should(Say("start command:\\s+\\$HOME/boot.sh && echo hello"))
				Eventually(session).Should(Exit(0))

				By("pushing the app with default uses detected command")
				session = helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir},
					PushCommandName, appName,
					"-c", "default")
				Eventually(session).Should(Say("(?m)start command:\\s+\\$HOME/boot.sh$"))
				Eventually(session).Should(Exit(0))

				By("pushing the app with null uses detected command")
				session = helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir},
					PushCommandName, appName,
					"-c", "null")
				Eventually(session).Should(Say("(?m)start command:\\s+\\$HOME/boot.sh$"))
				Eventually(session).Should(Exit(0))
			})
		})
	})

	Context("when the start command is provided in the manifest", func() {
		It("sets the start command correctly for the pushed app", func() {
			helpers.WithHelloWorldApp(func(dir string) {
				By("pushing the app with no provided start command uses detected command")
				helpers.WriteManifest(filepath.Join(dir, "manifest.yml"), map[string]interface{}{
					"applications": []map[string]interface{}{
						{
							"name": appName,
							"path": dir,
						},
					},
				})
				session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
				Eventually(session).Should(Say("start command:\\s+\\$HOME/boot.sh"))
				Eventually(session).Should(Exit(0))

				By("pushing the app with a start command uses provided start command")
				helpers.WriteManifest(filepath.Join(dir, "manifest.yml"), map[string]interface{}{
					"applications": []map[string]interface{}{
						{
							"name":    appName,
							"path":    dir,
							"command": "$HOME/boot.sh && echo hello",
						},
					},
				})
				session = helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
				Eventually(session).Should(Say("start command:\\s+\\$HOME/boot.sh && echo hello"))
				Eventually(session).Should(Exit(0))

				By("pushing the app with no provided start command again uses previously set command")
				helpers.WriteManifest(filepath.Join(dir, "manifest.yml"), map[string]interface{}{
					"applications": []map[string]interface{}{
						{
							"name": appName,
							"path": dir,
						},
					},
				})
				session = helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
				Eventually(session).Should(Say("start command:\\s+\\$HOME/boot.sh && echo hello"))
				Eventually(session).Should(Exit(0))

				By("pushing the app with default uses detected command")
				helpers.WriteManifest(filepath.Join(dir, "manifest.yml"), map[string]interface{}{
					"applications": []map[string]interface{}{
						{
							"name":    appName,
							"path":    dir,
							"command": "default",
						},
					},
				})
				session = helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
				Eventually(session).Should(Say("(?m)start command:\\s+\\$HOME/boot.sh$"))
				Eventually(session).Should(Exit(0))

				By("pushing the app with null uses detected command")
				helpers.WriteManifest(filepath.Join(dir, "manifest.yml"), map[string]interface{}{
					"applications": []map[string]interface{}{
						{
							"name":    appName,
							"path":    dir,
							"command": nil,
						},
					},
				})
				session = helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
				Eventually(session).Should(Say("(?m)start command:\\s+\\$HOME/boot.sh$"))
				Eventually(session).Should(Exit(0))
			})
		})
	})
})
