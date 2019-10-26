package service

import (
	"fmt"
	"errors"
	"strconv"
	"../utils"
	"../model"
)


// Require verification code by email
// Registration or Forgot Password will need this function to get verification code
func RequireVerificationCode(email, phone string) error {
	identifier, isEmail := email, true
	if len(identifier) == 0 {
		identifier, isEmail = phone, false
	}
	// Check if this email had required
	if utils.HasCacheKey("account", identifier) {
		return errors.New("requires too frequent.")
	}
	// Generate a verification code, and save into Cache
	code := utils.GenerateRandomNumber(4)
	if err := utils.SetCacheValue("account", identifier, code, 60 * 10); err != nil {
		return errors.New("set cache value error: " + err.Error())
	}
	// Send
	content := utils.GetVerificationText("Corp", code)
	fmt.Println(content)
	if isEmail {
		if err := utils.SendVerificationEmail(email, content); err != nil {
			return errors.New("send email error: " + err.Error())
		}
	} else {
		if err := utils.SendSMS(phone, content); err != nil {
			return errors.New("send sms error: " + err.Error())
		}
	}
	return nil
}

// How Registration Works?
// choise userId/email/phone to register.
// a. choise email without verification code.
//    1) automatically verify email is not registerred.
//    2) input the rest informatin(excluding UserId & Phone)
//    3) press button to submit(POST)
//    4) server side handles this requirement, and response(but now your account is not activated, you will need the url received by your email to activate your account)
//    5) click the URL in your email to activate.
// b. choise email with verification code: you will need to do things as follow:
//    1) automatically verify email is unregisterred.
//    2) press button to get Verification Code by Email
//    3) input your verification code
//    4) input the rest information(excluding userId & phone). 
//    5) press button to submit(POST)
//    6) server side handles this requirement and response success or not.
// c. choise phone with verfication code: 
//    the same process to Email Verification Code Requirement.
// d. simply input your userId, password and password comfirmation, and submit(POST) [Just a special implementation for special system/project]
//    1) input and verify UserId to figure out if it is exist.
//    2) fill up the rest information in this registration form
//    3) submit(POST)
//    4) server side handles this requirement(account will be in activated state throught this way)
//    5) done.

// situation a
func Register(email, password, activateUrl string) (string, error) {
	// FIXME: temp implementation for dealing with password verification
	account := model.Account{ Email: email, Password: password, IsActivated: false, IsEnabled: true, IsLocked: false }
	if err := account.IsPasswordValid(); err != nil {
		return "", errors.New("registration error: " + err.Error())
	}
	// FIXME: replace model to middleware(because of cache)
	if model.IsAccountExist(model.Account{Email: email}) {
		return "", errors.New("registration error: " + email + " is exist")
	}
	// FIXME: 
	if err := model.InsertAccount(&account); err != nil {
		return "", errors.New("registration error: " + err.Error())
	}
	return GenerateActivationCode(email, activateUrl)
}

// situation a (step 2)
func GenerateActivationCode(email, activateUrl string) (string, error) {
	// generate an activation code
	code := utils.GenerateRandomString(64)
	// 3 days
	if err := utils.SetCacheValue("account", code, email, 60 * 60 * 24 * 3); err != nil {
		return "", errors.New("save activation code error: " + err.Error())
	}
	// url: http://example.com/service/action?code=code
	// action in activateUrl supposes to be mapped to ActivateAccountWithHashCode below.
	url := fmt.Sprintf("%s?code=%s", activateUrl, code)
	// FIXME: company
	content := utils.GetActivationText("company", url)
	fmt.Println(content)
	if err := utils.SendActivationEmail(email, content); err != nil {
		return "", errors.New("send activation link(by email) error: " + err.Error())
	}
	return code, nil
}

// after Register(situation a) do
func ActivateAccountWithHashCode(code string) error {
	// see above to see how the code be saved in cache.
	email, err := utils.GetCacheValue("account", code, string(""))
	if err != nil {
		return errors.New("invalid verification code")
	}
	// FIXME:
	if err := model.UpdateActivatedState(model.Account{Email: email.(string), IsActivated: true}); err != nil {
		return err
	}
	utils.ResetCache("account", code)
	return nil
}

// situation b & c
func RegisterWithCode(email, phone, password, code string) error {
	identifier, isEmail := email, true
	if len(identifier) == 0 {
		identifier, isEmail = phone, false
	}
	saved, err := utils.GetCacheValue("account", identifier, string(""))
	if err != nil {
		return errors.New("registration error: invalid verification code.")
	}
	if saved.(string) != code {
		return errors.New("registration error: unmatched verification code.")
	}
	// FIXME: 
	account := model.Account{Email: email}
	if !isEmail {
		account = model.Account{Phone: phone}
	}
	if model.IsAccountExist(account) {
		return errors.New("registration error: " + identifier + " is exist")
	}
	// FIXME: 
	account = model.Account{ Email: email, Phone: phone, Password: password, IsActivated: true, IsEnabled: true, IsLocked: false }
	if err := model.InsertAccount(&account); err != nil {
		return errors.New("registration error: saving info failure")
	}
	utils.ResetCache("account", identifier)
	return nil
}

// situation d
func RegisterWithUsername(username, email, phone, password string) error {
	// FIXME: 
	if len(username) > 0 && model.IsAccountExist(model.Account{Username: username}) {
		return errors.New("registration error: " + username + " is exist")
	}
	// FIXME: 
	if len(email) > 0 && model.IsAccountExist(model.Account{Email: email}) {
		return errors.New("registration error: " + email + " is exist")
	}
	// FIXME: 
	if len(phone) > 0 && model.IsAccountExist(model.Account{Phone: phone}) {
		return errors.New("registration error: " + phone + " is exist")
	}
	// FIXME: 
	account := model.Account{ Username: username, Email: email, Phone: phone, Password: password, IsActivated: true, IsEnabled: true, IsLocked: false }
	if err := model.InsertAccount(&account); err != nil {
		return err
	}
	return nil
}

// How Login Works?
// 1. fill login form with UserId/Email/Phone, and submit(POST)
// 2. server side handle the this requirement.
// 3. response token or error text
func Login(username, email, phone, password string) (string, error) {
	// FIXME: 
	account := model.Account{ Username: username, Email: email, Phone: phone, Password: password }
	res, err := model.VerifyAccountWithPassword(account)
	if err != nil {
		return "", errors.New("login error: username or password is wrong")
	}
	// Further feature: 
	// if unmatched times > 5, account will be locked(implements in the futher)
	token, err2 := utils.GenerateToken(strconv.Itoa(int(res.ID)), res.Username, 60 * 60 * 24 * 30) // 30 days
	if err2 != nil {
		return "", errors.New("login error: unable to generate token")
	}
	return token, nil
}

func Logout(token string) error {
	if err := utils.ResetToken(token); err != nil {
		return errors.New("logout error: " + err.Error())
	}
	return nil
}

func EnabledAccount(username, email, phone string, enabled bool) error {
	// FIXME: 
	account := model.Account{ Username: username, Email: email, Phone: phone, IsLocked: enabled }
	if err := model.UpdateLockedState(account); err != nil {
		return errors.New("enable error: " + err.Error())
	}
	return nil
}

func LockAccount(username, email, phone string, locked bool) error {
	// FIXME: 
	account := model.Account{ Username: username, Email: email, Phone: phone, IsLocked: locked }
	if err := model.UpdateLockedState(account); err != nil {
		return errors.New("unlock error: " + err.Error())
	}
	return nil
}

// How Forgot Password Works?
// a. With Verification Code
//    1. two options(email or phone), select one of them in UI(default is email)
//       a. If email, then press the button "Require Verification Code by Email" -- above RequirePhoneCode
//       b. If phone, then press the button "Require Verification Code by Phone" -- above RequireEmailCode
//    2. check your email or mobile phone message to get your Verification Code.
//    3. input your Verification Code(and the rest information anyway)
//    4. press button "Update" to submit the form(POST).
//    5. server side will handle this POST, and update the password
// b. With Hash Code by Email
//    1. input your email, then click the button to send message to your email
//    2. check your email, there will be a link in the content of that email,   
//       may be here will be something wrong, we don't received the email, then just back to step 1
//    3. click that link(contains a hash code), to show the password update page
//    4. input your new password and submit(POST a hash code and new password)
//    5. server side will handle this POST, and update the password. 
// c. Come To Me

// situation a
func ForgotWithCode(email, phone, password, code string) error {
	var identifier string
	// only one of email and phone is available, then the other one supposes to be empty
	if identifier = email; len(identifier) == 0 {
		identifier = phone
	}
	saved, err := utils.GetCacheValue("account", identifier, string(""))
	if err != nil {
		return errors.New("forgot error: invalid verification code.")
	}
	if saved != code {
		return errors.New("forgot error: unmatched verification code.")
	}
	// FIXME: 
	// update password
	if err := model.UpdatePassword(model.Account{Email: email, Phone: phone, Password: password}); err != nil {
		return errors.New("forgot error: " + err.Error())
	}
	utils.ResetCache("account", identifier)
	return nil
}

// situation b
func Forgot(email, pageUrl string) (string, error) {
	// generate hash code(key: hash_code, value: email)
	code := utils.GenerateRandomString(64)
	if err := utils.SetCacheValue("account", code, email, 60 * 10); err != nil {
		return "", err
	}
	// url: http://example.com/account/action?code=code
	// action in pageUrl supposes to be mapped to ResetPasswordByHashCode below.
	url := fmt.Sprintf("%s?code=%s", pageUrl, code)
	content := utils.GetForgotPasswordText("company", url)
	fmt.Println(content)
	if err := utils.SendForgotEmail(email, content); err != nil {
		return "", err
	}
	return code, nil
}

// after Forgot(situation b)
func ResetPasswordByHashCode(code, password string) error {
	// get email by hashcode(key: hash_code, value: email)
	email, err := utils.GetCacheValue("account", code, "")
	if err != nil {
		return errors.New("reset error: " + err.Error())
	}
	// FIXME: 
	// update password
	if err := model.UpdatePassword(model.Account{Email: email.(string), Password: password}); err != nil {
		return err
	}
	utils.ResetCache("account", code)
	return nil
}

