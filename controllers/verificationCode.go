package controllers

import (
	"god-of-wealth/controllers/models"
	"god-of-wealth/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewVerificationApi() VerificationApi {
	return VerificationApi{services.GetAccountService(), services.NewVerificationService()}
}

type VerificationApi struct {
	AccountService      *services.AccountService
	VerificationService *services.VerificationService
}

func (v VerificationApi) CreateVerificationCode(ctx *gin.Context) {
	var createVerificationCodeRequest models.CreateVerificationCodeRequest
	ctx.ShouldBindJSON(&createVerificationCodeRequest)
	if createVerificationCodeRequest.PhoneNumber == nil {
		ctx.JSON(http.StatusBadRequest, "")
		return
	}
	if createVerificationCodeRequest.Purpose == "REGISTER" {
		account := v.AccountService.GetAccountByPhoneNumber(createVerificationCodeRequest.PhoneNumber.AreaCode, createVerificationCodeRequest.PhoneNumber.Number)
		if account != nil {
			ctx.JSON(http.StatusConflict, gin.H{
				"code":    "80000",
				"message": "account exists",
			})
			return
		}
	}
	if createVerificationCodeRequest.Purpose == "LOGIN" {
		account := v.AccountService.GetAccountByPhoneNumber(createVerificationCodeRequest.PhoneNumber.AreaCode, createVerificationCodeRequest.PhoneNumber.Number)
		if account == nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":    "80002",
				"message": "account not exists",
			})
			return
		}
	}
	ok := v.VerificationService.CreateVerificationCode(createVerificationCodeRequest.PhoneNumber.AreaCode, createVerificationCodeRequest.PhoneNumber.Number)
	if !ok {
		ctx.JSON(http.StatusConflict, gin.H{
			"code":    "80001",
			"message": "verification code conflict",
		})
		return
	}
	ctx.String(http.StatusCreated, "")
}
