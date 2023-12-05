package repositories

import (
	"hotelfortuna/dao"
	"hotelfortuna/dao/models"
	"time"

	"gorm.io/gorm"
)

var verificationCodeRepository *VerificationCodeRepository

func GetVerificationCodeRepository() *VerificationCodeRepository {
	if verificationCodeRepository == nil {
		verificationCodeRepository = NewVerificationCodeRepository()
	}
	return verificationCodeRepository
}

type VerificationCodeRepository struct {
	db *gorm.DB
}

func NewVerificationCodeRepository() *VerificationCodeRepository {
	return &VerificationCodeRepository{dao.GetDB()}
}

func (r VerificationCodeRepository) CreatePhone(areaCode, phoneNumber, code string) models.VerificationCode {
	verificationCode := models.VerificationCode{
		Type:        "PHONE",
		AreaCode:    areaCode,
		PhoneNumber: phoneNumber,
		Code:        code,
	}
	r.db.Save(&verificationCode)
	return verificationCode
}

func (r VerificationCodeRepository) Delete(verificationCode *models.VerificationCode) {
	r.db.Delete(verificationCode)
}

func (r VerificationCodeRepository) GetByPhoneNumber(areaCode, phoneNumber string) *models.VerificationCode {
	var verificationCode models.VerificationCode
	ctx := r.db.Where("area_code = ?", areaCode).Where(
		"phone_number = ?", phoneNumber).Find(&verificationCode)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &verificationCode
}

func (r VerificationCodeRepository) DeleteExpired() {
	r.db.Exec("DELETE FROM verification_codes WHERE created_at <= ?", time.Now().Add(-time.Minute*10))
}

func (r VerificationCodeRepository) DeletePhoneNumber(areaCode, phoneNumber string) {
	r.db.Exec("DELETE FROM verification_codes WHERE area_code = ? AND phone_number = ?", areaCode, phoneNumber)
}
