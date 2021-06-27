package command

import (
	"log"
	"os"
	"path/filepath"

	"github.com/JakBaranowski/gb-tools/common"
	"github.com/JakBaranowski/gb-tools/config"
	"github.com/JakBaranowski/gb-tools/fileops"
)

// Loadout will glob all files in original directories specified in config file,
// mirror all files from original directories in counterpart directories, and
// then remove files not present in the original directories from the
// counterpart directories.
func Loadout(config config.Config) {
	for i := range config.Loadouts {
		log.Printf(
			"Mirroring AI loadouts %s.",
			config.Loadouts[i].Name,
		)
		fileops.CreateDirIfDoesntExist(
			config.Loadouts[i].DestinationRelativePath,
			0755,
		)
		removeRedundantMirrorFiles(
			config.Loadouts[i].SourceRelativePath,
			config.Loadouts[i].DestinationRelativePath,
		)
		mirrorFiles(
			config.Loadouts[i].SourceRelativePath,
			config.Loadouts[i].DestinationRelativePath,
		)
		log.Printf(
			"Finished mirroring AI loadouts %s.",
			config.Loadouts[i].Name,
		)
	}
}

func removeRedundantMirrorFiles(srcPath string, dstPath string) {
	dstGlobPattern := filepath.Join(dstPath, "*.kit")
	dstFileList, err := filepath.Glob(dstGlobPattern)
	common.Must(err)

	for _, file := range dstFileList {
		srcFilePath := filepath.Join(
			srcPath,
			filepath.Base(file),
		)
		if !fileops.DoesExist(srcFilePath) {
			log.Printf(
				"Removing mirrored file \"%s\", as the original no longer exists.",
				file,
			)
			err = os.Remove(file)
			common.Must(err)
		}
	}
}

func mirrorFiles(srcPath string, dstPath string) {
	srcGlobPattern := filepath.Join(srcPath, "*.kit")
	srcFileList, err := filepath.Glob(srcGlobPattern)
	common.Must(err)

	for _, file := range srcFileList {
		dstFilePath := filepath.Join(
			dstPath,
			filepath.Base(file),
		)
		if !fileops.DoesExist(dstFilePath) {
			log.Printf(
				"Mirroring file \"%s\".",
				file,
			)
			fileops.Touch(dstFilePath)
		}
	}
}
