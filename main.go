package main

import (
	"fmt"

	"github.com/JakBaranowski/gb-tools/command"
	"github.com/JakBaranowski/gb-tools/common"
	"github.com/JakBaranowski/gb-tools/config"
)

func main() {
	Config := config.ParseConfig()
	commandString, err := common.GetArgument(1)
	common.Must(err)
	switch commandString {
	case "loadout":
		command.Loadout(Config)
	case "mission":
		command.Mission(Config)
	case "pack":
		command.Pack()
	case "config":
		command.Config()
	default:
		fmt.Printf("Unsopported command \"%s\".", commandString)
	}
}
