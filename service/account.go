package service

import (
	"fmt"
	"../utils"
	"../model"
)


// Require verification code by Phone
// Registration or Forgot Password will need this function to get verification code
func RequirePhoneCode(phone string) error {
	// Check if this phone number had required
	if utils.HasCacheKey("account", phone) {
		return "too frequent, please try again later."
	}
	// Generate a verification code, and save into Cache
	code := utils.GenerateRandomNumber(4)
	if err := utils.SetCacheValue("account", phone, code, 60 * 60 * 2); err != nil {
		return err
	}
	// Use SMS to send mobile short message
	content := utils.GetVerificationText("Corp", code)
	if err := utils.SendSMS(phone, content); err != nil {
		return err
	}
	return nil
}

// Require verification code by email
// Registration or Forgot Password will need this function to get verification code
func RequireEmailCode(email string) error {
	// Check if this email had required
	if utils.HasCacheKey("account", email) {
		return "Require too frequent, please try again later."
	}
	// Generate a verification code, and save into Cache
	code := utils.GenerateRandomNumber(4)
	if err := utils.SetCacheValue("account", email, code, 60); err != nil {
		return err
	}
	// Use Email service to send an email
	content := utils.GetVerificationText("Corp", code)
	if err := utils.SendEmail(email, content); err != nil {
		return err
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
func Register(email, password, activateUrl string) error {
	// FIXME: temp implementation for dealing with password verification
	account := model.Account{ Email: email, Password: password, IsActivated: false }
	if !account.IsPasswordValid() {
		return errors.New("invalid password format")
	}
	// FIXME: replace model to middleware(because of cache)
	if model.IsAccountExist(model.Account{Email: email}) {
		return err
	}
	// FIXME: 
	if err := model.InsertAccount(&account); err != nil {
		return err
	}
	// generate an activation code
	code := utils.GenerateRandomString(64)
	if err := utils.SetCacheValue("account", "registration", code, email, 60); err != nil {
		// FIXME: Rollback? or?
		return errors.New("failed to save activation code.")
	}
	// url: http://example.com/service/action?code=code
	// action in activateUrl supposes to be mapped to ActivateAccountWithHashCode below.
	url := fmt.Sprintf("%s?code=%s", activateUrl, code)
	content := utils.GetActivationText("company", url)
	if err := utils.SendEmail(email, content); err != nil {
		// FIXME: Rollback? or?
		return errors.New("failed to send email")
	}
	return nil
}

// after Register(situation a) do
func ActivateAccountWithHashCode(code string) error {
	// see above to see how the code be saved in cache.
	var email string
	if err := utils.GetCacheValue("account", "registration", code, &email); err != nil {
		return err
	}
	// FIXME:
	if err := model.UpdateActivatedState(model.Account{Email: email, IsActivated: true}); err != nil {
		return err
	}
	return nil
}

// situation b
func RegisterWithEmailAndCode(email, password, code string) error {
	// check if the code is valid
	var saved string
	if err = utils.GetCacheValue("account", "verify", email, &saved); err != nil {
		return errors.New("unmatched verification code.")
	}
	if saved == code {
		return errors.New("invalid verification code.")
	}
	// FIXME: 
	if model.IsAccountExist(model.Account{Email: email}) {
		return err
	}
	// FIXME: 
	account := model.Account{ Email: email, Phone: phone, Password: password, IsActivated: true }
	if err := model.InsertAccount(&account); err != nil {
		return err
	}
	return nil
}

// situation c
func RegisterWithPhoneAndCode(phone, password, code string) error {
	// check if the code is valid
	var saved string
	if err = utils.GetCacheValue("account", "verify", phone, &saved); err != nil {
		return errors.New("unmatched verification code.")
	}
	if saved == code {
		return errors.New("invalid verification code.")
	}
	// FIXME: 
	if model.IsAccountExist(model.Account{Phone: phone}) {
		return err
	}
	// FIXME: 
	account := model.Account{ Email: email, Phone: phone, Password: password, IsActivated: true }
	if err := model.InsertAccount(&account); err != nil {
		return err
	}
	return nil
}

// situation d
func RegisterWithUsername(username, email, phone, password string) error {
	// FIXME: 
	if model.IsAccountExist(model.Account{Username: username}) {
		return err
	}
	// FIXME: 
	if model.IsAccountExist(model.Account{Email: email}) {
		return err
	}
	// FIXME: 
	if model.IsAccountExist(model.Account{Phone: phone}) {
		return err
	}
	// FIXME: 
	account := model.Account{ Username: username, Email: email, Phone: phone, Password: password, IsActivated: true }
	if err := model.InsertAccount(&account); err != nil {
		return err
	}
	return nil
}

// How Login Works?
// 1. fill login form with UserId/Email/Phone, and submit(POST)
// 2. server side handle the this requirement.
// 3. response token or error text
func Login(username, email, phone, password string, token *string) error {
	// FIXME: 
	account := model.Account{ Username: username, Email: email, Phone: phone, Password: password }
	if err := model.VerifyAccountWithPassword(&account); err != nil {
		return err
	}
	// Further feature: 
	// if unmatched times > 5, account will be locked(implements in the futher)
	*token = utils.GenerateToken(account.ID, account.Username)
	return nil
}

// feature for the futher... 
// you will need to answer some questions to unlock.
// think about what questions I should use?
func UnlockAccount(username, email, phone string) error {
	// FIXME: 
	account := model.Account{ Username: username, Email: email, Phone: phone }
	if err := model.UpdateLockedState(account); err != nil {
		return errors.New("failed to unlock account")
	}
	return nil
}

// If just only use JWT, there is nothing we can do, because everything is useless. 
// JWT itself has expired time
func Logout(userId, email, phone string) error {
	// If JWT, there is no way to logout in server side, unless JWT had expired.
	// The only way is through client side
	// Or use Redis to store it when logged in, and remove it here.
	// But that is not a good way to use JWT. Annoyance...
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
	var identifier, saved string
	// only one of email and phone is available, then the other one supposes to be empty
	if identifier = email; len(identifier) == 0 {
		identifier = phone
	}
	if err := utils.GetCacheValue("account", "verify", identifier, &saved); err != nil {
		return errors.New("unmatched verification code.")
	}
	if saved != code {
		return errors.New("invalid verification code.")
	}
	// FIXME: 
	// update password
	account := model.Account{ Username: username, Email: email, Phone: phone, Password: password }
	if err := model.UpdatePassword(&account); err != nil {
		return errors.New("failed to update password")
	}
	return nil
}

// situation b
func Forgot(email, pageUrl string) error {
	// generate hash code(key: hash_code, value: email)
	code := utils.GenerateRandomString(64)
	if err := utils.SetCacheValue("account", "forgot", code, email); err != nil {
		return err
	}
	// url: http://example.com/account/action?code=code
	// action in pageUrl supposes to be mapped to ResetPasswordByHashCode below.
	url := fmt.Sprintf("%s?code=%s", pageUrl, code)
	content := utils.GetForgotPasswordText("company", url)
	if err := utils.SendEmail(email, content); err != nil {
		return err
	}
	return nil
}

// after Forgot(situation b)
func ResetPasswordByHashCode(code, password string) error {
	// get email by hashcode(key: hash_code, value: email)
	var email string
	if err := utils.GetCacheValue("account", "forgot", code, &email); err != nil {
		return err
	}
	// FIXME: 
	// update password
	if err := model.UpdatePassword(&Account{ Email: email, Password: password }); err != nil {
		return err
	}
	return nil
}

