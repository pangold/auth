package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pangold/auth/middleware"
)

type Filter struct {
	token middleware.Token
}

func NewFilterController(token middleware.Token) *Filter {
	return &Filter{
		token: token,
	}
}

func (f *Filter) Filter(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	var cid, uid, name string
	if err := f.token.TokenVerification(token, &cid, &uid, &name); err != nil {
		ctx.Abort()
		return
	}
	ctx.Next()
}

