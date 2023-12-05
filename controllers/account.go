package controllers

import (
	"errors"
	"god-of-wealth/common/config"
	"god-of-wealth/controllers/models"
	"god-of-wealth/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NewAccountApi() *AccountApi {
	return &AccountApi{services.GetAccountService()}
}

type AccountApi struct {
	accountService *services.AccountService
}

func (a AccountApi) GetAccount(ctx *gin.Context) {
	account := getAccount(ctx)
	if account == nil {
		return
	}
	ctx.JSON(http.StatusOK, AccountBackword(*account))
}

func (a AccountApi) UpdateAccountInfo(ctx *gin.Context) {
	account := getAccount(ctx)
	if account == nil {
		return
	}
	var updateAccountInfoRequest models.UpdateAccountInfoRequest
	ctx.ShouldBindJSON(&updateAccountInfoRequest)
	account.SetName(updateAccountInfoRequest.Name).SetBirthday(
		time.Unix(updateAccountInfoRequest.Birthday, 0)).SetGender(
		updateAccountInfoRequest.Gender)
	account.Submit()
	ctx.JSON(http.StatusOK, AccountBackword(*account))
}

func (a AccountApi) CreateAccount(ctx *gin.Context) {
	var createAccountRequest models.CreateAccountRequest
	ctx.ShouldBindJSON(&createAccountRequest)
	token, err := a.accountService.CreateAccount(
		createAccountRequest.PhoneNumber.AreaCode,
		createAccountRequest.PhoneNumber.Number,
		createAccountRequest.VerificationCode,
		createAccountRequest.Name,
		time.Unix(createAccountRequest.Birthday, 0),
		config.GetString("default.password"),
		createAccountRequest.From,
		createAccountRequest.Gender,
		createAccountRequest.Services,
	)
	switch {
	case errors.Is(err, services.ErrAccountExists):
		ctx.JSON(http.StatusConflict, gin.H{
			"code":    "40009",
			"message": err.Error(),
		})
	case errors.Is(err, services.ErrVerificationFault):
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"code":    "40006",
			"message": err.Error(),
		})
	}
	if err == nil {
		ctx.JSON(http.StatusCreated, models.Session{Token: token})
	}
}

func (a AccountApi) CreateSession(ctx *gin.Context) {
	var createSessionRequest models.CreateSessionRequest
	err := ctx.ShouldBindJSON(&createSessionRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    "40000",
			"message": err.Error(),
		})
		return
	}
	token, err := a.accountService.CreateSession(
		createSessionRequest.PhoneNumber.AreaCode,
		createSessionRequest.PhoneNumber.Number,
		createSessionRequest.Password,
		createSessionRequest.VerificationCode)
	if err != nil {
		switch err {
		case services.ErrorAccountNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":    "40002",
				"message": err.Error(),
			})
		case services.ErrorUnauthorized:
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    "40001",
				"message": err.Error(),
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    "50000",
				"message": err.Error(),
			})
		}
		return
	}
	ctx.JSON(http.StatusCreated, models.Session{Token: token})
}
