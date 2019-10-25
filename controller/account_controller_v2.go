package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"../service"
)

type RequireCodeForm struct {
	Email string `form:"email" json:"email"`
	Phone string `form:"phone" json:"phone"`
}

type RegistryForm struct {
	UserName string `form:"username" json:"username"`
	Email    string `form:"email"    json:"email"`
	Phone    string `form:"phone"    json:"phone"`
	Code     string `form:"code"     json:"code"`
	Password string `form:"password" json:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required"`
}

// Email Verification Code
func RequireVerificationCode(c *gin.Context) {
	form := RequireCodeForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: "Error: Please fill out the form"})
		return
	}
	if form.Email == "" && form.Phone == "" {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: "Error: Empty account name"})
		return
	}
	if err := service.RequireVerificationCode(form.Email, form.Phone); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	ResponseJson(c, ResponseOption{Code: http.StatusOK, Msg: "Require verification code successfully"})
}

func RegisterWithCode(c *gin.Context) {
	form := RegistryForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: "Error: Please fill out your account name"})
		return
	}
	if form.Email == "" && form.Phone == "" {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: "Error: Empty account name"})
		return
	}
	if form.Password != form.ConfirmPassword {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: "Error: Password and confirm password should be the same"})
		return
	}
	if err := service.RegisterWithCode(form.Email, form.Phone, form.Password, form.Code); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	ResponseJson(c, ResponseOption{Code: http.StatusOK, Msg: "Register successfully"})
}

func ForgotWithCode(c *gin.Context) {
	form := RegistryForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: "Error: Please fill out your account name"})
		return
	}
	if form.Email == "" && form.Phone == "" {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: "Error: Empty account name"})
		return
	}
	if form.Password != form.ConfirmPassword {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: "Error: Password and confirm password should be the same"})
		return
	}
	if err := service.ForgotWithCode(form.Email, form.Phone, form.Password, form.Code); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	ResponseJson(c, ResponseOption{Code: http.StatusOK, Msg: "Updated password successfully"})
}

