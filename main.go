package main

import (
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/controller"
	"gitlab.com/pangold/auth/middleware/system"
	"gitlab.com/pangold/auth/service"
)

func main() {
	// loadConfig("config.yml")
	as := service.AccountService{}
	us := service.UserService{}
	ac := controller.NewAccountController(&as)
	uc := controller.NewUserController(&us)
	fc := controller.NewFilterController(system.NewDefaultToken("my-secret-key"))
	//
	router := NewRouter(config.HttpConfig{Addr: "0.0.0.0:9999"})
	router.AccountRouter(ac)
	router.AccountV2Router(ac)
	router.UserRouter(fc, uc)
	router.Run()
}
