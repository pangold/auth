package controller

import (
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"../service"
)


type EmailForm struct {
	Email string `form:"email" json:"email" binding:"required"`
}

type ResetForm struct {
	Code     string `form:"code"     json:"code"     binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required"`
}

// Register above have already done it, 
// But just incase our client didn't receive that email or phone verfication code.
func RequireActivatedCode(c *gin.Context) {
	// Big problem about the activated link
	// In micro services, url will be changed by gateway dynamically...
	// FIXME: how to fix it
	url := "http://sample.com/activate"
	form := EmailForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: "Error: please fill out your email"})
		return
	}
	if _, err := service.GenerateActivationCode(form.Email, url); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	ResponseJson(c, ResponseOption{Code: http.StatusOK, Msg: "Require activated code successfully"})
}

// Register with Email(unactivated state)
func Register(c *gin.Context) {
	url := "http://sample.com/activate"
	form := LoginForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: "Error: Please fill out the form"})
		return
	}
	if _, err := service.Register(form.Email, form.Password, url); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	ResponseJson(c, ResponseOption{Code: http.StatusOK, Msg: "Registry successfully"})
}

// GET
func Activate(c *gin.Context) {
	code := c.Query("code")
	if len(code) == 0 {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: "Error: Empty activated code"})
		return
	}
	if err := service.ActivateAccountWithHashCode(code); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	ResponseJson(c, ResponseOption{Code: http.StatusOK, Msg: "Activation success"})
}

func Forgot(c *gin.Context) {
	// Big problem about the activated link
	// In micro services, url will be changed by gateway dynamically...
	// FIXME: how to fix it
	url := "http://sample.com/reset"
	form := EmailForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: "Error: Empty email"})
		return
	}
	if _, err := service.Forgot(form.Email, url); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	ResponseJson(c, ResponseOption{Code: http.StatusOK, Msg: "Require resetting password successfully"})
}

func ResetPassword(c *gin.Context) {
	form := ResetForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: "Error: Please fill out the form completely"})
		return
	}
	if form.Password != form.ConfirmPassword {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: "Error: Password and confirm password should be the same"})
		return
	}
	if err := service.ResetPasswordByHashCode(form.Code, form.Password); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	ResponseJson(c, ResponseOption{Code: http.StatusOK, Msg: "Reset password successfully"})
}
