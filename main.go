package main

import (
	"github.com/JakBaranowski/gb-tools/config"
)

func main() {
	var config config.Config
	println("Parsing")
	config.ParseConfig("config.json")
	println(config.Loadouts.SourceRelativePath)
	println(config)
	for key, value := range config.Missions.StringsToReplace {
		print(key, value)
	}
	// commandString, err := command.CheckForArgument(1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// switch commandString {
	// case "mirror":
	// 	command.Mirror(config)
	// case "mission":
	// 	command.Mission(config)
	// default:
	// 	fmt.Printf("Unsopported command %s.", os.Args[1])
	// }
}
