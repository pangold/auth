package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"regexp"
)

type Account struct {
	gorm.Model
	// You can use any one of UserId, Email or PhoneNumber to login
	// the UserId here is the same to the UserId of table User
	UserId  	string `json:"username"           gorm:"type:varchar(100);unique_index"`
	Email		string `json:"email"              gorm:"type:varchar(100)"`
	Phone		string `json:"phone"              gorm:"type:varchar(100)"`
	Password	string `json:"password"           gorm:""`
	Activated	bool   `json:"activated"          gorm:"default:false"`	// Needs to activate when registerred
	Locked      bool   `json:"locked"             gorm:"default:false"`	// Input wrong userId & password 5 times
	VCode       string `json:"vcode"              gorm:""`
	// LastLogInAt	time.Time
	// LastLogOutAt	time.Time
	// LastLoggedIp	string
}

func (this *Account) IsAccountValid() error {
	if this.UserId == "" && this.Email == "" && this.Phone == "" {
		return errors.New("invalid account")
	}
	return nil
}

func (this *Account) IsUserIdValid() error {
	// user id could be empty
	if this.UserId == "" {
		return nil
	}
	if len(this.UserId) < 6 {
		return errors.New("user id is too short")
	}
	return nil
}

func (this *Account) IsEmailValid() error {
	// email could be empty
	if this.Email == "" {
		return nil
	}
	if matched, _ := regexp.MatchString("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+", this.Email); !matched {
		return errors.New("invalid email")
	}
	return nil
}

func (this *Account) IsPhoneValid() error {
	// phone could be empty
	if this.Phone == "" {
		return nil
	}
	reg := regexp.MustCompile("^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}$")
	if matched := reg.MatchString(this.Phone); !matched {
		return errors.New("invalid phone number")
	}
	return nil
}

func (this *Account) IsPasswordValid() error {
	// TODO: too simple
	if len(this.Password) < 6 {
		return errors.New("password is too short")
	}
	return nil
}

func (this *Account) IsValid() error {
	// must be one of uid, email, phone is not empty
	if err := this.IsAccountValid(); err != nil {
		return err
	}
	if err := this.IsUserIdValid(); err != nil {
		return err
	}
	if err := this.IsEmailValid(); err != nil {
		return err
	}
	if err := this.IsPhoneValid(); err != nil {
		return err
	}
	if err := this.IsPasswordValid(); err != nil {
		return err
	}
	return nil
}

func (this *Account) IsValidState() error {
	if !this.Activated {
		return errors.New("inactivated account")
	}
	if this.Locked {
		return errors.New("locked account")
	}
	return nil
}