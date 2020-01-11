package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pangold/auth/service"
)

type UserController struct {
	service *service.User
}

func NewUserController(service *service.User) *UserController {
	return &UserController{
		service: service,
	}
}

func (this *UserController) Index(c *gin.Context) {
	// this.List(c)
}

func (this *UserController) Show(c *gin.Context) {

}

func (this *UserController) Create(c *gin.Context) {

}

func (this *UserController) Update(c *gin.Context) {

}

func (this *UserController) Delete(c *gin.Context) {

}