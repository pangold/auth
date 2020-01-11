package db

import (
	"database/sql"
	"errors"
	"gitlab.com/pangold/auth/model"
	"gitlab.com/pangold/auth/utils"
)

type Auth struct {
	db *sql.DB
}

func NewAuth(conn *sql.DB) *Auth {
	return &Auth{
		db: conn,
	}
}

func (this *Auth) Create(account *model.Account) error {
	if err := account.IsValid(); err != nil {
		return err
	}
	stmt, err := this.db.Prepare("INSERT accounts(user_id, email, phone, password, activated, locked) VALUES(?,?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	pwd := utils.GenerateMD5String(account.Password)
	res, err := stmt.Exec(account.UserId, account.Email, account.Phone, pwd, account.Activated, account.Locked)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	account.ID = uint64(id)
	return nil
}

func (this *Auth) condition(a model.Account) (cond string, val interface{}) {
	if a.ID != 0 {
		cond, val = " WHERE id = ?", a.ID
	} else if a.Phone != "" {
		cond, val = " WHERE phone = ?", a.Phone
	} else if a.Email != "" {
		cond, val = " WHERE email = ?", a.Email
	} else if a.UserId != "" {
		cond, val = " WHERE user_id = ?", a.UserId
	}
	return cond, val
}

func (this *Auth) VerifyPassword(a *model.Account) error {
	if err := a.IsAccountValid(); err != nil {
		return err
	}
	cond, val := this.condition(*a)
	row := this.db.QueryRow("SELECT id, user_id, password, activated, locked FROM accounts" + cond, val)
	old := a.Password
	if err := row.Scan(&a.ID, &a.UserId, &a.Password, &a.Activated, &a.Locked); err != nil {
		return err
	}
	if err := a.IsValidState(); err != nil {
		return err
	}
	if a.Password != utils.GenerateMD5String(old) {
		return errors.New("incorrect account or password")
	}
	return nil
}

func (this *Auth) IsAccountExist(a *model.Account) bool {
	if err := a.IsAccountValid(); err != nil {
		return false
	}
	cond, val := this.condition(*a)
	row := this.db.QueryRow("SELECT id FROM accounts" + cond, val)
	if err := row.Scan(&a.ID); err != nil {
		return false
	}
	return true
}

func (this *Auth) UpdatePassword(a model.Account) error {
	if err := a.IsAccountValid(); err != nil {
		return err
	}
	if err := a.IsPasswordValid(); err != nil {
		return err
	}
	cond, val := this.condition(a)
	stmt, err := this.db.Prepare("UPDATE accounts SET password = ? " + cond)
	defer stmt.Close()
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(utils.GenerateMD5String(a.Password), val); err != nil {
		return err
	}
	return nil
}

func (this *Auth) UpdateActivated(a model.Account) error {
	if err := a.IsAccountValid(); err != nil {
		return err
	}
	cond, val := this.condition(a)
	stmt, err := this.db.Prepare("UPDATE accounts SET activated = ? " + cond)
	defer stmt.Close()
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(a.Activated, val); err != nil {
		return err
	}
	return nil
}

func (this *Auth) UpdateLocked(a model.Account) error {
	if err := a.IsAccountValid(); err != nil {
		return err
	}
	cond, val := this.condition(a)
	stmt, err := this.db.Prepare("UPDATE accounts SET locked = ? " + cond)
	defer stmt.Close()
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(a.Locked, val); err != nil {
		return err
	}
	return nil
}

func (this *Auth) UpdateUserId(a model.Account) error {
	sql := "UPDATE accounts SET user_id = ? WHERE id = ?"
	stmt, err := this.db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(a.UserId, a.ID); err != nil {
		return err
	}
	return nil
}

func (this *Auth) UpdateEmail(a model.Account) error {
	sql := "UPDATE accounts SET email = ? WHERE id = ?"
	stmt, err := this.db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(a.Email, a.ID); err != nil {
		return err
	}
	return nil
}

func (this *Auth) UpdatePhone(a model.Account) error {
	sql := "UPDATE accounts SET phone = ? WHERE id = ?"
	stmt, err := this.db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(a.Phone, a.ID); err != nil {
		return err
	}
	return nil
}

