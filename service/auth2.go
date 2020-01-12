package service

import (
	"errors"
	"gitlab.com/pangold/auth/model"
	"gitlab.com/pangold/auth/utils"
)

// Generate verification code, store it and send it out
func (this *Auth) RequestVCode(a model.Account) error {
	target := a.Email
	if a.Phone != "" {
		target = a.Phone
	}
	if target == "" {
		return errors.New("invalid account")
	}
	code := utils.GenerateRandomNumber(4)
	// expire in 5 minutes
	if err := this.cache.SetCacheValue("auth", target, code, this.config.VCodeExpire); err != nil {
		return errors.New("server error with " + err.Error())
	}
	if err := this.vcode.SendVerificationCode(target, code); err != nil {
		return errors.New("server error with " + err.Error())
	}
	return nil
}

// compare to the stored verification code
func (this *Auth) CheckVCode(a model.Account) error {
	target := a.Email
	if a.Phone != "" {
		target = a.Phone
	}
	if target == "" {
		return errors.New("invalid account")
	}
	if a.Code == "" {
		return errors.New("invalid verification code")
	}
	code, err := this.cache.GetCacheValue("auth", target, string(""))
	if err != nil {
		return err
	}
	if code.(string) != a.Code {
		return errors.New("incorrect verification code")
	}
	this.cache.ResetCacheKey("auth", target)
	return nil
}

func (this *Auth) RegisterWithVCode(a model.Account) error {
	if err := a.IsValid(); err != nil {
		return err
	}
	if err := this.CheckVCode(a); err != nil {
		return err
	}
	a.Activated = true
	if err := this.db.Create(&a); err != nil {
		return err
	}
	return nil
}

func (this *Auth) LoginWithVCode(a model.Account) (string, error) {
	if err := this.CheckVCode(a); err != nil {
		return "", err
	}
	return this.Login(a)
}

func (this *Auth) ResetByVCode(a model.Account) error {
	if err := this.CheckVCode(a); err != nil {
		return err
	}
	if err := this.db.UpdatePassword(a); err != nil {
		return err
	}
	return nil
}

func (this *Auth) Lock(a model.Account, lock bool) error {
	a.Locked = lock
	if err := this.db.UpdateActivated(a); err != nil {
		return err
	}
	return nil
}

func (this *Auth) BindEmail(a model.Account, bind bool) error {
	if a.Email == "" || a.ID == 0 {
		return errors.New("invalid params")
	}
	if !bind {
		a.Email = ""
	}
	if err := a.IsAccountValid(); err != nil {
		return errors.New("unbind limited")
	}
	if err := this.db.UpdateEmail(a); err != nil {
		return err
	}
	return nil
}

func (this *Auth) BindPhone(a model.Account, bind bool) error {
	if a.Phone == "" || a.ID == 0 {
		return errors.New("invalid params")
	}
	if !bind {
		a.Phone = ""
	}
	if err := a.IsAccountValid(); err != nil {
		return errors.New("unbind limited")
	}
	if err := this.db.UpdatePhone(a); err != nil {
		return err
	}
	return nil
}

