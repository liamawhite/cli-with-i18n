package requirements

import (
	"errors"

	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/terminal"
)

type LoginRequirement struct {
	config                 coreconfig.Reader
	apiEndpointRequirement APIEndpointRequirement
}

func NewLoginRequirement(config coreconfig.Reader) LoginRequirement {
	return LoginRequirement{config, APIEndpointRequirement{config}}
}

func (req LoginRequirement) Execute() error {

	if err := req.apiEndpointRequirement.Execute(); err != nil {
		return err
	}

	if !req.config.IsLoggedIn() {
		return errors.New(terminal.NotLoggedInText())
	}

	return nil
}
