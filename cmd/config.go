/*
Copyright © 2023 Jakub Baranowski jbaranowski@rubberduckling.dev

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Saves Ground Branch tools config file to working directory",
	Long: `Saves Ground Branch tools config file to working directory for editing. 
Config file contains Ground Branch installation folder, used for install and
uninstall commands. Usage:
	
gbt config`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(viper.SafeWriteConfig())
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
