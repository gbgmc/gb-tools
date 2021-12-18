package main

import (
	"log"

	"github.com/JakBaranowski/gb-tools/common"
	"github.com/JakBaranowski/gb-tools/config"
	"github.com/JakBaranowski/gb-tools/pack"
)

func main() {
	conf := config.ParseConfig()
	commandString, err := common.GetRequiredArgument(
		1,
		"Supported commands are: 'loadout', 'pack' and 'config'",
	)
	common.Must(err)
	switch commandString {
	case "pack":
		pack.CommandPack()
	case "install":
		pack.CommandInstall(conf)
	case "uninstall":
		pack.CommandUninstall(conf)
	case "config":
		config.CommandConfig(conf)
	default:
		log.Printf(
			`Unsupported command \"%s\". Supported commands are: 'pack', 
			'install', 'uninstall' and 'config'`,
			commandString,
		)
	}
}
