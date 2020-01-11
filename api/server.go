package api

import (
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/controller"
	"gitlab.com/pangold/auth/middleware"
	"gitlab.com/pangold/auth/middleware/system"
	"gitlab.com/pangold/auth/service"
)

type Server struct {
	config config.Config
	token middleware.Token
	cache middleware.Cache
	email middleware.Email
	vcode middleware.VerificationCode
}

func NewServer(conf config.Config) *Server {
	return &Server{
		config: conf,
		token:  system.NewDefaultToken(conf.Jwt.SecretKey),
		cache:  system.NewSimpleCache(),
		email:  system.NewDefaultEmail(conf.Email), // FIXME: should as one
		vcode:  system.NewDefaultEmail(conf.Email),
	}
}

// customize token middleware
func (this *Server) UseToken(token middleware.Token) {
	this.token = token
}

// customize cache middleware
func (this *Server) UseCache(cache middleware.Cache) {
	this.cache = cache
}

func (this *Server) UseEmail(email middleware.Email) {
	this.email = email
}

func (this *Server) UseVCode(vcode middleware.VerificationCode) {
	this.vcode = vcode
}

func (this *Server) Run() {
	as := service.NewAuthService(this.config, this.email, this.vcode, this.token, this.cache)
	us := service.NewUserService(this.config, this.cache)
	a3 := service.NewAccountService(this.config, this.cache)
	ac := controller.NewAuthController(as)
	uc := controller.NewUserController(us)
	a4 := controller.NewAccountController(a3)
	f := controller.NewFilterController(this.token)
	//
	r := NewRouter(this.config.Server)
	r.AuthRouter(ac)
	r.Auth2Router(ac)
	r.AccountRouter(a4, f.Filter)
	r.UserRouter(uc, f.Filter)
	r.Run()
}
