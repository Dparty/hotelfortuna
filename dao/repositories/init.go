package repositories

import (
	"god-of-wealth/dao"
	"god-of-wealth/dao/models"
)

func init() {
	dao.GetDB().AutoMigrate(&models.Account{}, &models.VerificationCode{})
}
