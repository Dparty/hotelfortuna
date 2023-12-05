package repositories

import (
	"hotelfortuna/dao"
	"hotelfortuna/dao/models"
)

func init() {
	dao.GetDB().AutoMigrate(&models.Account{}, &models.VerificationCode{})
}
