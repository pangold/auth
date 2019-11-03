package utils

import (
	"errors"

	"./config"
)

var (
	conf config.Config
	system *config.SystemConfig
)

func getConfig() config.Config {
	if conf == nil {
		conf = config.UseYaml()
	}
	return conf
}

func ReadConfig(path string) error {
	sys, err := getConfig().ReadConfig(path)
	if err != nil {
		return errors.New("cannot read configuration file")
	}
	system = sys
	return nil
}

func GetConfig() *config.SystemConfig {
	if system == nil {
		panic("configuration havn't been loaded")
	}
	return system
}

