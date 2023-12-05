package models

import (
	"time"

	"god-of-wealth/common/utils"

	"github.com/Dparty/dao/common"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AreaCode               string `json:"areaCode" gorm:"type:CHAR(6)"`
	PhoneNumber            string `json:"phoneNumber" gorm:"type:VARCHAR(11);index:phonenumber_index"`
	AreaCodeAndPhoneNumber string `json:"areaCodeAndPhoneNumber" gorm:"type:VARCHAR(21);unique"`
	Password               string `json:"password" gorm:"type:CHAR(128)"`
	Salt                   []byte `json:"salt"`
	Role                   string `json:"role" gorm:"type:VARCHAR(128)"`
	Gender                 string `json:"gender" gorm:"type:VARCHAR(128)"`
	Birthday               time.Time
	Name                   string            `gorm:"type:VARCHAR(128)"`
	From                   string            `gorm:"type:VARCHAR(128)"`
	Services               common.StringList `gorm:"type:JSON"`
	Points                 int64             `gorm:"default:0"`
}

func (a Account) ID() uint {
	return a.Model.ID
}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.Model.ID = utils.GenerteId()
	return err
}
