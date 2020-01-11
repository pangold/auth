package api

import (
	"database/sql"
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/controller"
	"gitlab.com/pangold/auth/middleware"
	"gitlab.com/pangold/auth/middleware/system"
	"gitlab.com/pangold/auth/service"
	"gitlab.com/pangold/auth/utils"
)

type Server struct {
	config config.Config
	token middleware.Token
	cache middleware.Cache
	email middleware.Email
	vcode middleware.VerificationCode
	dbconn *sql.DB
}

func NewDBConn(conf config.MySQL) *sql.DB {
	return utils.GetDBConn(conf.User, conf.Password, conf.Host, conf.DBName, conf.Port)
}

func NewServer(conf config.Config) *Server {
	conn := NewDBConn(conf.MySQL)
	email := system.NewDefaultEmail(conf.Email)
	return &Server{
		config: conf,
		token:  system.NewDefaultToken(conf.Jwt.SecretKey),
		cache:  system.NewSimpleCache(),
		email:  email,
		vcode:  email,
		dbconn: conn,
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
	as := service.NewAuthService(this.config, this.dbconn, this.email, this.vcode, this.token, this.cache)
	us := service.NewUserService(this.config, this.dbconn, this.cache)
	a3 := service.NewAccountService(this.config, this.dbconn, this.cache)
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
