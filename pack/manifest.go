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

// Gets all files from a manifest
func GetFiles(manifest Manifest) (filesList []string) {
	log.Println("Getting files from manifest")
	for _, dependency := range manifest.Dependencies {
		dependencyFiles := GetFiles(ParseManifest(dependency))
		filesList = append(filesList, dependencyFiles...)
	}
	for _, file := range manifest.Files {
		matches, err := filepath.Glob(file)
		common.Must(err)
		filesList = append(filesList, matches...)
	}
	filesList = normalizeSlashes(removeDuplicateFiles(filesList))
	return
}

func normalizeSlashes(filesList []string) []string {
	log.Println("Normalizing slashes in file paths")
	filesListNormalized := []string{}
	for _, file := range filesList {
		filesListNormalized = append(filesListNormalized, filepath.ToSlash(file))
	}
	return filesListNormalized
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
