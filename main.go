package main

import (
	"fmt"

	"github.com/JakBaranowski/gb-tools/command"
	"github.com/JakBaranowski/gb-tools/common"
	"github.com/JakBaranowski/gb-tools/config"
)

func main() {
	conf := config.ParseConfig()
	commandString, err := common.GetArgument(1)
	common.Must(err)
	switch commandString {
	case "loadout":
		command.Loadout(conf)
	case "pack":
		command.Pack()
	case "config":
		command.Config(conf)
	default:
		fmt.Printf("Unsopported command \"%s\".", commandString)
	}
}
