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
		manifestPath := args[0]
		manifest := common.ParseManifest(manifestPath)
		log.Printf("Started installing %s", manifest.Name)
		files := manifest.GetFiles()
		for _, file := range files {
			targetFile := filepath.Join(viper.GetString("gameDir"), file)
			common.CreateDirIfDoesntExist(filepath.Dir(targetFile), 0755)
			log.Printf("Copying file %s to %s ", file, targetFile)
			common.Copy(file, targetFile)
		}
		log.Printf("Finished installing %s", manifest.Name)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
