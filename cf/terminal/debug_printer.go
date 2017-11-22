package terminal

import (
	"time"

	. "github.com/liamawhite/cli-with-i18n/cf/i18n"
	"github.com/liamawhite/cli-with-i18n/cf/trace"
)

type DebugPrinter struct {
	Logger trace.Printer
}

func (p DebugPrinter) Print(title, dump string) {
	p.Logger.Printf("\n%s [%s]\n%s\n", HeaderColor(T(title)), time.Now().Format(time.RFC3339), trace.Sanitize(dump))
}
