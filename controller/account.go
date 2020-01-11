package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pangold/auth/service"
)

type AccountController struct {
	service *service.Account
}

func NewAccountController(service *service.Account) *AccountController {
	return &AccountController{
		service: service,
	}
}

func (this *AccountController) Index(c *gin.Context) {

}

func (this *AccountController) Show(c *gin.Context) {

}

func (this *AccountController) Create(c *gin.Context) {

}

func (this *AccountController) Update(c *gin.Context) {

}

func (this *AccountController) Delete(c *gin.Context) {

}
