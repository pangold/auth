package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/controller"
	"log"
)

type Router struct {
	config   config.HttpConfig
	router  *gin.Engine
}

func NewRouter(conf config.HttpConfig) *Router {
	return &Router{
		config: conf,
		router: gin.Default(),
	}
}

func (this *Router) Run() {
	log.Printf("http account server start running at %s\n", this.config.Addr)
	if err := this.router.Run(this.config.Addr); err != nil {
		panic(err)
	}
}

func (this *Router) AccountRouter(account *controller.AccountController) {
	a := this.router.Group("/api/v1/accounts")
	a.POST("/sign_up",          account.SignUp)
	a.POST("/activation_url",   account.GetActivationUrl)
	a.POST("/activate",         account.Activate)
	a.POST("/sign_in",          account.SignIn)
	a.POST("/sign_out",         account.SignOut)
	a.POST("/forgot",           account.Forgot)
	a.POST("/reset",            account.Reset)
	a.POST("/is_user_id_exist", account.IsUserIdExist)
	a.POST("/is_email_exist",   account.IsEmailExist)
	a.POST("/is_phone_exist",   account.IsPhoneExist)
}

func (this *Router) AccountV2Router(account *controller.AccountController) {
	a := this.router.Group("/api/v2/accounts")
	a.POST("/get_code",     account.GetVCode)
	a.POST("/sign_up",      account.SignUpWithVCode)
	a.POST("/sign_in",      account.SignInWithVCode)
	a.POST("/reset",        account.ResetWithVCode)    // v2 of forgot
	a.POST("/lock",         account.Lock)
	a.POST("/unlock",       account.Unlock)
	a.POST("/bind_email",   account.BindEmail)
	a.POST("/unbind_email", account.UnbindEmail)
	a.POST("/bind_phone",   account.BindPhone)
	a.POST("/unbind_phone", account.UnbindPhone)
	// a.POST("/3rd",          account.ThirdPartyLogin)
}

func (this *Router) UserRouter(filter *controller.Filter, user *controller.UserController) {
	u := this.router.Group("/api/v1").Use(filter.Filter)
	u.GET   ("/users",     user.Index)
	u.GET   ("/user/{id}", user.Show)
	u.POST  ("/user/",     user.Create)
	u.PUT   ("/user/{id}", user.Update)
	u.DELETE("/user/{id}", user.Delete)
}


