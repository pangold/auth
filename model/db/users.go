package db

import (
	"database/sql"
	"fmt"
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/model"
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

//ID uint64
//UserId string
//NickName string
//Profile string
//Sex bool
//Age int
//Hobby string

func (this *User) GetUsers() (res []*model.User) {
	// TODO: joint account
	rows, err := this.db.Query("SELECT id, user_id, nickname, url, sex, age, hobby FROM users")
	if err != nil {
		return nil
	}
	for rows.Next() {
		u := &model.User{}
		if err := rows.Scan(&u.ID, &u.UserId, &u.NickName, &u.Url, &u.Sex, &u.Age, &u.Hobby); err != nil {
			res = append(res, u)
		}
	}
	return res
}

func (this *User) GetUserById(id uint64) *model.User {
	row := this.db.QueryRow("SELECT id, user_id, nickname, url, sex, age, hobby FROM users WHERE id = ?", id)
	u := &model.User{}
	if err := row.Scan(&u.ID, &u.UserId, &u.NickName, &u.Url, &u.Sex, &u.Age, &u.Hobby); err != nil {
		return nil
	}
	return u
}

func (this *User) GetAccountByUserId(uid string) *model.User {
	row := this.db.QueryRow("SELECT " +
		"u.id, u.user_id, u.nickname, u.url, u.sex, u.age, u.hobby " +
		"a.email, a.phone, a.activated, a.locked " +
		"FROM users AS u " +
		"LEFT JOIN accounts AS a " +
		"ON a.user_id = b.user_id " +
		"WHERE u.user_id = ?", uid)
	a := &model.User{}
	err := row.Scan(&a.ID, &a.UserId, &a.NickName, &a.Url, &a.Sex, &a.Age, &a.Hobby,
		&a.Account.Email, &a.Account.Phone, &a.Account.Activated, &a.Account.Locked)
	if err != nil {
		return nil
	}
	return a
}

func (this *User) GetAccountByEmail(email string) *model.User {
	row := this.db.QueryRow("SELECT " +
		"u.id, u.user_id, u.nickname, u.url, u.sex, u.age, u.hobby " +
		"a.email, a.phone, a.activated, a.locked " +
		"FROM users AS u " +
		"LEFT JOIN accounts AS a " +
		"ON a.user_id = b.user_id " +
		"WHERE u.user_id = ?", email)
	a := &model.User{}
	err := row.Scan(&a.ID, &a.UserId, &a.NickName, &a.Url, &a.Sex, &a.Age, &a.Hobby,
		&a.Account.Email, &a.Account.Phone, &a.Account.Activated, &a.Account.Locked)
	if err != nil {
		return nil
	}
	return a
}

func (this *User) GetAccountByPhone(phone string) *model.User {
	row := this.db.QueryRow("SELECT " +
		"u.id, u.user_id, u.nickname, u.url, u.sex, u.age, u.hobby " +
		"a.email, a.phone, a.activated, a.locked " +
		"FROM users AS u " +
		"LEFT JOIN accounts AS a " +
		"ON a.user_id = b.user_id " +
		"WHERE u.user_id = ?", phone)
	a := &model.User{}
	err := row.Scan(&a.ID, &a.UserId, &a.NickName, &a.Url, &a.Sex, &a.Age, &a.Hobby,
		&a.Account.Email, &a.Account.Phone, &a.Account.Activated, &a.Account.Locked)
	if err != nil {
		return nil
	}
	return a
}

func (this *User) Create(u *model.User) error {
	stmt, err := this.db.Prepare("INSERT users(user_id, nickname, url, sex, age, hobby) VALUES(?,?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	res, err := stmt.Exec(u.UserId, u.NickName, u.Url, u.Sex, u.Age, u.Hobby)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = uint64(id)
	// TODO: Create Account
	return nil
}

func (this *User) Update(u *model.User) error {
	stmt, err := this.db.Prepare("UPDATE users SET user_id = ?, nickname = ?, url = ?, sex = ?, age = ?, hobby = ? WHERE id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(u.UserId, u.NickName, u.Url, u.Sex, u.Age, u.Hobby, u.ID); err != nil {
		return err
	}
	return nil
}

func (this *User) Delete(u *model.User) error {
	stmt, err := this.db.Prepare("UPDATE FROM users WHERE id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(u.ID); err != nil {
		return err
	}
	return nil
}
