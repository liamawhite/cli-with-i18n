package plugininstaller

import (
	"github.com/liamawhite/cli-with-i18n/cf/actors/pluginrepo"
	"github.com/liamawhite/cli-with-i18n/cf/models"
	"github.com/liamawhite/cli-with-i18n/cf/terminal"
	"github.com/liamawhite/cli-with-i18n/util"
	"github.com/liamawhite/cli-with-i18n/util/downloader"
)

//go:generate counterfeiter . PluginInstaller

type PluginInstaller interface {
	Install(inputSourceFilepath string) string
}

type Context struct {
	Checksummer    util.Sha1Checksum
	FileDownloader downloader.Downloader
	GetPluginRepos pluginReposFetcher
	PluginRepo     pluginrepo.PluginRepo
	RepoName       string
	UI             terminal.UI
}

type pluginReposFetcher func() []models.PluginRepo

func NewPluginInstaller(context *Context) PluginInstaller {
	var installer PluginInstaller

	pluginDownloader := &PluginDownloader{UI: context.UI, FileDownloader: context.FileDownloader}
	if context.RepoName == "" {
		installer = &pluginInstallerWithoutRepo{
			UI:               context.UI,
			PluginDownloader: pluginDownloader,
			RepoName:         context.RepoName,
		}
	} else {
		installer = &pluginInstallerWithRepo{
			UI:               context.UI,
			PluginDownloader: pluginDownloader,
			RepoName:         context.RepoName,
			Checksummer:      context.Checksummer,
			PluginRepo:       context.PluginRepo,
			GetPluginRepos:   context.GetPluginRepos,
		}
	}
	return installer
}
