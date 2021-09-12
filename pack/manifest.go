package pack

import (
	"encoding/json"
	"log"
	"path/filepath"

	"github.com/JakBaranowski/gb-tools/common"
	"github.com/JakBaranowski/gb-tools/fileops"
)

// Struct for parsing json formatted game mode manifests.
type Manifest struct {
	Name         string
	Version      string
	Dependencies []string
	Globs        []string
	Files        []string
}

// Parses the manifest file under the provided manifestPath. Returns manifest
// with parsed manifest values.
func ParseManifest(manifestPath string) (manifest Manifest) {
	log.Println("Opening manifest " + manifestPath)
	manifestFile := fileops.OpenAndReadFile(manifestPath)
	log.Println("Parsing manifest " + manifestPath)
	err := json.Unmarshal(manifestFile, &manifest)
	common.Must(err)
	return
}

// Gets all files
func GetFiles(manifest Manifest) (filesList []string) {
	log.Println("Getting files from manifest")
	for _, dependency := range manifest.Dependencies {
		dependencyFiles := GetFiles(ParseManifest(dependency))
		filesList = append(filesList, dependencyFiles...)
	}
	for _, glob := range manifest.Globs {
		matches, err := filepath.Glob(glob)
		common.Must(err)
		filesList = append(filesList, matches...)
	}
	filesList = append(filesList, manifest.Files...)
	log.Println("Found files")
	for _, file := range filesList {
		log.Println(file)
	}
	filesList = removeDuplicateFiles(filesList)
	return
}

func removeDuplicateFiles(filesList []string) []string {
	log.Println("Removing duplicated files")
	foundFiles := make(map[string]bool)
	cleanFileList := []string{}
	for _, file := range filesList {
		if _, value := foundFiles[file]; !value {
			foundFiles[file] = true
			cleanFileList = append(cleanFileList, file)
		}
	}
	return cleanFileList
}
