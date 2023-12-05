package models

import (
	"god-of-wealth/common/utils"
	"time"

	"gorm.io/gorm"
)

type VerificationCode struct {
	gorm.Model
	Type        string `gorm:"type:VARCHAR(12)"`
	AreaCode    string `gorm:"type:CHAR(6)"`
	PhoneNumber string `gorm:"type:CHAR(11);index:phone_email_index"`
	Email       string `json:"email" gorm:"type:CHAR(128);index:verification_email_index"`
	Code        string `json:"code" gorm:"type:VARCHAR(12)"`
	Tries       int64
}

func (a *VerificationCode) BeforeCreate(tx *gorm.DB) (err error) {
	a.Model.ID = utils.GenerteId()
	return err
}

func (v *VerificationCode) Expired() bool {
	return time.Now().After(v.CreatedAt.Add(time.Minute * 10))
}
