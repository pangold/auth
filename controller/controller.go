package controller

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"../service"
)

type LoginForm struct {
	Username string `form:"username" json:"username"`
	Email    string `form:"email"    json:"email"`
	Phone    string `form:"phone"    json:"phone"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Header struct {
	Token    string `header:"Token" binding:"required"`
}

type ResponseOption struct {
	Code     int
	Msg      string
	Data     string
}

func ResponseJson(c *gin.Context, r ResponseOption) {
	c.JSON(r.Code, gin.H {
		"status": r.Code,
		"message": r.Msg,
		"data": r.Data,
	})
}

func New() *gin.Engine {
	rounter := gin.Default()
	rounter.Use(middleware())
	rounter.POST("/sign_in", Login)
	rounter.POST("/sign_out", Logout)
	// Email Link(With Hash Code)
	v1 := rounter.Group("/v1")
	{
		v1.POST("/require_activated_code", RequireActivatedCode)
		v1.POST("/sign_up", Register)
		v1.GET("/activate", Activate)
		v1.POST("/forgot", Forgot)
		v1.POST("/reset", ResetPassword)
	}
	// Verification Code
	v2 := rounter.Group("/v2")
	{
		v2.POST("/require_verification_code", RequireVerificationCode)
		v2.POST("/sign_up", RegisterWithCode)
		v2.POST("/forgot", ForgotWithCode)
	}
	// Without any of VCode or Email Link
	v3 := rounter.Group("/v3")
	{
		v3.POST("/sign_up", RegisterRaw)
		// forgot? how? answer questions?
	}
	return rounter
}

func Login(c *gin.Context) {
	var login LoginForm
	if err := c.ShouldBindJSON(&login); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: "Error: please fill out the form completely"})
		return
	}
	token, err := service.Login(login.Username, login.Email, login.Phone, login.Password)
	if err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	// FIXME:
	data := fmt.Sprintf("{token: \"%s\"}", token)
	ResponseJson(c, ResponseOption{Code: http.StatusOK, Data: data})
}

func Logout(c *gin.Context) {
	var header Header
	if err := c.ShouldBindHeader(&header); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusUnauthorized, Msg: "unauthorized"})
		return
	}
	if err := service.Logout(header.Token); err != nil {
		ResponseJson(c, ResponseOption{Code: http.StatusBadRequest, Msg: err.Error()})
		return
	}
	ResponseJson(c, ResponseOption{Code: http.StatusOK, Data: "signed out successfully"})
}

