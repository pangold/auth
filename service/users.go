package service

import (
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/middleware"
	"gitlab.com/pangold/auth/model"
)

type UserService struct {
	config config.Server
	db *model.DB
	cache middleware.Cache
}

func NewUserService(conf config.Server, db *model.DB, c middleware.Cache) *UserService {
	return &UserService{
		config: conf,
		db: db,
		cache: c,
	}
}