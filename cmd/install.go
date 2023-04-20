/*
Copyright © 2023 Jakub Baranowski jbaranowski@rubberduckling.dev

*/
package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/JakBaranowski/gb-tools/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install [manifestFilePath]",
	Short: "Moves files specified in manifest to Ground Branch installation folder",
	Long: `Parses the provided game mode manifest file and moves matching files to
the provided game directory. The manifest file can have any extension but
has to be json formatted.Usage:

gbt install GBGMC.json -c gbt.conf`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalf("Missing required argument \"manifestFilePath\"")
		}
		projectManifestPath := args[0]
		projectManifest := common.ParseManifest(projectManifestPath)
		log.Printf("Started installing %s", projectManifest.Name)
		files := projectManifest.GetFiles()
		gameDir := viper.GetString("gameDir")
		for _, file := range files {
			targetFile := filepath.Join(gameDir, file)
			common.CreateDirIfDoesntExist(filepath.Dir(targetFile), 0755)
			log.Printf("Copying file %s to %s ", file, targetFile)
			common.Copy(file, targetFile)
		}
		log.Printf("Finished installing %s", projectManifest.Name)
		if !createInstallManifest {
			return
		}
		installManifest := common.Manifest{
			Name:         projectManifest.Name,
			Version:      projectManifest.Version,
			Dependencies: nil,
			Files:        files,
		}
		installManifestPath := filepath.Join(
			gameDir,
			fmt.Sprintf("%s.json", installManifest.Name),
		)
		installManifest.Save(installManifestPath)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().BoolVarP(
		&createInstallManifest,
		"createPostManifest",
		"m",
		false,
		"Creates post installation manifest in game directory. Post installation manifest contains a list of exact files copied to game directory.",
	)
}
