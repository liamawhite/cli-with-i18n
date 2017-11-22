package plugin_test

import (
	"path/filepath"

	"github.com/liamawhite/cli-with-i18n/util/testhelpers/pluginbuilder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPlugin(t *testing.T) {
	RegisterFailHandler(Fail)
	pluginbuilder.BuildTestBinary(filepath.Join("..", "fixtures", "plugins"), "test_1")
	RunSpecs(t, "Plugin Suite")
}
