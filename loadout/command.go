package loadout

import (
	"log"
	"path"
	"path/filepath"
	"strings"

	"github.com/JakBaranowski/gb-tools/common"
	"github.com/JakBaranowski/gb-tools/config"
	"github.com/JakBaranowski/gb-tools/fileops"
)

// CommandLoadout will glob all files in original directories specified in
// config file, mirror all files from original directories in counterpart
// directories, and then remove files not present in the original directories
// from the counterpart directories.
func CommandLoadout(conf config.Config) {
	argument, exist := common.GetOptionalArgument(2)
	if exist {
		mirrorLoadoutsFromPath(argument)
	} else {
		mirrorLoadoutsFromConfig(conf)
	}
}

// mirrorLoadoutsFromPath mirrors directory specified as an argument passed by
// the user.
func mirrorLoadoutsFromPath(sourceRelativePath string) {
	destinationRelativePath := createPrefixedDirPath(sourceRelativePath, "_")
	if destinationRelativePath == "" {
		return
	}
	log.Printf(
		"Mirroring AI loadouts '%s'.",
		path.Base(sourceRelativePath),
	)
	mirrorLoadout(
		sourceRelativePath,
		destinationRelativePath,
	)
	log.Printf(
		"Finished mirroring AI loadouts '%s'.",
		path.Base(sourceRelativePath),
	)
}

// createPrefixedDirPath is a helper function that takes a originalPath and
// returns a new path almost identical to the originalPath but with a prefix
// added at the beginning of last directory name.
func createPrefixedDirPath(originalPath string, prefix string) string {
	separator := ""
	if strings.Contains(originalPath, "\\") {
		separator = "\\"
	} else {
		separator = "/"
	}
	i := strings.LastIndex(originalPath, separator)
	if i < 0 {
		return ""
	}
	pathArray := strings.Split(strings.Trim(originalPath, separator), separator)
	cleanPath := path.Join(pathArray[:len(pathArray)-1]...)
	return path.Join(cleanPath, prefix+pathArray[len(pathArray)-1])
}

// mirrorLoadoutsFromConfig mirrors directories specified in the tool config.
func mirrorLoadoutsFromConfig(conf config.Config) {
	for i := range conf.Loadouts {
		log.Printf(
			"Mirroring AI loadouts '%s'.",
			conf.Loadouts[i].Name,
		)
		mirrorLoadout(
			conf.Loadouts[i].SourceRelativePath,
			conf.Loadouts[i].DestinationRelativePath,
		)
		log.Printf(
			"Finished mirroring AI loadouts '%s'.",
			conf.Loadouts[i].Name,
		)
	}
}

// mirrorLoadout is an "passthrough" function that calls functions in order
// to synchronize original and mirror directories.
func mirrorLoadout(source string, destination string) {
	fileops.CreateDirIfDoesntExist(
		destination,
		0755,
	)
	removeRedundantFiles(
		source,
		destination,
	)
	mirrorNewFiles(
		source,
		destination,
	)
}

// removeRedundantFiles makes sure that the mirror directory does not contain
// files that don't have counterparts in the original directory.
func removeRedundantFiles(sorucePath string, destinationPath string) {
	destinationGlobPattern := filepath.Join(destinationPath, "*.kit")
	destinationFileList, err := filepath.Glob(destinationGlobPattern)
	common.Must(err)
	for _, file := range destinationFileList {
		srcFilePath := filepath.Join(
			sorucePath,
			filepath.Base(file),
		)
		if !fileops.DoesExist(srcFilePath) {
			log.Printf(
				"Removing mirrored file '%s', as the original no longer exists.",
				file,
			)
			fileops.RemoveFile(file)
		}
	}
}

// mirrorNewFiles makes sure the new files from the original directory have
// counterparts in the mirror directory.
func mirrorNewFiles(sourcePath string, destinationPath string) {
	sourceGlobPattern := filepath.Join(sourcePath, "*.kit")
	sourceFileList, err := filepath.Glob(sourceGlobPattern)
	common.Must(err)
	for _, file := range sourceFileList {
		dstFilePath := filepath.Join(
			destinationPath,
			filepath.Base(file),
		)
		if !fileops.DoesExist(dstFilePath) {
			log.Printf(
				"Mirroring file '%s'.",
				file,
			)
			fileops.Touch(dstFilePath)
		}
	}
}
