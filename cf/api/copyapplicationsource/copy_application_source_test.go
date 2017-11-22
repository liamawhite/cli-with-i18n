package copyapplicationsource_test

import (
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/liamawhite/cli-with-i18n/cf/api/apifakes"
	. "github.com/liamawhite/cli-with-i18n/cf/api/copyapplicationsource"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/net"
	"github.com/liamawhite/cli-with-i18n/cf/terminal/terminalfakes"
	testconfig "github.com/liamawhite/cli-with-i18n/util/testhelpers/configuration"
	testnet "github.com/liamawhite/cli-with-i18n/util/testhelpers/net"

	"github.com/liamawhite/cli-with-i18n/cf/trace/tracefakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CopyApplicationSource", func() {
	var (
		repo       Repository
		testServer *httptest.Server
		configRepo coreconfig.ReadWriter
	)

	setupTestServer := func(reqs ...testnet.TestRequest) {
		testServer, _ = testnet.NewServer(reqs)
		configRepo.SetAPIEndpoint(testServer.URL)
	}

	BeforeEach(func() {
		configRepo = testconfig.NewRepositoryWithDefaults()
		gateway := net.NewCloudControllerGateway(configRepo, time.Now, new(terminalfakes.FakeUI), new(tracefakes.FakePrinter), "")
		repo = NewCloudControllerCopyApplicationSourceRepository(configRepo, gateway)
	})

	AfterEach(func() {
		testServer.Close()
	})

	Describe(".CopyApplication", func() {
		BeforeEach(func() {
			setupTestServer(apifakes.NewCloudControllerTestRequest(testnet.TestRequest{
				Method: "POST",
				Path:   "/v2/apps/target-app-guid/copy_bits",
				Matcher: testnet.RequestBodyMatcher(`{
					"source_app_guid": "source-app-guid"
				}`),
				Response: testnet.TestResponse{
					Status: http.StatusCreated,
				},
			}))
		})

		It("should return a CopyApplicationModel", func() {
			err := repo.CopyApplication("source-app-guid", "target-app-guid")
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
