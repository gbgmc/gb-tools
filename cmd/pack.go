/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/JakBaranowski/gb-tools/common"
	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack [manifestFilePath]",
	Short: "Packs all files specified in manifest into an archive",
	Long: `Parses the provided game mode manifest file and packages it into easy to
use archives. The manifest file can have any extension but has to be json
formatted. Usage:

gbt pack GBGMC.json`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalf("Missing required argument \"manifestFilePath\"")
		}
		projectManifestPath := args[0]
		projectManifest := common.ParseManifest(projectManifestPath)
		files := projectManifest.GetFiles()
		packageName := projectManifest.Name + ".zip"
		common.CompressFiles(packageName, files)
	},
}

func init() {
	rootCmd.AddCommand(packCmd)
}
