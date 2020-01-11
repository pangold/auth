package api

import (
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/controller"
	"gitlab.com/pangold/auth/middleware"
	"gitlab.com/pangold/auth/middleware/system"
	"gitlab.com/pangold/auth/model/db"
	"gitlab.com/pangold/auth/service"
)

type Server struct {
	config config.Server
	token middleware.Token
	cache middleware.Cache
	email middleware.Email
	vcode middleware.VerificationCode
	auth *db.Auth
	user *db.User
}

func NewServer(conf config.Config) *Server {
	return &Server{
		config: conf.Server,
		token:  system.NewDefaultToken(conf.Jwt.SecretKey),
		cache:  system.NewSimpleCache(),
		email:  system.NewDefaultEmail(conf.Email), // FIXME: should as one
		vcode:  system.NewDefaultEmail(conf.Email),
		auth:   db.NewAuth(conf.MySQL),
		user:   &db.User{},
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
	as := service.NewAuthService(this.config, this.auth, this.email, this.vcode, this.token, this.cache)
	us := service.NewUserService(this.config, this.user, this.cache)
	ac := controller.NewAuthController(as)
	uc := controller.NewUserController(us)
	auth := controller.NewFilterController(this.token)
	//
	r := NewRouter(this.config)
	r.AccountRouter(ac)
	r.AccountV2Router(ac)
	r.UserRouter(uc, auth.Filter)
	r.Run()
}