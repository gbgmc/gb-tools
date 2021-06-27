package command

import "github.com/JakBaranowski/gb-tools/config"

func Config(conf config.Config) {
	config.SaveConfig(conf)
}
