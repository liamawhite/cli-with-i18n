package net

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/liamawhite/cli-with-i18n/cf/configuration/coreconfig"
	"github.com/liamawhite/cli-with-i18n/cf/errors"
	"github.com/liamawhite/cli-with-i18n/cf/terminal"
	"github.com/liamawhite/cli-with-i18n/cf/trace"
)

type errorResponse struct {
	Name    string
	Message string
}

func errorHandler(statusCode int, body []byte) error {
	response := errorResponse{}
	err := json.Unmarshal(body, &response)
	if err != nil {
		return errors.NewHTTPError(http.StatusInternalServerError, "", "")
	}

	return errors.NewHTTPError(statusCode, response.Name, response.Message)
}

func NewRoutingAPIGateway(config coreconfig.Reader, clock func() time.Time, ui terminal.UI, logger trace.Printer, envDialTimeout string) Gateway {
	return Gateway{
		errHandler:      errorHandler,
		config:          config,
		PollingThrottle: DefaultPollingThrottle,
		warnings:        &[]string{},
		Clock:           clock,
		ui:              ui,
		logger:          logger,
		PollingEnabled:  true,
		DialTimeout:     dialTimeout(envDialTimeout),
	}
}
