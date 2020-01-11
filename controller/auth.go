package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pangold/auth/model"
	"gitlab.com/pangold/auth/service"
	"net/http"
)

type AuthController struct {
	service *service.Auth
}

func NewAuthController(s *service.Auth) *AuthController {
	return &AuthController{
		service: s,
	}
}

// @Summary Sign up: needs to activate your account by activation-url if sign-up-strategy=activation
// @Tags Filter
// @version 1.0
// @Accept application/x-json-stream
// @Param userId string true "option 1: user name(must specify at lease one as account in 2 options)"
// @Param email string true "option 2: email(must specify at lease one as account in 2 options); must be specified if sign-up-strategy=activation"
// @Param password string true "new password(frontend developers must confirm password before posting the form)"
// @Success 200
// @Router /api/v1/account/sign_up [post]
func (this *AuthController) SignUp(ctx *gin.Context) {
	form := model.Account{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		failure(ctx, http.StatusBadRequest, "invalid params")
		return
	}
	if err := this.service.Register(form); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}

// @Summary Sign up: didn't receive activation-url
// @Tags Filter
// @version 1.0
// @Accept application/x-json-stream
// @Param email string true "email that uses to receive activation-url"
// @Success 200
// @Router /api/v1/account/activation [post]
func (this *AuthController) GetActivationUrl(ctx *gin.Context) {
	form := model.Account{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		failure(ctx, http.StatusBadRequest, "invalid params")
		return
	}
	if err := this.service.GetActivationUrl(form); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}

// @Summary Sign up: received activation-url and activate
// @Tags Filter
// @version 1.0
// @Success 200
// @Router /api/v1/account/activate?email={email}&code={code} [get]
func (this *AuthController) Activate(ctx *gin.Context) {
	c := ctx.Query("code")
	if err := this.service.Activate(c); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}

// @Summary Sign in with password
// @Tags Filter
// @version 1.0
// @Accept application/x-json-stream
// @Param account "could be anyone of userId, email or phone"
// @Param password string true "password"
// @Success 200
// @Router /api/v1/account/sign_in [post]
func (this *AuthController) SignIn(ctx *gin.Context) {
	form := model.Account{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		failure(ctx, http.StatusBadRequest, "invalid params")
		return
	}
	token, err := this.service.Login(form)
	if err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, token)
}

// @Summary Sign in with verification code
// @Tags Filter
// @version 1.0
// @Success 200
// @Router /api/v1/account/sign_out [post]
func (this *AuthController) SignOut(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if err := this.service.Logout(token); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}

// @Summary Forgot password: request a reset-password-link
// @Tags Filter
// @version 1.0
// @Accept application/x-json-stream
// @Param email string true "email address that can receive e-mail to get reset password url"
// @Success 200
// @Router /api/v1/account/forgot [post]
func (this *AuthController) Forgot(ctx *gin.Context) {
	form := model.Account{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		failure(ctx, http.StatusBadRequest, "invalid params")
		return
	}
	if err := this.service.Forgot(form); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}

// @Summary Reset password by hash code in reset-password-link.
// @Tags Filter
// @version 1.0
// @Accept application/x-json-stream
// @Param hash_code "hash code that contains with the url in email"
// @Param password "new password(front-end developers need to confirm password)"
// @Success 200
// @Router /api/v1/account/reset [post]
func (this *AuthController) Reset(ctx *gin.Context) {
	form := model.Account{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		failure(ctx, http.StatusBadRequest, "invalid params")
		return
	}
	if err := this.service.ResetByHashCode(form); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}

// @Summary Check if user id is exist
// @Tags Filter
// @version 1.0
// @Accept application/x-json-stream
// @Param userId "user id"
// @Success 200
// @Router /api/v1/account/is_uid_exist [post]
func (this *AuthController) IsUserIdExist(ctx *gin.Context) {
	form := model.Account{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		failure(ctx, http.StatusBadRequest, "invalid params")
		return
	}
	if !this.service.IsUserIdExist(form) {
		success(ctx, "false")
		return
	}
	success(ctx, "true")
}

// @Summary Check if email is exist
// @Tags Filter
// @version 1.0
// @Accept application/x-json-stream
// @Param email "email"
// @Success 200
// @Router /api/v1/account/is_email_exist [post]
func (this *AuthController) IsEmailExist(ctx *gin.Context) {
	form := model.Account{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		failure(ctx, http.StatusBadRequest, "invalid params")
		return
	}
	if !this.service.IsEmailExist(form) {
		success(ctx, "false")
		return
	}
	success(ctx, "true")
}

// @Summary Check if phone is exist
// @Tags Filter
// @version 1.0
// @Accept application/x-json-stream
// @Param phone "phone"
// @Success 200
// @Router /api/v1/account/is_phone_exist [post]
func (this *AuthController) IsPhoneExist(ctx *gin.Context) {
	form := model.Account{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		failure(ctx, http.StatusBadRequest, "invalid params")
		return
	}
	if !this.service.IsPhoneExist(form) {
		success(ctx, "false")
		return
	}
	success(ctx, "true")
}