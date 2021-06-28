package pack

import (
	"github.com/JakBaranowski/gb-tools/common"
	"github.com/JakBaranowski/gb-tools/fileops"
)

// Parses the game mode manifest file under the provided path and packages it
// into easy to use zip files.
// The manifest file can have any extension but has to be json formatted.
func CommandPack() {
	manifestPath, err := common.GetRequiredArgument(2, "Expected path to manifest file")
	common.Must(err)
	manifest := ParseManifest(manifestPath)
	packageName := manifest.Name + ".zip"
	fileops.CompressFiles(packageName, manifest.Files)
}
