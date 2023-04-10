package common

import (
	"encoding/json"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
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
	manifestFile := OpenAndReadFile(manifestPath)
	log.Println("Parsing manifest " + manifestPath)
	err := json.Unmarshal(manifestFile, &manifest)
	cobra.CheckErr(err)
	return
}

// Gets all files from a manifest
func (manifest *Manifest) GetFiles() (filesList []string) {
	log.Println("Getting files from manifest")
	for _, dependencyPath := range manifest.Dependencies {
		dependency := ParseManifest(dependencyPath)
		dependencyFiles := dependency.GetFiles()
		filesList = append(filesList, dependencyFiles...)
	}
	for _, file := range manifest.Files {
		matches, err := filepath.Glob(file)
		cobra.CheckErr(err)
		filesList = append(filesList, matches...)
	}
	filesList = normalizeSlashes(removeDuplicateFiles(filesList))
	return
}

// Normalizes slashes in paths from the provided fileList in order to avoid
// issues with paths on different OS.
func normalizeSlashes(filesList []string) []string {
	log.Println("Normalizing slashes in file paths")
	filesListNormalized := []string{}
	for _, file := range filesList {
		filesListNormalized = append(filesListNormalized, filepath.ToSlash(file))
	}
	return filesListNormalized
}

// Removes duplicated entries from a provided fileList.
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
