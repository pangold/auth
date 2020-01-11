package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pangold/auth/model"
	"net/http"
)

// @Summary Get verification code
// @Tags Account
// @version 2.0
// @Accept application/x-json-stream
// @Param email string true "option 1 to receive verification code"
// @Param phone string true "option 2 to receive verification code"
// @Success 200
// @Router /api/v2/account/get_vcode [post]
func (this *AuthController) GetVCode(ctx *gin.Context) {
	form := model.Account{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		failure(ctx, http.StatusBadRequest, "invalid params")
		return
	}
	if err := this.service.RequestVCode(form); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}

// @Summary Sign up with verification code, must specify email or phone to receive verification code
// @Tags Account
// @version 2.0
// @Accept application/x-json-stream
// @Param userId string true "option 1: user name(must specify at lease one as account in 3 options)"
// @Param email string true "option 2: email(must specify at lease one as account in 3 options)"
// @Param phone string true "option 3: phone number(must specify at lease one as account in 3 options)"
// @Param verification_code string true ""
// @Param password string true "new password(frontend developers must confirm password before posting the form)"
// @Success 200
// @Router /api/v2/account/sign_up [post]
func (this *AuthController) SignUpWithVCode(ctx *gin.Context) {
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

// @Summary Sign in with verification code
// @Tags Account
// @version 2.0
// @Accept application/x-json-stream
// @Param account "could be either of email or phone that can receive verification code"
// @Param verification code "verification code"
// @Success 200
// @Router /api/v2/account/sign_in [post]
func (this *AuthController) SignInWithVCode(ctx *gin.Context) {
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

// @Summary Reset password by verification code. It's v2 of forgot password
// @Tags Account
// @version 2.0
// @Accept application/x-json-stream
// @Param email or phone string true "email or phone number"
// @Param verification code "verification code"
// @Param password "new password"
// @Success 200
// @Router /api/v2/account/reset [post]
func (this *AuthController) ResetWithVCode(ctx *gin.Context) {
	form := model.Account{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		failure(ctx, http.StatusBadRequest, "invalid params")
		return
	}
	if err := this.service.ResetByVCode(form); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}

// @Summary Lock account
// @Tags Account
// @version 2.0
// @Accept application/x-json-stream
// @Param account "could be anyone of userId, email or phone"
// @Success 200
// @Router /api/account/lock [post]
func (this *AuthController) Lock(ctx *gin.Context) {
	form := model.Account{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		failure(ctx, http.StatusBadRequest, "invalid params")
		return
	}
	if err := this.service.Lock(form, true); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}

// @Summary Unlock account
// @Tags Account
// @version 2.0
// @Accept application/x-json-stream
// @Param account "could be anyone of userId, email or phone"
// @Success 200
// @Router /api/account/unlock [post]
func (this *AuthController) Unlock(ctx *gin.Context) {
	form := model.Account{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		failure(ctx, http.StatusBadRequest, "invalid params")
		return
	}
	if err := this.service.Lock(form, false); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}

// @Summary Bind email
// @Tags Account
// @version 2.0
// @Accept application/x-json-stream
// @Param account "could be anyone of userId, email or phone"
// @Success 200
// @Router /api/account/bind_email [post]
func (this *AuthController) BindEmail(ctx *gin.Context) {
	form := model.Account{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		failure(ctx, http.StatusBadRequest, "invalid params")
		return
	}
	if err := this.service.BindEmail(form, true); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}

// @Summary Unbind email
// @Tags Account
// @version 2.0
// @Accept application/x-json-stream
// @Param account "could be anyone of userId, email or phone"
// @Success 200
// @Router /api/account/unbind_email [post]
func (this *AuthController) UnbindEmail(ctx *gin.Context) {
	form := model.Account{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		failure(ctx, http.StatusBadRequest, "invalid params")
		return
	}
	if err := this.service.BindEmail(form, false); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}

// @Summary Bind phone
// @Tags Account
// @version 2.0
// @Accept application/x-json-stream
// @Param account "could be anyone of userId, email or phone"
// @Success 200
// @Router /api/account/bind_phone [post]
func (this *AuthController) BindPhone(ctx *gin.Context) {
	form := model.Account{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		failure(ctx, http.StatusBadRequest, "invalid params")
		return
	}
	if err := this.service.BindPhone(form, true); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}

// @Summary Unbind phone
// @Tags Account
// @version 2.0
// @Accept application/x-json-stream
// @Param account "could be anyone of userId, email or phone"
// @Success 200
// @Router /api/account/unbind_phone [post]
func (this *AuthController) UnbindPhone(ctx *gin.Context) {
	form := model.Account{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		failure(ctx, http.StatusBadRequest, "invalid params")
		return
	}
	if err := this.service.BindPhone(form, false); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}
