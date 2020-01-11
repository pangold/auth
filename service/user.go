package service

import (
	"database/sql"
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/middleware"
	"gitlab.com/pangold/auth/model"
	"gitlab.com/pangold/auth/model/db"
)

type User struct {
	config config.Server
	db *db.User
	cache middleware.Cache
}

func NewUserService(conf config.Config, conn *sql.DB, c middleware.Cache) *User {
	return &User{
		config: conf.Server,
		db: db.NewUser(conn),
		cache: c,
	}
}

func (this *User) GetUsers() []*model.User {
	return this.db.GetUsers()
}

func (this *User) GetUser(id uint64) *model.User {
	return this.db.GetUserById(id)
}

func (this *User) Create(u *model.User) error {
	return this.db.Create(u)
}

func (this *User) Update(u *model.User) error {
	return this.db.Update(u)
}

func (this *User) Delete(u *model.User) error {
	return this.db.Delete(u)
}