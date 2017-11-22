package shared_test

import (
	"github.com/liamawhite/cli-with-i18n/command/commandfakes"
	. "github.com/liamawhite/cli-with-i18n/command/v3/shared"
	"github.com/liamawhite/cli-with-i18n/util/ui"

	"github.com/liamawhite/cli-with-i18n/api/uaa"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("New Clients", func() {
	var (
		binaryName    string
		fakeConfig    *commandfakes.FakeConfig
		testUI        *ui.UI
		fakeUAAClient *uaa.Client
	)

	BeforeEach(func() {
		binaryName = "faceman"
		fakeConfig = new(commandfakes.FakeConfig)
		fakeConfig.BinaryNameReturns(binaryName)

		testUI = ui.NewTestUI(NewBuffer(), NewBuffer(), NewBuffer())
	})

	It("returns a networking client", func() {
		client, err := NewNetworkingClient("some-url", fakeConfig, fakeUAAClient, testUI)
		Expect(err).NotTo(HaveOccurred())
		Expect(client).NotTo(BeNil())
	})

	Context("when the network policy api endpoint is not set", func() {
		It("returns an error", func() {
			_, err := NewNetworkingClient("", fakeConfig, fakeUAAClient, testUI)
			Expect(err).To(MatchError("This command requires Network Policy API V1. Your targeted endpoint does not expose it."))
		})
	})
})
