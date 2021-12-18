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
	files := manifest.GetFiles()
	packageName := manifest.Name + ".zip"
	fileops.CompressFiles(packageName, files)
}

// Parses the game mode manifest file under the provided path and moves matching files
// to the provided game path.
// The manifest file can have any extension but has to be json formatted.
func CommandInstall(conf config.Config) {
	manifestPath, err := common.GetRequiredArgument(2, "Expected path to manifest file")
	common.Must(err)
	gamePath, err := getGamePath(conf)
	common.Must(err)
	manifest := ParseManifest(manifestPath)
	log.Printf("Started installing %s", manifest.Name)
	files := manifest.GetFiles()
	for _, file := range files {
		targetFile := filepath.Join(gamePath, file)
		fileops.CreateDirIfDoesntExist(filepath.Dir(targetFile), 0755)
		log.Printf("Copying file %s to %s ", file, targetFile)
		fileops.Copy(file, targetFile)
	}
	log.Printf("Finished installing %s", manifest.Name)
}

// Parses the game mode manifest file under the provided path and removes matching files
// from the provided game path.
// The manifest file can have any extension but has to be json formatted.
func CommandUninstall(conf config.Config) {
	manifestPath, err := common.GetRequiredArgument(2, "Expected path to manifest file")
	common.Must(err)
	gamePath, err := getGamePath(conf)
	common.Must(err)
	manifest := ParseManifest(manifestPath)
	log.Printf("Started uninstalling %s", manifest.Name)
	files := manifest.GetFiles()
	for _, file := range files {
		targetFile := filepath.Join(gamePath, file)
		log.Printf("Removing file %s ", targetFile)
		if fileops.DoesExist(targetFile) {
			fileops.RemoveFile(targetFile)
		}
	}
	log.Printf("Finished uninstalling %s", manifest.Name)
}

func getGamePath(conf config.Config) (string, error) {
	gamePath := ""
	err := error(nil)
	if gamePathName, exist := common.GetOptionalArgument(3); exist {
		gamePath, err = conf.GetGamePath(gamePathName)
	} else {
		gamePath, err = conf.GetGamePath("default")
	}
	return gamePath, err
}
