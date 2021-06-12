package command

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/JakBaranowski/gb-tools/config"
	"github.com/JakBaranowski/gb-tools/fileops"
)

func Loadout(config config.Config) {
	fmt.Printf("Started mirroring AI loadouts.\n")
	fileops.CreateDirIfDoesntExist(config.Loadouts.DestinationRelativePath, 0755)
	removeRedundantMirrorFiles(config)
	mirrorFiles(config)
	fmt.Printf("Finished mirroring AI loadouts.\n")
}

func removeRedundantMirrorFiles(config config.Config) {
	fmt.Printf("Removing redundant files from mirror directory.\n")
	dstGlobPattern := filepath.Join(config.Loadouts.DestinationRelativePath, "*.kit")
	dstFileList, err := filepath.Glob(dstGlobPattern)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range dstFileList {
		srcFilePath := filepath.Join(
			config.Loadouts.SourceRelativePath,
			filepath.Base(file),
		)
		if !fileops.DoesExist(srcFilePath) {
			fmt.Printf(
				"Removing mirror file \"%s\", as the original no longer exists.\n",
				file,
			)
			err = os.Remove(file)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func mirrorFiles(config config.Config) {
	fmt.Printf("Copying orginal files to mirror directory.\n")
	srcGlobPattern := filepath.Join(config.Loadouts.SourceRelativePath, "*.kit")
	srcFileList, err := filepath.Glob(srcGlobPattern)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range srcFileList {
		dstFilePath := filepath.Join(
			config.Loadouts.DestinationRelativePath,
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
