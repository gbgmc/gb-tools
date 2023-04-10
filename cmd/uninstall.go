/*
Copyright © 2023 Jakub Baranowski jbaranowski@rubberduckling.dev

*/
package cmd

import (
	"log"
	"path/filepath"

	"github.com/JakBaranowski/gb-tools/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall [manifestFilePath]",
	Short: "Removes files specified in manifest from Ground Branch installation folder",
	Long: `Parses the game mode manifest file under the provided path and removes
matching files from the provided game path. The manifest file can have 
any extension but has to be json formatted. Usage:

gbt uninstall GBGMC.json -c gbt.conf`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalf("Missing required argument \"manifestFilePath\"")
		}
		manifestPath := args[0]
		manifest := common.ParseManifest(manifestPath)
		log.Printf("Started uninstalling %s", manifest.Name)
		files := manifest.GetFiles()
		for _, file := range files {
			targetFile := filepath.Join(viper.GetString("gameDir"), file)
			log.Printf("Removing file %s ", targetFile)
			if common.DoesExist(targetFile) {
				common.RemoveFile(targetFile)
			}
		}
		log.Printf("Finished uninstalling %s", manifest.Name)
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
