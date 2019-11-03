package config

import (
	"fmt"
	"testing"
)

func TestWriteConfig(t *testing.T) {
	c := UseYaml()
	conf := SystemConfig{Server: Server{Port: 8888}}
	if err := c.WriteConfig("config2.yml", conf); err != nil {
		t.Errorf(err.Error())
	}
	conf2, err := c.ReadConfig("config2.yml")
	if err != nil {
		t.Errorf(err.Error())
	}
	if conf.Server.Port != conf2.Server.Port {
		t.Errorf("unexptected port")
	}
}
