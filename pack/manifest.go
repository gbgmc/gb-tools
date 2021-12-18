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

// Creates a new Manifest struct with the provided data.
func NewManifest(
	name string,
	version string,
	dependencies []string,
	files []string,
) (manifest Manifest) {
	manifest.Name = name
	manifest.Version = version
	manifest.Dependencies = dependencies
	manifest.Files = files
	return manifest
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

// Saves the manifest under specified path.
func (manifest *Manifest) Save(path string) {
	jsonData, _ := json.MarshalIndent(manifest, "", "    ")
	fileops.Save(jsonData, path)
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
		common.Must(err)
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
