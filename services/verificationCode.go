package services

import (
	"god-of-wealth/dao/repositories"
	verificationcode "god-of-wealth/verificationCode"
	"time"

	"github.com/Dparty/common/sms"
	"github.com/Dparty/common/utils/random"
)

var verificationService *VerificationService

func GetVerificationService() *VerificationService {
	if verificationService == nil {
		verificationService = NewVerificationService()
	}
	return verificationService
}

func NewVerificationService() *VerificationService {
	return &VerificationService{
		repositories.GetVerificationCodeRepository(),
	}
}

type VerificationService struct {
	verificationCodeRepository *repositories.VerificationCodeRepository
}

func (v VerificationService) CreateVerificationCode(areaCode, phoneNumber string) bool {
	verificationCode := v.verificationCodeRepository.GetByPhoneNumber(areaCode, phoneNumber)
	if verificationCode == nil || time.Now().After(verificationCode.CreatedAt.Add(time.Minute)) {
		v.verificationCodeRepository.DeletePhoneNumber(areaCode, phoneNumber)
		code := random.RandomNumberString(6)
		verificationcode.SendVerificationCode(sms.PhoneNumber{
			AreaCode: areaCode,
			Number:   phoneNumber,
		}, code)
		v.verificationCodeRepository.CreatePhone(areaCode, phoneNumber, code)
		return true
	}
	return false
}
