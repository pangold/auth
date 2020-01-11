package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pangold/auth/middleware"
)

type Auth struct {
	token middleware.Token
}

func NewAuthController(token middleware.Token) *Auth {
	return &Auth{
		token: token,
	}
}

func (f *Auth) Filter(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	var cid, uid, name string
	if err := f.token.CheckToken(token, &cid, &uid, &name); err != nil {
		ctx.Abort()
		return
	}
	ctx.Next()
}

