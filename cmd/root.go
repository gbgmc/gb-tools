/*
Copyright © 2023 Jakub Baranowski jbaranowski@rubberduckling.dev

*/
package cmd

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	gameDir string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gbt",
	Short: "Ground Branch Tools aim to make Ground Branch Game Mode development easier",
	Long: `Ground Branch Tools automate a lot of repetetive tasks that are required when 
developing Game Modes for Ground Branch. For example to pack all files required 
by a game mode you would first create a manifest file for it and then use gbt 
to pack all files into one archive using the pack comand:

gbt pack GBGMC.json`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is \"./gbt.json\")")
	rootCmd.PersistentFlags().StringVar(&gameDir, "gameDir", "C:/Program Files (x86)/Steam/steamapps/common/Ground Branch", "game installation directory")
	cobra.CheckErr(viper.BindPFlag("gameDir", rootCmd.PersistentFlags().Lookup("gameDir")))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		workingDir, err := os.Getwd()
		cobra.CheckErr(err)
		viper.AddConfigPath(workingDir)
		viper.SetConfigType("json")
		viper.SetConfigName("gbt")
	}

	if err := viper.ReadInConfig(); err == nil {
		log.Printf("Using config file: %s.", viper.ConfigFileUsed())
	} else if errors.As(err, &viper.ConfigFileNotFoundError{}) {
		log.Print("Config file not found. Using defaults.")
	} else {
		cobra.CheckErr(err)
	}
}
