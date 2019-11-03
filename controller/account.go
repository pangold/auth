package controller

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"../service"
)

func Login(c *gin.Context) {
	form := AccountForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		failure(c, http.StatusBadRequest, "Error: Incomplete form")
		return
	}
	token, err := service.Login(form.UserName, form.Email, form.Phone, form.Password)
	if err != nil {
		failure(c, http.StatusBadRequest, err.Error())
		return
	}
	// FIXME:
	success(c, fmt.Sprintf("{token: \"%s\"}", token))
}

func Logout(c *gin.Context) {
	header := AuthHeader{}
	if err := c.ShouldBindHeader(&header); err != nil {
		failure(c, http.StatusUnauthorized, "unauthorized")
		return
	}
	if err := service.Logout(header.Token); err != nil {
		failure(c, http.StatusBadRequest, err.Error())
		return
	}
	success(c, "")
}

func IsUsernameExist(c *gin.Context) {
	form := AccountForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		failure(c, http.StatusBadRequest, "Error: Incomplete form")
		return
	}
	if !service.IsUsernameExist(form.UserName) {
		failure(c, http.StatusBadRequest, "Error: Username is unavailable")
		return
	}
	success(c, "")
}

func IsEmailExist(c *gin.Context) {
	form := AccountForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		failure(c, http.StatusBadRequest, "Error: Incomplete form")
		return
	}
	if !service.IsEmailExist(form.Email) {
		failure(c, http.StatusBadRequest, "Error: Email is unavailable")
		return
	}
	success(c, "")
}

func IsPhoneExist(c *gin.Context) {
	form := AccountForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		failure(c, http.StatusBadRequest, "Error: Incomplete form")
		return
	}
	if !service.IsPhoneExist(form.Phone) {
		failure(c, http.StatusBadRequest, "Error: Phone is unavailable")
		return
	}
	success(c, "")
}

// Email Verification Code
func RequireVerificationCode(c *gin.Context) {
	form := AccountForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		failure(c, http.StatusBadRequest, "Error: Incomplete form")
		return
	}
	if form.Email == "" && form.Phone == "" {
		failure(c, http.StatusBadRequest, "Error: Email or phone is required")
		return
	}
	if err := service.RequireVerificationCode(form.Email, form.Phone); err != nil {
		failure(c, http.StatusBadRequest, err.Error())
		return
	}
	success(c, "")
}

// In case clients didn't receive that email / phone activation code.
func RequireActivationCode(c *gin.Context) {
	// FIXME: URL may be changed by gateway
	url := "http://sample.com/activate"
	form := AccountForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		failure(c, http.StatusBadRequest, "Error: Incomplete form")
		return
	}
	if form.Email == "" {
		failure(c, http.StatusBadRequest, "Error: Email is required")
		return
	}
	if _, err := service.GenerateActivationCode(form.Email, url); err != nil {
		failure(c, http.StatusBadRequest, err.Error())
		return
	}
	success(c, "")
}

// Register with Email(unactivated state)
func RegisterUnactivatedState(c *gin.Context) {
	url := "http://sample.com/activate"
	form := AccountForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		failure(c, http.StatusBadRequest, "Error: incomplete form")
		return
	}
	if form.Email == "" {
		failure(c, http.StatusBadRequest, "Error: Email is required")
		return
	}
	if _, err := service.Register(form.Email, form.Password, url); err != nil {
		failure(c, http.StatusBadRequest, err.Error())
		return
	}
	success(c, "")
}

// GET
func Activate(c *gin.Context) {
	code := c.Query("code")
	if len(code) == 0 {
		failure(c, http.StatusBadRequest, "Error: Activation code is required")
		return
	}
	if err := service.ActivateAccountWithHashCode(code); err != nil {
		failure(c, http.StatusBadRequest, err.Error())
		return
	}
	success(c, "")
}

// Have been activated
func RegisterWithCode(c *gin.Context) {
	form := AccountForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		failure(c, http.StatusBadRequest, "Error: Incomplete form")
		return
	}
	if form.Code == "" {
		failure(c, http.StatusBadRequest, "Error: Activation code is required")
		return
	}
	if form.Email == "" && form.Phone == "" {
		failure(c, http.StatusBadRequest, "Error: Email or phone is required")
		return
	}
	if form.Password != form.ConfirmPassword {
		failure(c, http.StatusBadRequest, "Error: Passwords should be the same")
		return
	}
	if err := service.RegisterWithCode(form.Email, form.Phone, form.Password, form.Code); err != nil {
		failure(c, http.StatusBadRequest, err.Error())
		return
	}
	success(c, "")
}

// Unsafe
func RegisterRaw(c *gin.Context) {
	form := AccountForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		failure(c, http.StatusBadRequest, "Error: Incomplete form")
		return
	}
	if form.UserName == "" {
		failure(c, http.StatusBadRequest, "Error: Username is required")
		return
	}
	if form.Password != form.ConfirmPassword {
		failure(c, http.StatusBadRequest, "Error: Passwords should be the same")
		return
	}
	if err := service.RegisterWithUsername(form.UserName, form.Email, form.Phone, form.Password); err != nil {
		failure(c, http.StatusBadRequest, err.Error())
		return
	}
	success(c, "")
}

func Forgot(c *gin.Context) {
	// FIXME: URL may be changed by gateway
	url := "http://sample.com/reset"
	form := AccountForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		failure(c, http.StatusBadRequest, "Error: Incomplete form")
		return
	}
	if form.Email == "" {
		failure(c, http.StatusBadRequest, "Error: Email is required")
		return
	}
	if _, err := service.Forgot(form.Email, url); err != nil {
		failure(c, http.StatusBadRequest, err.Error())
		return
	}
	success(c, "")
}

func ResetByHashCode(c *gin.Context) {
	form := AccountForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		failure(c, http.StatusBadRequest, "Error: Incomplete form")
		return
	}
	if form.Code == "" {
		failure(c, http.StatusBadRequest, "Error: Code is required")
		return
	}
	if form.Password != form.ConfirmPassword {
		failure(c, http.StatusBadRequest, "Error: Passwords should be the same")
		return
	}
	if err := service.ResetByHashCode(form.Code, form.Password); err != nil {
		failure(c, http.StatusBadRequest, err.Error())
		return
	}
	success(c, "")
}

func ResetWithVerificationCode(c *gin.Context) {
	form := AccountForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		failure(c, http.StatusBadRequest, "Error: Incomplete form")
		return
	}
	if form.Code == "" {
		failure(c, http.StatusBadRequest, "Error: Code is required")
		return
	}
	if form.Email == "" && form.Phone == "" {
		failure(c, http.StatusBadRequest, "Error: Email or phone is required")
		return
	}
	if form.Password != form.ConfirmPassword {
		failure(c, http.StatusBadRequest, "Error: Passwords should be the same")
		return
	}
	if err := service.ResetWithVerificationCode(form.Email, form.Phone, form.Password, form.Code); err != nil {
		failure(c, http.StatusBadRequest, err.Error())
		return
	}
	success(c, "")
}

