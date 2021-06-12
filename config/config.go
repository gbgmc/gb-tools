package config

import (
	"embed"
	"encoding/json"
	"log"
)

//go:embed config.json
var embededConfig embed.FS

// Config is a struct used to hold the contents of the configuration json file.
type Config struct {
	Missions Mission
	Loadouts Loadout
}

type Mission struct {
	RelativePath     string
	StringsToReplace map[string]string
}

type Loadout struct {
	SourceRelativePath      string
	DestinationRelativePath string
}

func (conf *Config) ParseConfig() {
	configFile, err := embededConfig.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(configFile, conf)
}
