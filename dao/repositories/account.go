package repositories

import (
	"god-of-wealth/dao"
	"god-of-wealth/dao/models"

	"gorm.io/gorm"
)

var accountRepository *AccountRepository

func GetAccountRepository() *AccountRepository {
	if accountRepository == nil {
		accountRepository = NewAccountRepository()
	}
	return accountRepository
}

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{dao.GetDB()}
}

func (a AccountRepository) Create(account *models.Account) {
	a.db.Save(account)
}

func (a AccountRepository) Save(account *models.Account) {
	a.db.Save(account)
}

func (a AccountRepository) GetById(id uint) *models.Account {
	var account models.Account
	ctx := a.db.Find(&account, id)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &account
}

func (a AccountRepository) GetByPhoneNumber(areaCode, phoneNumber string) *models.Account {
	var account models.Account
	ctx := a.db.Where(
		"area_code = ?",
		areaCode).Where(
		"phone_number = ?",
		phoneNumber).Find(&account)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &account
}
