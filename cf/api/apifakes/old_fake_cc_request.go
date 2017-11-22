package apifakes

import (
	"net/http"

	testnet "github.com/liamawhite/cli-with-i18n/util/testhelpers/net"
)

func NewCloudControllerTestRequest(request testnet.TestRequest) testnet.TestRequest {
	request.Header = http.Header{
		"accept":        {"application/json"},
		"authorization": {"BEARER my_access_token"},
	}

	return request
}
