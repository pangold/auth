package db

import (
	"database/sql"
	"fmt"
	"gitlab.com/pangold/auth/config"
)

type User struct {
	db *sql.DB
}

func NewUser(c config.MySQL) *User {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", c.User, c.Password, c.Host, c.DBName))
	if err != nil {
		panic(err.Error())
	}
	return &User{
		db: db,
	}
}
