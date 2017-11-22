package shared

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/cloudfoundry/noaa/consumer"
	"github.com/liamawhite/cli-with-i18n/api/uaa"
	"github.com/liamawhite/cli-with-i18n/api/uaa/noaabridge"
	"github.com/liamawhite/cli-with-i18n/command"
)

type RequestLoggerOutput interface {
	Start() error
	Stop() error
	DisplayType(name string, requestDate time.Time) error
	DisplayDump(dump string) error
}

type DebugPrinter struct {
	outputs []RequestLoggerOutput
}

func (p *DebugPrinter) addOutput(output RequestLoggerOutput) {
	p.outputs = append(p.outputs, output)
}

func (p DebugPrinter) Print(title string, dump string) {
	for _, output := range p.outputs {
		_ = output.Start()
		defer output.Stop()

		output.DisplayType(title, time.Now())
		output.DisplayDump(dump)
	}

}

// NewNOAAClient returns back a configured NOAA Client.
func NewNOAAClient(apiURL string, config command.Config, uaaClient *uaa.Client, ui command.UI) *consumer.Consumer {
	client := consumer.New(
		apiURL,
		&tls.Config{
			InsecureSkipVerify: config.SkipSSLValidation(),
		},
		http.ProxyFromEnvironment,
	)
	client.RefreshTokenFrom(noaabridge.NewTokenRefresher(uaaClient, config))
	client.SetMaxRetryCount(5)

	noaaDebugPrinter := DebugPrinter{}

	// if verbose, set debug printer on noaa client
	verbose, location := config.Verbose()

	client.SetDebugPrinter(&noaaDebugPrinter)

	if verbose {
		noaaDebugPrinter.addOutput(ui.RequestLoggerTerminalDisplay())
	}
	if location != nil {
		noaaDebugPrinter.addOutput(ui.RequestLoggerFileWriter(location))
	}

	return client
}
