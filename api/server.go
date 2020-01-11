package api

import (
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/controller"
	"gitlab.com/pangold/auth/middleware"
	"gitlab.com/pangold/auth/middleware/system"
	"gitlab.com/pangold/auth/service"
)

type Server struct {
	conf config.Config
	token middleware.Token
	cache middleware.Cache
}

func NewServer(conf config.Config) *Server {
	server := &Server{
		conf: conf,
		token: system.NewDefaultToken(conf.Jwt.SecretKey),
		cache: system.NewSimpleCache(),
	}
	return server
}

//
func (this *Server) UseToken(token middleware.Token) {
	this.token = token
}

func (this *Server) UseCache(cache middleware.Cache) {
	this.cache = cache
}

func (this *Server) Run() {
	as := service.NewAccountService(this.conf, this.token, this.cache)
	us := service.NewUserService(this.conf, this.cache)
	ac := controller.NewAccountController(as)
	uc := controller.NewUserController(us)
	fc := controller.NewFilterController(this.token)
	//
	r := NewRouter(this.conf.Server)
	r.AccountRouter(ac)
	r.AccountV2Router(ac)
	r.UserRouter(fc, uc)
	r.Run()
}
