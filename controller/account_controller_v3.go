package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"../service"
)

func RegisterRaw(c *gin.Context) {
	form := RegistryForm{}
	if err := c.ShouldBindJSON(&form); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: "Error: Please fill out your account name"})
		return
	}
	if form.UserName == "" {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: "Error: Empty username"})
		return
	}
	if form.Password != form.ConfirmPassword {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: "Error: Password and confirm password should be the same"})
		return
	}
	if err := service.RegisterWithUsername(form.UserName, form.Email, form.Phone, form.Password); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	ResponseJson(c, ResponseOption{Code: http.StatusOK, Msg: "Register successfully"})
}

