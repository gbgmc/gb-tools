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
		dirs := make(map[string]bool)
		for _, file := range files {
			targetFile := filepath.Join(viper.GetString("gameDir"), file)
			if common.DoesExist(targetFile) {
				common.Remove(targetFile)
				log.Printf("Removed file %s ", targetFile)
			}
			dir := filepath.Dir(targetFile)
			if _, value := dirs[dir]; !value {
				dirs[dir] = true
			}
		}
		for dir := range dirs {
			if common.DoesExist(dir) {
				common.Remove(dir)
			}
		}
		log.Printf("Finished uninstalling %s", manifest.Name)
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
