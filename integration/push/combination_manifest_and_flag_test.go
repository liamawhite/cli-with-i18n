package push

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/liamawhite/cli-with-i18n/integration/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("push with a simple manifest and flags", func() {
	var (
		appName        string
		pathToManifest string
	)

	BeforeEach(func() {
		appName = helpers.NewAppName()

		tmpFile, err := ioutil.TempFile("", "combination-manifest")
		Expect(err).ToNot(HaveOccurred())
		pathToManifest = tmpFile.Name()
		Expect(tmpFile.Close()).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		Expect(os.RemoveAll(pathToManifest)).ToNot(HaveOccurred())
	})

	Context("when the app is new", func() {
		Context("when pushing a single app from the manifest", func() {
			Context("when the manifest is passed via '-f'", func() {
				Context("when pushing the app from the current directory", func() {
					BeforeEach(func() {
						helpers.WriteManifest(pathToManifest, map[string]interface{}{
							"applications": []map[string]string{
								{
									"name": appName,
								},
							},
						})
					})

					It("pushes the app from the current directory and the manifest for app settings", func() {
						helpers.WithHelloWorldApp(func(dir string) {
							session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, "-f", pathToManifest)
							Eventually(session).Should(Say("Getting app info\\.\\.\\."))
							Eventually(session).Should(Say("Creating app with these attributes\\.\\.\\."))
							Eventually(session).Should(Say("\\+\\s+name:\\s+%s", appName))
							Eventually(session).Should(Say("\\s+routes:"))
							Eventually(session).Should(Say("(?i)\\+\\s+%s.%s", appName, defaultSharedDomain()))
							Eventually(session).Should(Say("Mapping routes\\.\\.\\."))
							Eventually(session).Should(Say("Uploading files\\.\\.\\."))
							Eventually(session).Should(Say("100.00%"))
							Eventually(session).Should(Say("Waiting for API to complete processing files\\.\\.\\."))
							helpers.ConfirmStagingLogs(session)
							Eventually(session).Should(Say("Waiting for app to start\\.\\.\\."))
							Eventually(session).Should(Say("requested state:\\s+started"))
							Eventually(session).Should(Exit(0))
						})

						session := helpers.CF("app", appName)
						Eventually(session).Should(Say("name:\\s+%s", appName))
						Eventually(session).Should(Exit(0))
					})
				})

				Context("when the path to the application is provided in the manifest", func() {
					It("pushes the app from the path specified in the manifest and uses the manifest for app settings", func() {
						helpers.WithHelloWorldApp(func(dir string) {
							helpers.WriteManifest(pathToManifest, map[string]interface{}{
								"applications": []map[string]string{
									{
										"name": appName,
										"path": filepath.Base(dir),
									},
								},
							})

							session := helpers.CF(PushCommandName, "-f", pathToManifest)
							Eventually(session).Should(Say("Getting app info\\.\\.\\."))
							Eventually(session).Should(Say("Creating app with these attributes\\.\\.\\."))
							Eventually(session).Should(Say("\\+\\s+name:\\s+%s", appName))
							Eventually(session).Should(Say("\\s+path:\\s+%s", regexp.QuoteMeta(dir)))
							Eventually(session).Should(Say("requested state:\\s+started"))
							Eventually(session).Should(Exit(0))
						})

						session := helpers.CF("app", appName)
						Eventually(session).Should(Say("name:\\s+%s", appName))
						Eventually(session).Should(Exit(0))
					})
				})
			})

			Context("manifest contains a path and a '-p' is provided", func() {
				var tempDir string

				BeforeEach(func() {
					var err error
					tempDir, err = ioutil.TempDir("", "combination-manifest-with-p")
					Expect(err).ToNot(HaveOccurred())

					helpers.WriteManifest(filepath.Join(tempDir, "manifest.yml"), map[string]interface{}{
						"applications": []map[string]string{
							{
								"name": appName,
								"path": "does-not-exist",
							},
						},
					})
				})

				AfterEach(func() {
					Expect(os.RemoveAll(tempDir)).ToNot(HaveOccurred())
				})

				It("overrides the manifest path with the '-p' path", func() {
					helpers.WithHelloWorldApp(func(dir string) {
						session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: tempDir}, PushCommandName, "-p", dir)
						Eventually(session).Should(Say("\\+\\s+name:\\s+%s", appName))
						Eventually(session).Should(Say("\\s+path:\\s+%s", regexp.QuoteMeta(dir)))
						Eventually(session).Should(Say("requested state:\\s+started"))
						Eventually(session).Should(Exit(0))
					})

					session := helpers.CF("app", appName)
					Eventually(session).Should(Say("name:\\s+%s", appName))
					Eventually(session).Should(Exit(0))
				})
			})

			Context("manifest contains a name and a name is provided", func() {
				It("overrides the manifest name", func() {
					helpers.WithHelloWorldApp(func(dir string) {
						helpers.WriteManifest(filepath.Join(dir, "manifest.yml"), map[string]interface{}{
							"applications": []map[string]string{
								{
									"name": "earle",
								},
							},
						})

						session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, appName)
						Eventually(session).Should(Say("\\+\\s+name:\\s+%s", appName))
						Eventually(session).Should(Say("requested state:\\s+started"))
						Eventually(session).Should(Exit(0))
					})

					session := helpers.CF("app", appName)
					Eventually(session).Should(Say("name:\\s+%s", appName))
					Eventually(session).Should(Exit(0))
				})
			})

			Context("when the --no-manifest flag is passed", func() {
				It("does not use the provided manifest", func() {
					helpers.WithHelloWorldApp(func(dir string) {
						helpers.WriteManifest(filepath.Join(dir, "manifest.yml"), map[string]interface{}{
							"applications": []map[string]string{
								{
									"name": "crazy-jerry",
								},
							},
						})

						session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, "--no-manifest", appName)
						Eventually(session).Should(Say("Getting app info\\.\\.\\."))
						Eventually(session).Should(Say("Creating app with these attributes\\.\\.\\."))
						Eventually(session).Should(Say("\\+\\s+name:\\s+%s", appName))
						Eventually(session).Should(Say("\\s+routes:"))
						Eventually(session).Should(Say("(?i)\\+\\s+%s.%s", appName, defaultSharedDomain()))
						Eventually(session).Should(Say("Mapping routes\\.\\.\\."))
						Eventually(session).Should(Say("Uploading files\\.\\.\\."))
						Eventually(session).Should(Say("100.00%"))
						Eventually(session).Should(Say("Waiting for API to complete processing files\\.\\.\\."))
						helpers.ConfirmStagingLogs(session)
						Eventually(session).Should(Say("Waiting for app to start\\.\\.\\."))
						Eventually(session).Should(Say("requested state:\\s+started"))
						Eventually(session).Should(Exit(0))
					})

					session := helpers.CF("app", appName)
					Eventually(session.Out).Should(Say("name:\\s+%s", appName))
					Eventually(session).Should(Exit(0))
				})
			})
		})

		Context("when pushing multiple apps from the manifest", func() {
			Context("manifest contains multiple apps and a '-p' is provided", func() {
				var tempDir string

				BeforeEach(func() {
					var err error
					tempDir, err = ioutil.TempDir("", "combination-manifest-with-p")
					Expect(err).ToNot(HaveOccurred())

					helpers.WriteManifest(filepath.Join(tempDir, "manifest.yml"), map[string]interface{}{
						"applications": []map[string]string{
							{
								"name": "name-1",
							},
							{
								"name": "name-2",
							},
						},
					})
				})

				AfterEach(func() {
					Expect(os.RemoveAll(tempDir)).ToNot(HaveOccurred())
				})

				It("returns an error", func() {
					helpers.WithHelloWorldApp(func(dir string) {
						session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: tempDir}, PushCommandName, "-p", dir)
						Eventually(session.Err).Should(Say("Incorrect Usage: Command line flags \\(except -f\\) cannot be applied when pushing multiple apps from a manifest file\\."))
						Eventually(session).Should(Exit(1))
					})
				})
			})

			Context("manifest contains multiple apps and '--no-start' is provided", func() {
				var appName1, appName2 string

				BeforeEach(func() {
					appName1 = helpers.NewAppName()
					appName2 = helpers.NewAppName()
				})

				It("does not start the apps", func() {
					helpers.WithHelloWorldApp(func(dir string) {
						helpers.WriteManifest(filepath.Join(dir, "manifest.yml"), map[string]interface{}{
							"applications": []map[string]string{
								{"name": appName1},
								{"name": appName2},
							},
						})

						session := helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, PushCommandName, "--no-start")
						Eventually(session).Should(Say("Getting app info\\.\\.\\."))
						Eventually(session).Should(Say("Creating app with these attributes\\.\\.\\."))
						Eventually(session).Should(Say("\\s+name:\\s+%s", appName1))
						Eventually(session).Should(Say("requested state:\\s+stopped"))
						Eventually(session).Should(Say("\\s+name:\\s+%s", appName2))
						Eventually(session).Should(Say("requested state:\\s+stopped"))
						Eventually(session).Should(Exit(0))
					})
				})
			})
		})
	})
})
