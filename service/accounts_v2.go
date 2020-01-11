package service

import (
	"errors"
	"gitlab.com/pangold/auth/model"
	"gitlab.com/pangold/auth/utils"
)

// Generate verification code, store it and send it out
func (this *Account) RequestVCode(a model.Account) error {
	// relative account(only email or phone)
	target := a.Email
	if a.Phone != "" {
		target = a.Phone
	}
	code := utils.GenerateRandomNumber(4)
	// expire in 5 minutes
	if err := this.cache.SetCacheValue("auth", target, code, 60 * 5); err != nil {
		return errors.New("server error with " + err.Error())
	}
	if err := this.vcode.SendVerificationCode(target, code); err != nil {
		return errors.New("server error with " + err.Error())
	}
	return nil
}

// compare to the stored verification code
func (this *Account) CheckVCode(a model.Account) bool {
	target := a.Email
	if a.Phone != "" {
		target = a.Phone
	}
	code, err := this.cache.GetCacheValue("auth", target, string(""))
	if err != nil {
		return false
	}
	if code.(string) != a.VCode {
		return false
	}
	this.cache.ResetCacheKey("auth", target)
	return true
}

func (this *Account) RegisterWithVCode(a model.Account) error {
	if err := a.IsValid(); err != nil {
		return err
	}
	if a.VCode == "" || !this.CheckVCode(a) {
		return errors.New("invalid verification code")
	}
	if err := this.db.Create(&a); err != nil {
		return err
	}
	return nil
}

func (this *Account) LoginWithVCode(a model.Account) (string, error) {
	if a.VCode == "" || !this.CheckVCode(a) {
		return "", errors.New("invalid verification code")
	}
	return this.Login(a)
}

func (this *Account) ResetByVCode(a model.Account) error {
	if a.VCode == "" || !this.CheckVCode(a) {
		return errors.New("invalid verification code")
	}
	if err := this.db.UpdatePassword(a); err != nil {
		return err
	}
	return nil
}

func (this *Account) Lock(a model.Account, lock bool) error {
	if err := this.db.UpdateActivated(a); err != nil {
		return err
	}
	return nil
}

func (this *Account) BindEmail(a model.Account, lock bool) error {
	if a.Email == "" || a.UserId == "" {
		return errors.New("invalid params")
	}
	if err := this.db.UpdateEmail(a); err != nil {
		return err
	}
	return nil
}

func (this *Account) BindPhone(a model.Account, lock bool) error {
	if a.Phone == "" || a.UserId == "" {
		return errors.New("invalid params")
	}
	if err := this.db.UpdatePhone(a); err != nil {
		return err
	}
	return nil
}

