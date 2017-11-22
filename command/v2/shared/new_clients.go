package shared

import (
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccv2"
	ccWrapper "github.com/liamawhite/cli-with-i18n/api/cloudcontroller/wrapper"
	"github.com/liamawhite/cli-with-i18n/api/uaa"
	uaaWrapper "github.com/liamawhite/cli-with-i18n/api/uaa/wrapper"
	"github.com/liamawhite/cli-with-i18n/command"
	"github.com/liamawhite/cli-with-i18n/command/translatableerror"
)

// NewClients creates a new V2 Cloud Controller client and UAA client using the
// passed in config.
func NewClients(config command.Config, ui command.UI, targetCF bool) (*ccv2.Client, *uaa.Client, error) {
	ccWrappers := []ccv2.ConnectionWrapper{}

	verbose, location := config.Verbose()
	if verbose {
		ccWrappers = append(ccWrappers, ccWrapper.NewRequestLogger(ui.RequestLoggerTerminalDisplay()))
	}

	if location != nil {
		ccWrappers = append(ccWrappers, ccWrapper.NewRequestLogger(ui.RequestLoggerFileWriter(location)))
	}

	authWrapper := ccWrapper.NewUAAAuthentication(nil, config)

	ccWrappers = append(ccWrappers, authWrapper)
	ccWrappers = append(ccWrappers, ccWrapper.NewRetryRequest(2))

	ccClient := ccv2.NewClient(ccv2.Config{
		AppName:            config.BinaryName(),
		AppVersion:         config.BinaryVersion(),
		JobPollingTimeout:  config.OverallPollingTimeout(),
		JobPollingInterval: config.PollingInterval(),
		Wrappers:           ccWrappers,
	})

	if !targetCF {
		return ccClient, nil, nil
	}

	if config.Target() == "" {
		return nil, nil, translatableerror.NoAPISetError{
			BinaryName: config.BinaryName(),
		}
	}

	_, err := ccClient.TargetCF(ccv2.TargetSettings{
		URL:               config.Target(),
		SkipSSLValidation: config.SkipSSLValidation(),
		DialTimeout:       config.DialTimeout(),
	})
	if err != nil {
		return nil, nil, HandleError(err)
	}

	if ccClient.AuthorizationEndpoint() == "" {
		return nil, nil, translatableerror.AuthorizationEndpointNotFoundError{}
	}

	uaaClient := uaa.NewClient(uaa.Config{
		AppName:           config.BinaryName(),
		AppVersion:        config.BinaryVersion(),
		ClientID:          config.UAAOAuthClient(),
		ClientSecret:      config.UAAOAuthClientSecret(),
		DialTimeout:       config.DialTimeout(),
		SkipSSLValidation: config.SkipSSLValidation(),
	})

	if verbose {
		uaaClient.WrapConnection(uaaWrapper.NewRequestLogger(ui.RequestLoggerTerminalDisplay()))
	}
	if location != nil {
		uaaClient.WrapConnection(uaaWrapper.NewRequestLogger(ui.RequestLoggerFileWriter(location)))
	}

	uaaAuthWrapper := uaaWrapper.NewUAAAuthentication(nil, config)
	uaaClient.WrapConnection(uaaAuthWrapper)
	uaaClient.WrapConnection(uaaWrapper.NewRetryRequest(2))

	err = uaaClient.SetupResources(config, ccClient.AuthorizationEndpoint())
	if err != nil {
		return nil, nil, err
	}

	uaaAuthWrapper.SetClient(uaaClient)
	authWrapper.SetClient(uaaClient)

	return ccClient, uaaClient, err
}
