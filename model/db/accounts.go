package db

import (
	"database/sql"
	"fmt"
	"gitlab.com/pangold/auth/config"
)

type Account struct {
	db *sql.DB
}

func NewAccount(c config.MySQL) *Account {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", c.User, c.Password, c.Host, c.DBName))
	if err != nil {
		panic(err.Error())
	}
	return &Account{
		db: db,
	}
}