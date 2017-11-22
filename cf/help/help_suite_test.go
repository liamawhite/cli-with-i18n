package help_test

import (
	"github.com/liamawhite/cli-with-i18n/cf/commandsloader"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHelp(t *testing.T) {
	RegisterFailHandler(Fail)

	commandsloader.Load()

	RunSpecs(t, "Help Suite")
}
