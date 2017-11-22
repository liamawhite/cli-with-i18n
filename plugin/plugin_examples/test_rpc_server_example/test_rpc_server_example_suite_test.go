package main_test

import (
	"github.com/liamawhite/cli-with-i18n/util/testhelpers/pluginbuilder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTestRpcServerExample(t *testing.T) {
	RegisterFailHandler(Fail)

	pluginbuilder.BuildTestBinary("", "test_rpc_server_example")

	RunSpecs(t, "Test RPC Server Example Suite")
}
