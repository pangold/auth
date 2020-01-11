package api

import (
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/controller"
	"gitlab.com/pangold/auth/middleware"
	"gitlab.com/pangold/auth/middleware/system"
	"gitlab.com/pangold/auth/model"
	"gitlab.com/pangold/auth/service"
)

type Server struct {
	config config.Server
	token middleware.Token
	cache middleware.Cache
	email middleware.Email
	vcode middleware.VerificationCode
	auth *model.Auth
	user *model.User
}

func NewServer(conf config.Config) *Server {
	return &Server{
		config: conf.Server,
		token: system.NewDefaultToken(conf.Jwt.SecretKey),
		cache: system.NewSimpleCache(),
		email: system.NewDefaultEmail(conf.Email), // FIXME: should as one
		vcode: system.NewDefaultEmail(conf.Email),
		auth: model.NewAuth(conf.MySQL),
		user: &model.User{},
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
	as := service.NewAccount(this.config, this.auth, this.email, this.vcode, this.token, this.cache)
	us := service.NewUser(this.config, this.user, this.cache)
	ac := controller.NewAccountController(as)
	uc := controller.NewUserController(us)
	auth := controller.NewAuthController(this.token)
	//
	r := NewRouter(this.config)
	r.AccountRouter(ac)
	r.AccountV2Router(ac)
	r.UserRouter(uc, auth.Filter)
	r.Run()
}
