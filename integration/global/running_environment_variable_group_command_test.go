package global

import (
	"fmt"
	"math/rand"

	"github.com/liamawhite/cli-with-i18n/integration/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("running-environment-variable-group command", func() {
	var (
		key1 string
		key2 string
		val1 string
		val2 int
	)

	BeforeEach(func() {
		helpers.LoginCF()

		key1 = helpers.PrefixedRandomName("key1")
		key2 = helpers.PrefixedRandomName("key2")
		val1 = helpers.PrefixedRandomName("val1")
		val2 = rand.Intn(2000)

		json := fmt.Sprintf(`{"%s":"%s", "%s":%d}`, key1, val1, key2, val2)
		session := helpers.CF("set-running-environment-variable-group", json)
		Eventually(session).Should(Exit(0))
	})

	AfterEach(func() {
		session := helpers.CF("set-running-environment-variable-group", "{}")
		Eventually(session).Should(Exit(0))
	})

	It("gets running environment variables", func() {
		session := helpers.CF("running-environment-variable-group")
		Eventually(session).Should(Say("%s\\s+%s", key1, val1))
		Eventually(session).Should(Say("%s\\s+%d", key2, val2))
		Eventually(session).Should(Exit(0))
	})
})
