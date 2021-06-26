package config

import (
	"embed"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/JakBaranowski/gb-tools/common"
	"github.com/JakBaranowski/gb-tools/fileops"
)

//go:embed config.json
var embededConfig embed.FS

// Config is a struct used to hold the contents of the configuration json file.
type Config struct {
	GamePath string
	Missions struct {
		RelativePath     string
		StringsToReplace map[string]string
	}
	Loadouts []struct {
		Name                    string
		SourceRelativePath      string
		DestinationRelativePath string
	}
}

// Parses embeded config file and returns a Config struct
func ParseEmbededConfig() (conf Config) {
	configFile, err := embededConfig.ReadFile("config.json")
	common.Must(err)
	json.Unmarshal(configFile, &conf)
	return
}

// Parses json formatted config file and returns a Config struct
func ParseFileConfig(configPath string) (conf Config) {
	configFile := fileops.OpenAndReadFile(configPath)
	json.Unmarshal(configFile, &conf)
	return
}

// Checks if a config file is present in the gbt config directory. If there is none
// will use the default embeded config.
func ParseConfig() (conf Config) {
	userConfigDir, err := os.UserConfigDir()
	common.Must(err)
	configFilePath := filepath.Join(userConfigDir, "gbt", "gbt.conf")
	if fileops.DoesExist(configFilePath) {
		log.Printf("Config found")
		conf = ParseFileConfig(configFilePath)
	} else {
		conf = ParseEmbededConfig()
	}
	return
}

// Saves the default config to gbt config directory. It can be then edited by hand
// to include custom settings.
func SaveConfig() {
	userConfigDir, err := os.UserConfigDir()
	common.Must(err)
	configDirPath := filepath.Join(userConfigDir, "gbt")
	fileops.CreateDirIfDoesntExist(configDirPath, 0755)
	configFilePath := filepath.Join(configDirPath, "gbt.conf")
	configFile, err := embededConfig.ReadFile("config.json")
	common.Must(err)
	file, err := os.Create(configFilePath)
	common.Must(err)
	defer file.Close()
	n, err := file.Write(configFile)
	log.Printf("Written %d bytes to %s", n, configFilePath)
	common.Must(err)
}
