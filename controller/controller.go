package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type AccountForm struct {
	UserName         string `form:"username" json:"username"`
	Email            string `form:"email"    json:"email"`
	Phone            string `form:"phone"    json:"phone"`
	Code             string `form:"code"     json:"code"`
	Password         string `form:"password" json:"password"`
	ConfirmPassword  string `form:"confirm_password" json:"confirm_password"`
}

type AuthHeader struct {
	Token            string `header:"Token" binding:"required"`
}

type ResponseOption struct {
	Code             int
	Msg              string
	Data             string
}


func responseJson(c *gin.Context, r ResponseOption) {
	c.JSON(r.Code, gin.H {
		"status": r.Code,
		"message": r.Msg,
		"data": r.Data,
	})
}

func failure(c *gin.Context, code int, msg string) {
	responseJson(c, ResponseOption{Code: code, Msg: msg})
}

func success(c *gin.Context, data string) {
	responseJson(c, ResponseOption{Code: http.StatusOK, Data: data})
}

