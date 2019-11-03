package main

import (
	"github.com/gin-gonic/gin"
	"./controller"
)

func main() {
	rounter := gin.Default()
	// rounter.Use(middleware())
	v1 := rounter.Group("/api/v1")
	{
		// Email Link(With Hash Code)
		v1.POST("/sign_in", controller.Login)
		v1.POST("/sign_out", controller.Logout)
		v1.POST("/sign_up", controller.RegisterUnactivatedState)
		// In case we don't receive activation code
		v1.POST("/activation_code", controller.RequireActivationCode)
		v1.GET ("/activate", controller.Activate)
		v1.POST("/forgot", controller.Forgot)
		v1.POST("/reset", controller.ResetByHashCode)
		// For checking
		v1.POST("/is_username_exist", controller.IsUsernameExist)
		v1.POST("/is_email_exist", controller.IsEmailExist)
		v1.POST("/is_phone_exist", controller.IsPhoneExist)
	}
	v2 := rounter.Group("/api/v2")
	{
		// Verification Code
		v2.POST("/sign_in", controller.Login)
		v2.POST("/sign_out", controller.Logout)
		v2.POST("/sign_up", controller.RegisterWithCode) // needs vcode? you should check if it is exist first
		v2.POST("/verification_code", controller.RequireVerificationCode)
		v2.POST("/reset", controller.ResetWithVerificationCode) // compare to v1::reset, there is an additional param phone/email
		// For checking
		v2.POST("/is_username_exist", controller.IsUsernameExist)
		v2.POST("/is_email_exist", controller.IsEmailExist)
		v2.POST("/is_phone_exist", controller.IsPhoneExist)
	}
	v3 := rounter.Group("/api/v3")
	{
		// Without any of VCode or Email Link
		v3.POST("/sign_in", controller.Login)
		v3.POST("/sign_out", controller.Logout)
		v3.POST("/sign_up", controller.RegisterRaw)
		// Forgot? You want to reset your password?
		// No, so far, what you can do is contact Administrators.
		// For checking
		v3.POST("/is_username_exist", controller.IsUsernameExist)
		v3.POST("/is_email_exist", controller.IsEmailExist)
		v3.POST("/is_phone_exist", controller.IsPhoneExist)
	}
	rounter.Run(":8080")
}
