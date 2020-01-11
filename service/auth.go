package service

import (
	"errors"
	"gitlab.com/pangold/auth/config"
	"gitlab.com/pangold/auth/middleware"
	"gitlab.com/pangold/auth/model"
	"gitlab.com/pangold/auth/model/db"
	"gitlab.com/pangold/auth/utils"
	"strconv"
)

type Auth struct {
	config config.Server
	db *db.Auth
	token middleware.Token
	cache middleware.Cache
	email middleware.Email
	vcode middleware.VerificationCode
}

func NewAuthService(conf config.Server, db *db.Auth, e middleware.Email, vc middleware.VerificationCode, t middleware.Token, c middleware.Cache) *Auth {
	return &Auth{
		config: conf,
		db: db,
		token: t,
		cache: c,
		email: e,
		vcode: vc,
	}
}

func (this *Auth) Register(a model.Account) error {
	if err := a.IsAccountValid(); err != nil {
		return err
	}
	if err := a.IsPasswordValid(); err != nil {
		return err
	}
	if err := this.db.Create(&a); err != nil {
		return err
	}
	// FIXME: set cache after registering if it's in activated state
	return errors.New("invalid params")
}

func (this *Auth) GetActivationUrl(a model.Account) error {
	if err := a.IsEmailValid(); err != nil {
		return err
	}
	if !this.db.IsAccountExist(&model.Account{Email: a.Email}) {
		return errors.New("email is not exist")
	}
	// generate activation hash code and save
	code := utils.GenerateRandomString(64)
	if err := this.cache.SetCacheValue("auth", code, a.Email, 60 * 5); err != nil {
		return errors.New("server error with " + err.Error())
	}
	// send activation code via email
	if err := this.email.SendActivationEmail(a.Email, code); err != nil {
		return errors.New("email service error with " + err.Error())
	}
	return nil
}

func (this *Auth) Activate(code string) error {
	email, err := this.cache.GetCacheValue("auth", code, string(""))
	if err != nil {
		return errors.New("invalid activation code")
	}
	a := model.Account{Activated: true, Email: email.(string)}
	if err := this.db.UpdateActivated(a); err != nil {
		return err
	}
	return nil
}

func (this *Auth) Login(a model.Account) (string, error) {
	// FIXME: get cache before checking password
	if err := this.db.VerifyPassword(&a); err != nil {
		return "", errors.New("invalid account or password")
	}
	token, err := this.token.GenerateToken(strconv.Itoa(int(a.ID)), a.UserId, "pc", this.config.TokenExpire)
	if err != nil {
		return "", errors.New("server error " + err.Error())
	}
	return token, nil
}

func (this *Auth) Logout(token string) error {
	return this.token.ResetToken(token)
}

func (this *Auth) Forgot(a model.Account) error {
	if err := a.IsEmailValid(); err != nil {
		return err
	}
	if !this.db.IsAccountExist(&model.Account{Email: a.Email}) {
		return errors.New("email is not exist")
	}
	// generate activation hash code and save
	code := utils.GenerateRandomString(64)
	if err := this.cache.SetCacheValue("auth", code, a.Email, 60 * 5); err != nil {
		return errors.New("server error with " + err.Error())
	}
	// send activation code via email
	if err := this.email.SendResetPasswordEmail(a.Email, code); err != nil {
		return errors.New("server error with " + err.Error())
	}
	return nil
}

func (this *Auth) ResetByHashCode(a model.Account) error {
	email, err := this.cache.GetCacheValue("auth", a.Code, string(""))
	if err != nil {
		return errors.New("invalid reset hash code")
	}
	a.Email = email.(string)
	if err := this.db.UpdatePassword(a); err != nil {
		return err
	}
	return nil
}

func (this *Auth) IsUserIdExist(a model.Account) bool {
	if a.UserId == "" {
		return false
	}
	// FIXME: check if exist from cache before reading db
	return this.db.IsAccountExist(&model.Account{UserId: a.UserId})
}

func (this *Auth) IsEmailExist(a model.Account) bool {
	if a.Email == "" {
		return false
	}
	// FIXME: check if exist from cache before reading db
	return this.db.IsAccountExist(&model.Account{Email: a.Email})
}

func (this *Auth) IsPhoneExist(a model.Account) bool {
	if a.Phone == "" {
		return false
	}
	// FIXME: check if exist from cache before reading db

	// wrap again just in case input is incorrect
	return this.db.IsAccountExist(&model.Account{Phone: a.Phone})
}