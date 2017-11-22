package pushaction

import "github.com/liamawhite/cli-with-i18n/util/manifest"

func (*Actor) ReadManifest(pathToManifest string) ([]manifest.Application, error) {
	// Cover method to make testing easier
	return manifest.ReadAndMergeManifests(pathToManifest)
}
