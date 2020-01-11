package service

import (
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/middleware"
	"gitlab.com/pangold/auth/model"
)

type User struct {
	config config.Server
	user *model.User
	cache middleware.Cache
}

func NewUser(conf config.Server, user *model.User, c middleware.Cache) *User {
	return &User{
		config: conf,
		user: user,
		cache: c,
	}
}