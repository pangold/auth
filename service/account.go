package service

import (
	"database/sql"
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/middleware"
	"gitlab.com/pangold/auth/model"
	"gitlab.com/pangold/auth/model/db"
)

type Account struct {
	config config.Server
	db *db.Account
	cache middleware.Cache
}

func NewAccountService(conf config.Config, conn *sql.DB, c middleware.Cache) *Account {
	return &Account{
		config: conf.Server,
		db: db.NewAccount(conn),
		cache: c,
	}
}

func (this *Account) GetAccounts() []*model.Account {
	return this.db.GetAccounts()
}

func (this *Account) GetAccount(id uint64) *model.Account {
	return this.db.GetAccountById(id)
}

func (this *Account) Create(a *model.Account) error {
	return this.db.Create(a)
}

func (this *Account) Update(a *model.Account) error {
	return this.db.Update(a)
}

func (this *Account) Delete(a *model.Account) error {
	return this.db.Delete(a)
}