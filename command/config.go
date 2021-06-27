package command

import "github.com/JakBaranowski/gb-tools/config"

// Config method will save the config so it can be edited by the user.
func Config(conf config.Config) {
	config.SaveConfig(conf)
}
