package service

import (
	"errors"
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/middleware"
	"gitlab.com/pangold/auth/model"
	"strconv"
)

type AccountService struct {
	config config.Server
	db *model.DB
	token middleware.Token
	cache middleware.Cache
}

func NewAccountService(conf config.Config, t middleware.Token, c middleware.Cache) *AccountService {
	return &AccountService{
		config: conf.Server,
		db: model.NewDB(conf.MySQL),
		token: t,
		cache: c,
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
	// FIXME: set cache after registering if it's in activated state
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
	// FIXME: get cache before checking password
	if err := this.db.VerifyPassword(&a); err != nil {
		return "", errors.New("invalid account or password")
	}
	token, err := this.token.GenerateToken(strconv.Itoa(int(a.ID)), a.UserId, "pc", this.config.TokenExpire)
	if err != nil {
		return "", errors.New("generate token error " + err.Error())
	}
	return token, nil
}

func (this *AccountService) Logout(token string) error {
	return this.token.ResetToken(token)
}

func (this *AccountService) Forgot(a model.Account) error {
	// check email
	// generate hash code with email
	// send email
	return nil
}

func (this *AccountService) ResetByHashCode(a model.Account) error {
	// get email by hash code
	a.Email = ""
	// FIXME: clear cache before updating
	if err := this.db.UpdatePassword(a); err != nil {
		return err
	}
	return nil
}

func (this *AccountService) IsUserIdExist(a model.Account) bool {
	if a.UserId == "" {
		return false
	}
	// FIXME: check if exist from cache before reading db
	return this.db.IsAccountExist(a)
}

func (this *AccountService) IsEmailExist(a model.Account) bool {
	if a.Email == "" {
		return false
	}
	// FIXME: check if exist from cache before reading db
	return this.db.IsAccountExist(a)
}

func (this *AccountService) IsPhoneExist(a model.Account) bool {
	if a.Phone == "" {
		return false
	}
	// FIXME: check if exist from cache before reading db
	return this.db.IsAccountExist(a)
}