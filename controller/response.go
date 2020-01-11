package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func failure(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H {
		"status": code,
		"message": msg,
		"data": "",
	})
}

func success(ctx *gin.Context, data string) {
	ctx.JSON(http.StatusOK, gin.H {
		"status": http.StatusOK,
		"message": "",
		"data": data,
	})
}

