package service

import (
	"errors"
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/middleware"
	"gitlab.com/pangold/auth/model"
	"strconv"
)

type AccountService struct {
	db *model.DB
	token middleware.Token
}

func NewAccountService(c config.MySqlConfig) *AccountService {
	return &AccountService{
		db: model.NewDB(c),
	}
}

func (this *AccountService) Register(a model.Account) error {
	if a.UserId == "" && a.Email == "" && a.Phone == "" {
		return errors.New("empty account")
	}
	if a.Password == "" {
		return errors.New("empty password")
	}
	if err := this.db.Create(&a); err != nil {
		return err
	}
	return errors.New("invalid params")
}

func (this *AccountService) GetActivationUrl(a model.Account) error {
	// send activation-url via email
	return nil
}

func (this *AccountService) Activate(email, code string) error {
	a := model.Account{Activated: true, Email: email}
	if err := this.db.UpdateActivated(a); err != nil {
		return err
	}
	return nil
}

func (this *AccountService) Login(a model.Account) (string, error) {
	if err := this.db.VerifyPassword(&a); err != nil {
		return "", errors.New("account or password is invalid")
	}
	token, err := this.token.GenerateToken(strconv.Itoa(int(a.ID)), a.UserId, "pc")
	if err != nil {
		return "", errors.New("generate token error " + err.Error())
	}
	return token, nil
}

func (this *AccountService) Logout(token string) error {
	// reset token
	return this.token.ResetToken(token)
}

func (this *AccountService) Forgot(a model.Account) error {
	// generate hash code with email
	// send email
	return nil
}

func (this *AccountService) ResetByHashCode(a model.Account) error {
	// get email by hash code
	a.Email = ""
	//
	if err := this.db.UpdatePassword(a); err != nil {
		return err
	}
	return nil
}

func (this *AccountService) IsUserIdExist(a model.Account) bool {
	if a.UserId == "" {
		return false
	}
	return this.db.IsAccountExist(a)
}

func (this *AccountService) IsEmailExist(a model.Account) bool {
	if a.Email == "" {
		return false
	}
	return this.db.IsAccountExist(a)
}

func (this *AccountService) IsPhoneExist(a model.Account) bool {
	if a.Phone == "" {
		return false
	}
	return this.db.IsAccountExist(a)
}