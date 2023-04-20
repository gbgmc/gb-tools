package config

import (
	"embed"
	"encoding/json"
	"errors"
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
	GamePaths []struct {
		Name string
		Path string
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
		conf = ParseFileConfig(configFilePath)
	} else {
		conf = ParseEmbededConfig()
	}
	return
}

// Saves the default config to gbt config directory. It can be then edited by hand
// to include custom settings.
func SaveConfig(conf Config) {
	userConfigDir, err := os.UserConfigDir()
	common.Must(err)
	configDirPath := filepath.Join(userConfigDir, "gbt")
	fileops.CreateDirIfDoesntExist(configDirPath, 0755)
	configFilePath := filepath.Join(configDirPath, "gbt.conf")
	if fileops.DoesExist(configFilePath) {
		log.Printf("Config file already exists.")
		return
	}
	configFileContent, err := json.MarshalIndent(conf, "", "  ")
	common.Must(err)
	err = os.WriteFile(configFilePath, configFileContent, 0666)
	common.Must(err)
	log.Printf("Config file saved in '%s'.", configFilePath)
}

// Returns the game path with the provided name from the config. If the game path
// with the given name does not exists returns nil.
func (conf *Config) GetGamePath(name string) (string, error) {
	for _, gamePath := range conf.GamePaths {
		if gamePath.Name == name {
			return gamePath.Path, nil
		}
	}
	return "", errors.New("game path not found")
}
