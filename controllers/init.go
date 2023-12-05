package controllers

import (
	"hotelfortuna/services"
	"net/http"

	"github.com/Dparty/common/server"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	accountApi := NewAccountApi()
	verificationApi := NewVerificationApi()
	server.MetricsMiddleware(router)
	router.Use(server.CorsMiddleware())
	router.Use(accountApi.accountService.Auth())
	router.GET("/me", accountApi.GetAccount)
	router.PUT("/me", accountApi.UpdateAccountInfo)
	router.POST("/sessions", accountApi.CreateSession)
	router.POST("/verification", verificationApi.CreateVerificationCode)
	router.POST("/accounts", accountApi.CreateAccount)
	router.GET("/ddFaFYgBYU.txt", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "6d39d050c4ab6af0e647c8337b037bf0")
	})
	router.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"version": "0.0.3",
		})
	})
	router.Run(":8080")
}

func getAccount(ctx *gin.Context) *services.Account {
	accountInterface, ok := ctx.Get("account")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, "")
		return nil
	}
	account, ok := accountInterface.(services.Account)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, "")
		return nil
	}
	return &account
}
