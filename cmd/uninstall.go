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
		projectManifestPath, err := filepath.Abs(args[0])
		cobra.CheckErr(err)
		projectManifest := common.ParseManifest(projectManifestPath)
		log.Printf("Started uninstalling %s", projectManifest.Name)
		files := projectManifest.GetFiles()
		dirs := make(map[string]bool)
		gameDir := viper.GetString("gameDir")
		for _, file := range files {
			targetFile := filepath.Join(gameDir, file)
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
		if installManifest {
			log.Printf("Finished uninstalling %s", projectManifest.Name)
			return
		}
		log.Print("Removing install manifest")
		if filepath.Dir(projectManifestPath) == gameDir {
			common.Remove(projectManifestPath)
		} else if installManifestPath := filepath.Join(
			gameDir,
			fmt.Sprintf("%s.json", projectManifest.Name),
		); common.DoesExist(installManifestPath) {
			common.Remove(installManifestPath)
		}
		log.Printf("Finished uninstalling %s", projectManifest.Name)
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)

	uninstallCmd.Flags().BoolVarP(
		&installManifest,
		"manifest",
		"m",
		false,
		"If set will not remove post installation manifest from game directory. Post installation manifest contains a list of exact files copied to game directory.",
	)
}
