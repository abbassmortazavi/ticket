package config

import "ticket/config"

func Get() config.Config {
	return configurations
}
