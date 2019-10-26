package config

import (
	"gopkg.in/yaml.v2"
)

type Yaml struct {
	Path string
}

func UseYaml(path string) Config {
	return &Yaml{Path: path}
}

func (y Yaml) ReadConfig() (*SystemConfig, error) {
	conf := SystemConfig{}
	data, err := ReadFile(y.Path)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal([]byte(data), &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

func (y Yaml) WriteConfig(conf SystemConfig) error {
	data, err := yaml.Marshal(&conf)
	if err != nil {
		return err
	}
	if err := WriteFile(y.Path, string(data)); err != nil {
		return err
	}
	return nil
}
