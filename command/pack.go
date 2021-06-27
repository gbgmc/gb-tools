package command

import (
	"github.com/JakBaranowski/gb-tools/common"
	"github.com/JakBaranowski/gb-tools/fileops"
	"github.com/JakBaranowski/gb-tools/gamemode"
)

// Parses the game mode manifest file under the provided path and packages it
// into easy to use zip files.
// The manifest file can have any extension but has to be json formatted.
func Pack() {
	manifestPath, err := common.GetArgument(2)
	common.Must(err)
	pack(manifestPath)
}

func pack(manifestPath string) {
	manifest := gamemode.ParseManifest(manifestPath)
	packageName := manifest.Name + ".zip"
	fileops.CompressFiles(packageName, manifest.Files)
}
