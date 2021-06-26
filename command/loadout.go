package command

import (
	"fmt"
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
		fmt.Printf(
			"Started mirroring AI loadouts %s.\n",
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
		fmt.Printf(
			"Finished mirroring AI loadouts %s.\n",
			config.Loadouts[i].Name,
		)
	}
}

func removeRedundantMirrorFiles(srcPath string, dstPath string) {
	fmt.Printf("Removing redundant files from mirror directory.\n")
	dstGlobPattern := filepath.Join(dstPath, "*.kit")
	dstFileList, err := filepath.Glob(dstGlobPattern)
	common.Must(err)

	for _, file := range dstFileList {
		srcFilePath := filepath.Join(
			srcPath,
			filepath.Base(file),
		)
		if !fileops.DoesExist(srcFilePath) {
			fmt.Printf(
				"Removing mirror file \"%s\", as the original no longer exists.\n",
				file,
			)
			err = os.Remove(file)
			common.Must(err)
		}
	}
}

func mirrorFiles(srcPath string, dstPath string) {
	fmt.Printf("Copying orginal files to mirror directory.\n")
	srcGlobPattern := filepath.Join(srcPath, "*.kit")
	srcFileList, err := filepath.Glob(srcGlobPattern)
	common.Must(err)

	for _, file := range srcFileList {
		dstFilePath := filepath.Join(
			dstPath,
			filepath.Base(file),
		)
		if !fileops.DoesExist(dstFilePath) {
			fmt.Printf(
				"Mirroring file \"%s\".\n",
				file,
			)
			fileops.Touch(dstFilePath)
		}
	}
}
