package service

import (
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/middleware"
	"gitlab.com/pangold/auth/model/db"
)

type User struct {
	config config.Server
	user *db.User
	cache middleware.Cache
}

func NewUserService(conf config.Server, user *db.User, c middleware.Cache) *User {
	return &User{
		config: conf,
		user: user,
		cache: c,
	}
}