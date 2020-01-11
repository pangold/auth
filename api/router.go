package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/controller"
)

type Router struct {
	config   config.Server
	router  *gin.Engine
}

func NewRouter(conf config.Server) *Router {
	return &Router{
		config: conf,
		router: gin.Default(),
	}
}

func (this *Router) Run() {
	if err := this.router.Run(this.config.Addr); err != nil {
		panic(err)
	}
}

func (this *Router) AuthRouter(auth *controller.AuthController) {
	a := this.router.Group("/api/v1")
	a.POST("/sign_up",        auth.SignUp)
	a.POST("/activation_url", auth.GetActivationUrl)
	a.POST("/activate",       auth.Activate)
	a.POST("/sign_in",        auth.SignIn)
	a.POST("/sign_out",       auth.SignOut)
	a.POST("/forgot",         auth.Forgot)
	a.POST("/reset",          auth.Reset)
	a.POST("/uid_exist",      auth.IsUserIdExist)
	a.POST("/email_exist",    auth.IsEmailExist)
	a.POST("/phone_exist",    auth.IsPhoneExist)
}

func (this *Router) Auth2Router(auth *controller.AuthController) {
	a := this.router.Group("/api/v2")
	a.POST("/get_code",       auth.GetVCode)
	a.POST("/sign_up",        auth.SignUpWithVCode)
	a.POST("/sign_in",        auth.SignInWithVCode)
	a.POST("/forgot",         auth.ResetWithVCode)    // v2 of forgot
	a.POST("/reset",          auth.ResetWithVCode)    // v2 of forgot
	a.POST("/lock",           auth.Lock)
	a.POST("/unlock",         auth.Unlock)
	a.POST("/bind_email",     auth.BindEmail)
	a.POST("/unbind_email",   auth.UnbindEmail)
	a.POST("/bind_phone",     auth.BindPhone)
	a.POST("/unbind_phone",   auth.UnbindPhone)
	// a.POST("/3rd",          account.ThirdPartyLogin)
}

func (this *Router) AccountRouter(account *controller.AccountController, filter ...gin.HandlerFunc) {
	a := this.router.Group("/api/v1").Use(filter...)
	a.GET   ("/accounts",     account.Index)
	a.GET   ("/account/{id}", account.Show)
	a.POST  ("/account/",     account.Create)
	a.PUT   ("/account/{id}", account.Update)
	a.DELETE("/account/{id}", account.Delete)
}

func (this *Router) UserRouter(user *controller.UserController, filter ...gin.HandlerFunc) {
	u := this.router.Group("/api/v1").Use(filter...)
	u.GET   ("/users",     user.Index)
	u.GET   ("/user/{id}", user.Show)
	u.POST  ("/user/",     user.Create)
	u.PUT   ("/user/{id}", user.Update)
	u.DELETE("/user/{id}", user.Delete)
}


