package db

import (
	"database/sql"
	"gitlab.com/pangold/auth/model"
)

type User struct {
	db *sql.DB
}

func NewUser(conn *sql.DB) *User {
	return &User{
		db: conn,
	}
}

func (this *User) GetUsers() (res []*model.User) {
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
		"a.id, a.user_id, a.email, a.phone, a.activated, a.locked " +
		"FROM users AS u " +
		"RIGHT JOIN accounts AS a " +
		"ON a.user_id = u.user_id " +
		"WHERE u.user_id = ?", uid)
	a := &model.User{}
	err := row.Scan(&a.ID, &a.UserId, &a.NickName, &a.Url, &a.Sex, &a.Age, &a.Hobby,
		&a.ID, &a.UserId, &a.Account.Email, &a.Account.Phone, &a.Account.Activated, &a.Account.Locked)
	if err != nil {
		return nil
	}
	return a
}

func (this *User) GetAccountByEmail(email string) *model.User {
	row := this.db.QueryRow("SELECT " +
		"u.id, u.user_id, u.nickname, u.url, u.sex, u.age, u.hobby " +
		"a.id, a.user_id, a.email, a.phone, a.activated, a.locked " +
		"FROM users AS u " +
		"RIGHT JOIN accounts AS a " +
		"ON a.user_id = u.user_id " +
		"WHERE u.email = ?", email)
	a := &model.User{}
	err := row.Scan(&a.ID, &a.UserId, &a.NickName, &a.Url, &a.Sex, &a.Age, &a.Hobby,
		&a.ID, &a.UserId, &a.Account.Email, &a.Account.Phone, &a.Account.Activated, &a.Account.Locked)
	if err != nil {
		return nil
	}
	return a
}

func (this *User) GetAccountByPhone(phone string) *model.User {
	row := this.db.QueryRow("SELECT " +
		"u.id, u.user_id, u.nickname, u.url, u.sex, u.age, u.hobby " +
		"a.id, a.user_id, a.email, a.phone, a.activated, a.locked " +
		"FROM users AS u " +
		"RIGHT JOIN accounts AS a " +
		"ON a.user_id = u.user_id " +
		"WHERE u.phone = ?", phone)
	a := &model.User{}
	err := row.Scan(&a.ID, &a.UserId, &a.NickName, &a.Url, &a.Sex, &a.Age, &a.Hobby,
		&a.ID, &a.UserId, &a.Account.Email, &a.Account.Phone, &a.Account.Activated, &a.Account.Locked)
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
	// a := model.Account{UserId: u.UserId}
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
	stmt, err := this.db.Prepare("DELETE FROM users WHERE id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(u.ID); err != nil {
		return err
	}
	return nil
}
