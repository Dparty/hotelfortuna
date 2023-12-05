package controllers

import (
	"god-of-wealth/controllers/models"
	"god-of-wealth/services"

	"god-of-wealth/common/utils"
)

func AccountBackword(account services.Account) models.Account {
	entity := account.Entity()
	return models.Account{
		ID:   utils.UintToString(entity.ID()),
		Name: entity.Name,
		PhoneNumber: models.PhoneNumber{
			AreaCode: entity.AreaCode,
			Number:   entity.PhoneNumber,
		},
		Gender:   entity.Gender,
		Birthday: entity.Birthday.Unix(),
		Services: entity.Services,
		Points:   account.Points(),
	}
}
