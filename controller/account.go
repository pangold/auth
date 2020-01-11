package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.com/pangold/auth/model"
	"gitlab.com/pangold/auth/service"
	"net/http"
	"strconv"
)

type AccountController struct {
	service *service.Account
}

func NewAccountController(service *service.Account) *AccountController {
	return &AccountController{
		service: service,
	}
}

func (this *AccountController) Index(ctx *gin.Context) {
	accounts := this.service.GetAccounts()
	bytes, err := json.Marshal(accounts)
	if err != nil {
		failure(ctx, 5000, err.Error())
		return
	}
	success(ctx, string(bytes))
}

func (this *AccountController) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	iid, err := strconv.Atoi(id)
	if err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	account := this.service.GetAccount(uint64(iid))
	bytes, err := json.Marshal(account)
	if err != nil {
		failure(ctx, 5000, err.Error())
		return
	}
	success(ctx, string(bytes))
}

func (this *AccountController) Create(ctx *gin.Context) {
	account := model.Account{}
	if err := ctx.ShouldBindJSON(&account); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := this.service.Create(&account); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}

func (this *AccountController) Update(ctx *gin.Context) {
	account := model.Account{}
	if err := ctx.ShouldBindJSON(&account); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Param("id")
	iid, err := strconv.Atoi(id)
	if err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	account.ID = uint64(iid)
	if err := this.service.Update(&account); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}

func (this *AccountController) Delete(ctx *gin.Context) {
	account := model.Account{}
	if err := ctx.ShouldBindJSON(&account); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Param("id")
	iid, err := strconv.Atoi(id)
	if err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	account.ID = uint64(iid)
	if err := this.service.Delete(&account); err != nil {
		failure(ctx, http.StatusBadRequest, err.Error())
		return
	}
	success(ctx, "")
}
