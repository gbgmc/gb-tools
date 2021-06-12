package main

import (
	"fmt"
	"log"
	"os"

	"github.com/JakBaranowski/gb-tools/command"
	"github.com/JakBaranowski/gb-tools/config"
)

func main() {
	var config config.Config
	config.ParseConfig()
	commandString, err := command.GetArgument(1)
	if err != nil {
		log.Fatal(err)
	}
	switch commandString {
	case "loadout":
		command.Loadout(config)
	case "mission":
		command.Mission(config)
	default:
		fmt.Printf("Unsopported command \"%s\".", os.Args[1])
	}
}
