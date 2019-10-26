package service

import (
	"fmt"
	"testing"

	"../utils"
	"../model"
)

func init() {
	if err := model.DropTable("accounts"); err != nil {
		fmt.Printf("drop table error: %v\n", err)
	}
	model.MigrateAccount()
}

func TestRegister(t *testing.T) {
	code, err := Register("pangold@163.com", "88888888", "https://sample.com/account/activate")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	email, err := utils.GetCacheValue("account", code, "")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if email.(string) != "pangold@163.com" {
		t.Errorf("unmatch email with code")
	}
	if err := ActivateAccountWithHashCode(code); err != nil {
		t.Errorf(err.Error())
	}
}

func TestRegisterWithEmailAndCode(t *testing.T) {
	email := "panking@126.com"
	if err := RequireVerificationCode(email, ""); err != nil {
		t.Errorf(err.Error())
	}
	code, err := utils.GetCacheValue("account", email, "")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if err := RegisterWithCode(email, "", "88888888", code.(string)); err != nil {
		t.Errorf(err.Error())
	}
}

func TestRegisterWithPhoneAndCode(t *testing.T) {
	phone := "13800000000"
	if err := RequireVerificationCode("", phone); err != nil {
		t.Errorf(err.Error())
	}
	code, err := utils.GetCacheValue("account", phone, "")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if err := RegisterWithCode("", phone, "88888888", code.(string)); err != nil {
		t.Errorf(err.Error())
	}
}

func TestRegisterWithUsername(t *testing.T) {
	if err := RegisterWithUsername("dora", "dora@gmail.com", "18866888866", "88888888"); err != nil {
		t.Error(err.Error())
	}
}

func TestLogin(t *testing.T) {
	token, err := Login("dora", "", "", "88888888")
	if err != nil {
		t.Errorf(err.Error())
	}
	if err = Logout(token); err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println("token: " + token)
	//
	if token, err = Login("", "", "13800000000", "88888888"); err != nil {
		t.Errorf(err.Error())
	}
	if err = Logout(token); err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println("token: " + token)
	//
	if token, err = Login("", "pangold@163.com", "", "88888888"); err != nil {
		t.Errorf(err.Error())
	}
	if err = Logout(token); err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println("token: " + token)
	//
	if token, err = Login("", "panking@126.com", "", "88888888"); err != nil {
		t.Errorf(err.Error())
	}
	if err = Logout(token); err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println("token: " + token)
}

func TestForgotWithEmailCode(t *testing.T) {
	email := "panking@126.com"
	if err := RequireVerificationCode(email, ""); err != nil {
		t.Errorf(err.Error())
	}
	code, err := utils.GetCacheValue("account", email, "")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if err := ForgotWithCode(email, "", "888888", code.(string)); err != nil {
		t.Errorf(err.Error())
	}
}

func TestForgotWithPhoneAndCode(t *testing.T) {
	phone := "13800000000"
	if err := RequireVerificationCode("", phone); err != nil {
		t.Errorf(err.Error())
	}
	code, err := utils.GetCacheValue("account", phone, "")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if err := ForgotWithCode("", phone, "888888", code.(string)); err != nil {
		t.Errorf(err.Error())
	}
}

func TestForgot(t *testing.T) {
	code, err := Forgot("pangold@163.com", "https://sample.com/account/activate")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	email, err := utils.GetCacheValue("account", code, "")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if email.(string) != "pangold@163.com" {
		t.Errorf("unmatch email with code")
	}
	if err := ResetPasswordByHashCode(code, "888888"); err != nil {
		t.Errorf(err.Error())
	}
}
