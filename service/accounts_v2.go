package service

import (
	"errors"
	"gitlab.com/pangold/auth/model"
)

// Generate verification code, store it and send it out
func (this *AccountService) RequestVCode(a model.Account) error {
	// generate v code
	// store v code
	// send
	return nil
}

// compare to the stored verification code
func (this *AccountService) CheckVCode(a model.Account) bool {
	// get v code
	// check
	return true
}

func (this *AccountService) RegisterWithVCode(a model.Account) error {
	if a.VCode == "" {
		return errors.New("empty verification code")
	}
	if a.UserId == "" && a.Email == "" && a.Phone == "" {
		return errors.New("empty account")
	}
	if a.Password == "" {
		return errors.New("empty password")
	}
	if !this.CheckVCode(a) {
		return errors.New("invalid verification code")
	}
	if err := this.db.Create(&a); err != nil {
		return err
	}
	return nil
}

func (this *AccountService) LoginWithVCode(a model.Account) (string, error) {
	if !this.CheckVCode(a) {
		return "", errors.New("invalid verification code")
	}
	return this.Login(a)
}

func (this *AccountService) ResetByVCode(a model.Account) error {
	if !this.CheckVCode(a) {
		return errors.New("invalid verification code")
	}
	if err := this.db.UpdatePassword(a); err != nil {
		return err
	}
	return nil
}

func (this *AccountService) Lock(a model.Account, lock bool) error {
	if err := this.db.UpdateActivated(a); err != nil {
		return err
	}
	return nil
}

func (this *AccountService) BindEmail(a model.Account, lock bool) error {
	if a.Email == "" || a.UserId == "" {
		return errors.New("invalid params")
	}
	if err := this.db.UpdateEmail(a); err != nil {
		return err
	}
	return nil
}

func (this *AccountService) BindPhone(a model.Account, lock bool) error {
	if a.Phone == "" || a.UserId == "" {
		return errors.New("invalid params")
	}
	if err := this.db.UpdatePhone(a); err != nil {
		return err
	}
	return nil
}

