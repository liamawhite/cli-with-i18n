package errors

import (
	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
)

type InvalidSSLCert struct {
	URL    string
	Reason string
}

func NewInvalidSSLCert(url, reason string) *InvalidSSLCert {
	return &InvalidSSLCert{
		URL:    url,
		Reason: reason,
	}
}

func (err *InvalidSSLCert) Error() string {
	message := T("Received invalid SSL certificate from ") + err.URL
	if err.Reason != "" {
		message += " - " + err.Reason
	}
	return message
}
