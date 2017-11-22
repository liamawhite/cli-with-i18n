package actors_test

import (
	"github.com/liamawhite/cli-with-i18n/cf/i18n"
	"github.com/liamawhite/cli-with-i18n/util/testhelpers/configuration"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestActors(t *testing.T) {
	i18n.T = i18n.Init(configuration.NewRepositoryWithDefaults())
	RegisterFailHandler(Fail)
	RunSpecs(t, "Actors Suite")
}
