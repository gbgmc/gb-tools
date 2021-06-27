package gamemode

import (
	"encoding/json"

	"github.com/JakBaranowski/gb-tools/fileops"
)

type Manifest struct {
	Name    string
	Version string
	Files   []string
}

// Parses the manifest file under the provided manifestPath. Returns manifest
// with parsed manifest values.
func ParseManifest(manifestPath string) (manifest Manifest) {
	manifestFile := fileops.OpenAndReadFile(manifestPath)
	json.Unmarshal(manifestFile, &manifest)
	return
}
