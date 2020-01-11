package db

import (
	"database/sql"
	"gitlab.com/pangold/auth/model"
	"gitlab.com/pangold/auth/utils"
)

type Account struct {
	db *sql.DB
}

func NewAccount(conn *sql.DB) *Account {
	return &Account{
		db: conn,
	}
}

func (this *Account) GetAccounts() (res []*model.Account) {
	rows, err := this.db.Query("SELECT id, user_id, email, phone, activated, locked FROM accounts")
	if err != nil {
		return nil
	}
	for rows.Next() {
		a := &model.Account{}
		if err := rows.Scan(&a.ID, &a.UserId, &a.Email, &a.Phone, &a.Activated, &a.Locked); err != nil {
			res = append(res, a)
		}
	}
	return res
}

func (this *Account) GetActivatedAccounts(activated bool) (res []*model.Account) {
	rows, err := this.db.Query("SELECT id, user_id, email, phone, activated, locked FROM accounts WHERE activated = ?", activated)
	if err != nil {
		return nil
	}
	for rows.Next() {
		a := &model.Account{}
		if err := rows.Scan(&a.ID, &a.UserId, &a.Email, &a.Phone, &a.Activated, &a.Locked); err != nil {
			res = append(res, a)
		}
	}
	return res
}

func (this *Account) GetLockedAccounts(locked bool) (res []*model.Account) {
	rows, err := this.db.Query("SELECT id, user_id, email, phone, activated, locked FROM accounts WHERE locked = ?", locked)
	if err != nil {
		return nil
	}
	for rows.Next() {
		a := &model.Account{}
		if err := rows.Scan(&a.ID, &a.UserId, &a.Email, &a.Phone, &a.Activated, &a.Locked); err != nil {
			res = append(res, a)
		}
	}
	return res
}

func (this *Account) GetAccountById(id uint64) *model.Account {
	row := this.db.QueryRow("SELECT id, user_id, email, phone, activated, locked FROM accounts WHERE id = ?", id)
	a := &model.Account{}
	if err := row.Scan(&a.ID, &a.UserId, &a.Email, &a.Phone, &a.Activated, &a.Locked); err != nil {
		return nil
	}
	return a
}

func (this *Account) GetAccountByUserId(uid string) *model.Account {
	row := this.db.QueryRow("SELECT id, user_id, email, phone, activated, locked FROM accounts WHERE user_id = ?", uid)
	a := &model.Account{}
	if err := row.Scan(&a.ID, &a.UserId, &a.Email, &a.Phone, &a.Activated, &a.Locked); err != nil {
		return nil
	}
	return a
}

func (this *Account) GetAccountByEmail(email string) *model.Account {
	row := this.db.QueryRow("SELECT id, user_id, email, phone, activated, locked FROM accounts WHERE email = ?", email)
	a := &model.Account{}
	if err := row.Scan(&a.ID, &a.UserId, &a.Email, &a.Phone, &a.Activated, &a.Locked); err != nil {
		return nil
	}
	return a
}

func (this *Account) GetAccountByPhone(phone string) *model.Account {
	row := this.db.QueryRow("SELECT id, user_id, email, phone, activated, locked FROM accounts WHERE phone = ?", phone)
	a := &model.Account{}
	if err := row.Scan(&a.ID, &a.UserId, &a.Email, &a.Phone, &a.Activated, &a.Locked); err != nil {
		return nil
	}
	return a
}

func (this *Account) Create(a *model.Account) error {
	var pwd string
	if a.Password == "" {
		a.Password = "123456" // default password
	}
	if err := a.IsValid(); err != nil {
		return err
	}
	stmt, err := this.db.Prepare("INSERT accounts(user_id, email, phone, password, activated, locked) VALUES(?,?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	pwd = utils.GenerateMD5String(a.Password)
	res, err := stmt.Exec(a.UserId, a.Email, a.Phone, pwd, a.Activated, a.Locked)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	a.ID = uint64(id)
	return nil
}

func (this *Account) Update(a *model.Account) error {
	if err := a.IsValid(); err != nil {
		return err
	}
	stmt, err := this.db.Prepare("UPDATE accounts SET user_id = ?, email = ?, phone = ?, password = ?, activated = ?, locked = ? WHERE id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	pwd := utils.GenerateMD5String(a.Password)
	if _, err := stmt.Exec(a.UserId, a.Email, a.Phone, pwd, a.Activated, a.Locked, a.ID); err != nil {
		return err
	}
	return nil
}

func (this *Account) Delete(a *model.Account) error {
	stmt, err := this.db.Prepare("UPDATE FROM accounts WHERE id = ?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(a.ID); err != nil {
		return err
	}
	return nil
}