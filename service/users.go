package service

import (
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/middleware"
	"gitlab.com/pangold/auth/model"
)

type UserService struct {
	db *model.DB
	cache middleware.Cache
}

func NewUserService(conf config.Config, c middleware.Cache) *UserService {
	return &UserService{
		db: model.NewDB(conf.MySQL),
		cache: c,
	}
}