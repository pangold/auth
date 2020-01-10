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

// Hooks
func (a *Account) BeforeSave() error {
	return a.IsValid()
}

func (a *Account) BeforeCreate() error {
	return a.IsValid()
}

func (a *Account) BeforeUpdate() error {
	return a.IsValid()
}

func (a *Account) BeforeDelete() error {
	return nil
}

// Column Verification
func (a *Account) IsUsernameValid() error {
	if len(a.UserId) < 4 {
		return errors.New("user id is too short.")
	}
	return nil
}

func (a *Account) IsEmailValid() error {
	// the first way to use regexp
	if a.Email == "" {
		// Email could be empty, so return directly
		return nil
	}
	if matched, _ := regexp.MatchString("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+", a.Email); !matched {
		return errors.New("invalid email address")
	}
	return nil
}

func (a *Account) IsPhoneValid() error {
	// the second way to use regexp(similar)
	if a.Phone == "" {
		return nil
	}
	regular := "^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}$"
	reg := regexp.MustCompile(regular)
	if matched := reg.MatchString(a.Phone); !matched {
		return errors.New("invalid phone number")
	}
	return nil
}

func (a *Account) IsPasswordValid() error {
	// This is special, it needs to encrypt...
	// Now, this implementation is not correct, 
	// But it doesn't matter. It will always return nil
	if len(a.Password) < 6 {
		return errors.New("password is too short.")
	}
	return nil
}

func (a *Account) IsValid() error {
	if err := a.IsUsernameValid(); err != nil {
		return err
	}
	if err := a.IsEmailValid(); err != nil {
		return err
	}
	if err := a.IsPhoneValid(); err != nil {
		return err
	}
	if err := a.IsPasswordValid(); err != nil {
		return err
	}
	//? 
	//if len(a.Password) > 0 {
	//	a.Password = encryptText(a.Password)
	//}
	return nil
}

func (a *Account) IsStateValid() error {
	if !a.Activated {
		return errors.New("invalid account: in activated state")
	}
	if a.Locked {
		return errors.New("invalid account: in locked state")
	}
	return nil
}