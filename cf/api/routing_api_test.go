package api_test

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/liamawhite/cli-with-i18n/cf/api"
	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/errors"
	"github.com/liamawhite/cli-with-i18n/cf/models"
	"github.com/liamawhite/cli-with-i18n/cf/net"
	"github.com/liamawhite/cli-with-i18n/cf/terminal/terminalfakes"
	"github.com/liamawhite/cli-with-i18n/cf/trace/tracefakes"
	testconfig "github.com/liamawhite/cli-with-i18n/util/testhelpers/configuration"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("RoutingApi", func() {

	var (
		repo             api.RoutingAPIRepository
		configRepo       coreconfig.Repository
		routingAPIServer *ghttp.Server
	)

	BeforeEach(func() {
		configRepo = testconfig.NewRepositoryWithDefaults()
		gateway := net.NewRoutingAPIGateway(configRepo, time.Now, new(terminalfakes.FakeUI), new(tracefakes.FakePrinter), "")

		repo = api.NewRoutingAPIRepository(configRepo, gateway)
	})

	AfterEach(func() {
		routingAPIServer.Close()
	})

	Describe("ListRouterGroups", func() {

		Context("when routing api return router groups", func() {
			BeforeEach(func() {
				routingAPIServer = ghttp.NewServer()
				routingAPIServer.RouteToHandler("GET", "/v1/router_groups",
					func(w http.ResponseWriter, req *http.Request) {
						responseBody := []byte(`[
			{
				  "guid": "bad25cff-9332-48a6-8603-b619858e7992",
					"name": "default-tcp",
					"type": "tcp"
			}]`)
						w.Header().Set("Content-Length", strconv.Itoa(len(responseBody)))
						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(http.StatusOK)
						w.Write(responseBody)
					})
				configRepo.SetRoutingAPIEndpoint(routingAPIServer.URL())
			})

			It("lists routing groups", func() {
				cb := func(grp models.RouterGroup) bool {
					Expect(grp).To(Equal(models.RouterGroup{
						GUID: "bad25cff-9332-48a6-8603-b619858e7992",
						Name: "default-tcp",
						Type: "tcp",
					}))
					return true
				}
				err := repo.ListRouterGroups(cb)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when routing api returns an empty response ", func() {
			BeforeEach(func() {
				routingAPIServer = ghttp.NewServer()
				routingAPIServer.RouteToHandler("GET", "/v1/router_groups",
					func(w http.ResponseWriter, req *http.Request) {
						responseBody := []byte("[]")
						w.Header().Set("Content-Length", strconv.Itoa(len(responseBody)))
						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(http.StatusOK)
						w.Write(responseBody)
					})
				configRepo.SetRoutingAPIEndpoint(routingAPIServer.URL())
			})

			It("doesn't list any router groups", func() {
				cb := func(grp models.RouterGroup) bool {
					Fail(fmt.Sprintf("Not expected to receive callback for router group:%#v", grp))
					return false
				}
				err := repo.ListRouterGroups(cb)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when routing api returns an error ", func() {
			BeforeEach(func() {
				routingAPIServer = ghttp.NewServer()
				routingAPIServer.RouteToHandler("GET", "/v1/router_groups",
					func(w http.ResponseWriter, req *http.Request) {
						w.WriteHeader(http.StatusUnauthorized)
						w.Write([]byte(`{"name":"UnauthorizedError","message":"token is expired"}`))
					})
				configRepo.SetRoutingAPIEndpoint(routingAPIServer.URL())
			})

			It("doesn't list any router groups and displays error message", func() {
				cb := func(grp models.RouterGroup) bool {
					Fail(fmt.Sprintf("Not expected to receive callback for router group:%#v", grp))
					return false
				}

				err := repo.ListRouterGroups(cb)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("token is expired"))
				Expect(err.(errors.HTTPError).ErrorCode()).To(ContainSubstring("UnauthorizedError"))
				Expect(err.(errors.HTTPError).StatusCode()).To(Equal(http.StatusUnauthorized))
			})
		})
	})
})
