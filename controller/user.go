package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.com/pangold/auth/model"
	"gitlab.com/pangold/auth/service"
	"net/http"
	"strconv"
)

type UserController struct {
	service *service.User
}

func NewUserController(service *service.User) *UserController {
	return &UserController{
		service: service,
	}
}

func (this *UserController) Index(ctx *gin.Context) {
	users := this.service.GetUsers()
	bytes, err := json.Marshal(users)
	if err != nil {
		failure(ctx, 5000, err.Error())
		return
	}
	success(ctx, string(bytes))
}

func (this *UserController) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	iid, err := strconv.Atoi(id)
	if err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user := this.service.GetUser(uint64(iid))
	bytes, err := json.Marshal(user)
	if err != nil {
		failure(ctx, 5000, err.Error())
		return
	}
	success(ctx, string(bytes))
}

func (this *UserController) Create(ctx *gin.Context) {
	user := model.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := this.service.Create(&user); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}

func (this *UserController) Update(ctx *gin.Context) {
	user := model.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Param("id")
	iid, err := strconv.Atoi(id)
	if err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user.ID = uint64(iid)
	if err := this.service.Update(&user); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}

func (this *UserController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	iid, err := strconv.Atoi(id)
	if err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := this.service.Delete(&model.User{ID: uint64(iid)}); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}