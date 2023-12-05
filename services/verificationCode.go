package services

import (
	"hotelfortuna/dao/repositories"
	verificationcode "hotelfortuna/verificationCode"
	"sync"
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

var mu sync.Mutex

func (v VerificationService) CreateVerificationCode(areaCode, phoneNumber string) bool {
	mu.Lock()
	verificationCode := v.verificationCodeRepository.GetByPhoneNumber(areaCode, phoneNumber)
	if verificationCode == nil || time.Now().After(verificationCode.CreatedAt.Add(time.Minute)) {
		v.verificationCodeRepository.DeletePhoneNumber(areaCode, phoneNumber)
		code := random.RandomNumberString(6)
		verificationcode.SendVerificationCode(sms.PhoneNumber{
			AreaCode: areaCode,
			Number:   phoneNumber,
		}, code)
		v.verificationCodeRepository.CreatePhone(areaCode, phoneNumber, code)
		mu.Unlock()
		return true
	}
	mu.Unlock()
	return false
}
