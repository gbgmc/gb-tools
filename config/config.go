package config

import (
	"encoding/json"

	"github.com/JakBaranowski/gb-tools/helpers"
)

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

func (conf *Config) ParseConfig(configPath string) {
	configFile := helpers.OpenAndReadFile(configPath)
	json.Unmarshal(configFile, conf)
}
