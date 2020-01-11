package main

import (
	"gitlab.com/pangold/auth/api"
	"gitlab.com/pangold/auth/config"
)

func main() {
	conf := config.NewYaml("config/config.yml").ReadConfig()
	server := api.NewServer(*conf)
	// custom middleware
	// server.UseToken(system.NewDefaultToken("secret-key"))
	// server.UseCache(system.NewRedisCache(conf.Redis.Addr))
	server.Run()
}
