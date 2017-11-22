package appfiles_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/liamawhite/cli-with-i18n/cf/api/apifakes"
	"github.com/liamawhite/cli-with-i18n/cf/net"
	"github.com/liamawhite/cli-with-i18n/cf/terminal/terminalfakes"
	testconfig "github.com/liamawhite/cli-with-i18n/util/testhelpers/configuration"
	testnet "github.com/liamawhite/cli-with-i18n/util/testhelpers/net"

	. "github.com/liamawhite/cli-with-i18n/cf/api/appfiles"
	"github.com/liamawhite/cli-with-i18n/cf/trace/tracefakes"
	. "github.com/liamawhite/cli-with-i18n/util/testhelpers/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AppFilesRepository", func() {
	It("lists files", func() {
		expectedResponse := "file 1\n file 2\n file 3"

		listFilesEndpoint := func(writer http.ResponseWriter, request *http.Request) {
			methodMatches := request.Method == "GET"
			pathMatches := request.URL.Path == "/some/path"

			if !methodMatches || !pathMatches {
				fmt.Printf("One of the matchers did not match. Method [%t] Path [%t]",
					methodMatches, pathMatches)

				writer.WriteHeader(http.StatusInternalServerError)
				return
			}

			writer.WriteHeader(http.StatusOK)
			fmt.Fprint(writer, expectedResponse)
		}

		listFilesServer := httptest.NewServer(http.HandlerFunc(listFilesEndpoint))
		defer listFilesServer.Close()

		req := apifakes.NewCloudControllerTestRequest(testnet.TestRequest{
			Method: "GET",
			Path:   "/v2/apps/my-app-guid/instances/1/files/some/path",
			Response: testnet.TestResponse{
				Status: http.StatusTemporaryRedirect,
				Header: http.Header{
					"Location": {fmt.Sprintf("%s/some/path", listFilesServer.URL)},
				},
			},
		})

		listFilesRedirectServer, handler := testnet.NewServer([]testnet.TestRequest{req})
		defer listFilesRedirectServer.Close()

		configRepo := testconfig.NewRepositoryWithDefaults()
		configRepo.SetAPIEndpoint(listFilesRedirectServer.URL)

		gateway := net.NewCloudControllerGateway(configRepo, time.Now, new(terminalfakes.FakeUI), new(tracefakes.FakePrinter), "")
		repo := NewCloudControllerAppFilesRepository(configRepo, gateway)
		list, err := repo.ListFiles("my-app-guid", 1, "some/path")

		Expect(handler).To(HaveAllRequestsCalled())
		Expect(err).ToNot(HaveOccurred())
		Expect(list).To(Equal(expectedResponse))
	})
})
