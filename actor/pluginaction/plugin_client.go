package pluginaction

import "github.com/liamawhite/cli-with-i18n/api/plugin"

//go:generate counterfeiter . PluginClient

type PluginClient interface {
	GetPluginRepository(repositoryURL string) (plugin.PluginRepository, error)
	DownloadPlugin(pluginURL string, path string, proxyReader plugin.ProxyReader) error
}
