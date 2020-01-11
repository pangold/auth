package service

import (
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/middleware"
	"gitlab.com/pangold/auth/model/db"
)

type Account struct {
	config config.Server
	db *db.Account
	cache middleware.Cache
}

func NewAccountService(conf config.Config, c middleware.Cache) *Account {
	return &Account{
		config: conf.Server,
		db: db.NewAccount(conf.MySQL),
		cache: c,
	}
}