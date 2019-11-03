package config

import (
	"gopkg.in/yaml.v2"
)

type Yaml struct {

}

func UseYaml() Config {
	return &Yaml{}
}

func (y Yaml) ReadConfig(path string) (*SystemConfig, error) {
	conf := SystemConfig{}
	data, err := ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal([]byte(data), &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

func (y Yaml) WriteConfig(path string, conf SystemConfig) error {
	data, err := yaml.Marshal(&conf)
	if err != nil {
		return err
	}
	if err := WriteFile(path, string(data)); err != nil {
		return err
	}
	return nil
}
