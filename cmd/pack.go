/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"path/filepath"

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
		projectManifestPath, err := filepath.Abs(args[0])
		cobra.CheckErr(err)
		projectManifest := common.ParseManifest(projectManifestPath)
		files := projectManifest.GetFiles()
		bodies := make(map[string][]byte)
		if installManifest {
			log.Print("Adding install manifest to archive")
			installManifest := common.Manifest{
				Name:         projectManifest.Name,
				Version:      projectManifest.Version,
				Dependencies: nil,
				Files:        files,
			}
			bodies[fmt.Sprintf("%s.json", installManifest.Name)] = installManifest.Marshal()
		}
		packageName := projectManifest.Name + ".zip"
		common.Compress(packageName, files, bodies)
	},
}

func init() {
	rootCmd.AddCommand(packCmd)

	packCmd.Flags().BoolVarP(
		&installManifest,
		"manifest",
		"m",
		false,
		"If set will add post installation manifest to the archive. Post installation manifest contains a list of exact files copied to game directory.",
	)
}
