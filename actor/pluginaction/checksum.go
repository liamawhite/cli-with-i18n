package pluginaction

import "github.com/liamawhite/cli-with-i18n/util/configv3"

func (actor Actor) ValidateFileChecksum(path string, checksum string) bool {
	plugin := configv3.Plugin{Location: path}
	return plugin.CalculateSHA1() == checksum
}
