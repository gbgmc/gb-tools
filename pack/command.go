package pack

import (
	"log"
	"path/filepath"

	"github.com/JakBaranowski/gb-tools/common"
	"github.com/JakBaranowski/gb-tools/config"
	"github.com/JakBaranowski/gb-tools/fileops"
)

// Parses the game mode manifest file under the provided path and packages it
// into easy to use zip files.
// The manifest file can have any extension but has to be json formatted.
func CommandPack() {
	manifestPath, err := common.GetRequiredArgument(2, "Expected path to manifest file")
	common.Must(err)
	manifest := ParseManifest(manifestPath)
	files := GetFiles(manifest)
	packageName := manifest.Name + ".zip"
	fileops.CompressFiles(packageName, files)
}

func CommandInstall(conf config.Config) {
	manifestPath, err := common.GetRequiredArgument(2, "Expected path to manifest file")
	common.Must(err)
	manifest := ParseManifest(manifestPath)
	log.Printf("Started installing %s", manifest.Name)
	files := GetFiles(manifest)
	for _, file := range files {
		targetFile := filepath.Join(conf.GamePath, file)
		fileops.CreateDirIfDoesntExist(filepath.Dir(targetFile), 0755)
		log.Printf("Copying file %s to %s ", file, targetFile)
		fileops.Copy(file, targetFile)
	}
	log.Printf("Finished installing %s", manifest.Name)
}

func CommandUninstall(conf config.Config) {
	manifestPath, err := common.GetRequiredArgument(2, "Expected path to manifest file")
	common.Must(err)
	manifest := ParseManifest(manifestPath)
	log.Printf("Started uninstalling %s", manifest.Name)
	files := GetFiles(manifest)
	for _, file := range files {
		targetFile := filepath.Join(conf.GamePath, file)
		log.Printf("Removing file %s ", targetFile)
		if fileops.DoesExist(targetFile) {
			fileops.RemoveFile(targetFile)
		}
	}
	log.Printf("Finished uninstalling %s", manifest.Name)
}
