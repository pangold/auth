package service

import (
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/middleware"
	"gitlab.com/pangold/auth/model/db"
)

type User struct {
	config config.Server
	db *db.User
	cache middleware.Cache
}

func NewUserService(conf config.Config, c middleware.Cache) *User {
	return &User{
		config: conf.Server,
		db: db.NewUser(conf.MySQL),
		cache: c,
	}
}