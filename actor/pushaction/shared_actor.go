package pushaction

import "github.com/liamawhite/cli-with-i18n/actor/sharedaction"

//go:generate counterfeiter . SharedActor

type SharedActor interface {
	GatherArchiveResources(archivePath string) ([]sharedaction.Resource, error)
	GatherDirectoryResources(sourceDir string) ([]sharedaction.Resource, error)
	ZipArchiveResources(sourceArchivePath string, filesToInclude []sharedaction.Resource) (string, error)
	ZipDirectoryResources(sourceDir string, filesToInclude []sharedaction.Resource) (string, error)
}
