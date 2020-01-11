package model

import (
	"database/sql"
	"errors"
	"fmt"
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/utils"
)

type DB struct {
	db *sql.DB
}

func NewDB(c config.MySQL) *DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", c.User, c.Password, c.Host, c.DBName))
	if err != nil {
		panic(err.Error())
	}
	return &DB{
		db: db,
	}
}

func (this *DB) Create(account *Account) error {
	pre, err := this.db.Prepare("INSERT accounts(user_id, email, phone, password, activated, locked) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	res , err := pre.Exec(account.UserId, account.Email, account.Phone, account.Password, account.Activated, account.Locked)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	account.ID = uint(id)
	return nil
}

func (this *DB) VerifyPassword(account *Account) error {
	var rows *sql.Rows
	var err error
	if account.ID != 0 {
		rows, err = this.db.Query("SELECT * FROM accounts WHERE id=?", account.ID)
	} else if account.Phone != "" {
		rows, err = this.db.Query("SELECT * FROM accounts WHERE phone=?", account.Phone)
	} else if account.Email != "" {
		rows, err = this.db.Query("SELECT * FROM accounts WHERE email=?", account.Email)
	} else {
		rows, err = this.db.Query("SELECT * FROM accounts WHERE user_id=?", account.UserId)
	}
	if err != nil {
		return errors.New(err.Error())
	}
	old := account.Password
	err = rows.Scan(&account.ID, &account.UserId, &account.Email, &account.Phone, &account.Password, &account.Activated, &account.Locked)
	if err != nil {
		return errors.New(err.Error())
	}
	if account.Password != utils.GenerateMD5String(old) {
		return errors.New("incorrect account or password")
	}
	return nil
}

func (this *DB) IsAccountExist(account Account) bool {
	var rows *sql.Rows
	var err error
	if account.ID != 0 {
		rows, err = this.db.Query("SELECT * FROM accounts WHERE id=?", account.ID)
	} else if account.Phone != "" {
		rows, err = this.db.Query("SELECT * FROM accounts WHERE phone=?", account.Phone)
	} else if account.Email != "" {
		rows, err = this.db.Query("SELECT * FROM accounts WHERE email=?", account.Email)
	} else {
		rows, err = this.db.Query("SELECT * FROM accounts WHERE user_id=?", account.UserId)
	}
	if err != nil {
		return false
	}
	return rows.Next()
}

func (this *DB) UpdatePassword(account Account) error {
	var sql, value interface{}
	if account.ID != 0 {
		sql, value = "UPDATE accounts SET password=? WHERE id=?", account.ID
	} else if account.Phone != "" {
		sql, value = "UPDATE accounts SET password=? WHERE phone=?", account.Phone
	} else if account.Email != "" {
		sql, value = "UPDATE accounts SET password=? WHERE email=?", account.Email
	} else {
		sql, value = "UPDATE accounts SET password=? WHERE user_id=?", account.UserId
	}
	pre, err := this.db.Prepare(sql.(string))
	if err != nil {
		return err
	}
	res, err := pre.Exec(utils.GenerateMD5String(account.Password), value)
	if err != nil {
		return err
	}
	num, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if num == 0 {
		return errors.New("no such data")
	}
	return nil
}

func (this *DB) UpdateActivated(account Account) error {
	var sql, value interface{}
	if account.ID != 0 {
		sql, value = "UPDATE accounts SET activated=? WHERE id=?", account.ID
	} else if account.Phone != "" {
		sql, value = "UPDATE accounts SET activated=? WHERE phone=?", account.Phone
	} else if account.Email != "" {
		sql, value = "UPDATE accounts SET activated=? WHERE email=?", account.Email
	} else {
		sql, value = "UPDATE accounts SET activated=? WHERE user_id=?", account.UserId
	}
	pre, err := this.db.Prepare(sql.(string))
	if err != nil {
		return err
	}
	res, err := pre.Exec(account.Activated, value)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (this *DB) UpdateLocked(account Account) error {
	var sql, value interface{}
	if account.ID != 0 {
		sql, value = "UPDATE accounts SET locked=? WHERE id=?", account.ID
	} else if account.Phone != "" {
		sql, value = "UPDATE accounts SET locked=? WHERE phone=?", account.Phone
	} else if account.Email != "" {
		sql, value = "UPDATE accounts SET locked=? WHERE email=?", account.Email
	} else {
		sql, value = "UPDATE accounts SET locked=? WHERE user_id=?", account.UserId
	}
	pre, err := this.db.Prepare(sql.(string))
	if err != nil {
		return err
	}
	res, err := pre.Exec(account.Locked, value)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (this *DB) UpdateEmail(account Account) error {
	pre, err := this.db.Prepare("UPDATE accounts SET email=? WHERE user_id=?")
	if err != nil {
		return err
	}
	res, err := pre.Exec(account.Email, account.UserId)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (this *DB) UpdatePhone(account Account) error {
	pre, err := this.db.Prepare("UPDATE accounts SET phone=? WHERE user_id=?")
	if err != nil {
		return err
	}
	res, err := pre.Exec(account.Phone, account.UserId)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

