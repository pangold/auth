package config

import (
	"fmt"
	"testing"
)

func TestReadConfig(t *testing.T) {
	c := UseYaml("config.yml")
	conf, err := c.ReadConfig()
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(conf)
}

func TestWriteConfig(t *testing.T) {
	c := UseYaml("config2.yml")
	conf := SystemConfig{Server: Server{Port: 8888}}
	if err := c.WriteConfig(conf); err != nil {
		t.Errorf(err.Error())
	}
	conf2, err := c.ReadConfig()
	if err != nil {
		t.Errorf(err.Error())
	}
	if conf.Server.Port != conf2.Server.Port {
		t.Errorf("unexptected port")
	}
}
