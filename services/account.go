package services

import (
	"errors"
	"hotelfortuna/common/utils"
	"hotelfortuna/dao/models"
	daoModel "hotelfortuna/dao/models"
	"hotelfortuna/dao/repositories"
	"time"

	dutils "github.com/Dparty/common/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var accountRepository = repositories.GetAccountRepository()

var accountService *AccountService

func GetAccountService() *AccountService {
	if accountService == nil {
		accountService = NewAccountService()
	}
	return accountService
}

func NewAccountService() *AccountService {
	return &AccountService{
		repositories.NewAccountRepository(),
		repositories.GetVerificationCodeRepository(),
	}
}

type AccountService struct {
	accountRepository          *repositories.AccountRepository
	verificationCodeRepository *repositories.VerificationCodeRepository
}

var ErrAccountExists = errors.New("account exists")
var ErrVerificationFault = errors.New("verification fault")

func (a AccountService) CreateAccount(
	areaCode,
	phoneNumber,
	code,
	name string,
	birthday time.Time,
	password string,
	from string,
	gender string,
	services []string) (string, error) {
	if account := a.accountRepository.GetByPhoneNumber(areaCode, phoneNumber); account != nil {
		return "", ErrAccountExists
	}
	verificationCode := a.verificationCodeRepository.GetByPhoneNumber(areaCode, phoneNumber)
	if verificationCode == nil || verificationCode.Expired() || verificationCode.Code != code {
		return "", ErrVerificationFault
	}
	hashed, salt := utils.HashWithSalt(password)
	account := daoModel.Account{
		Password:               hashed,
		Salt:                   salt,
		AreaCode:               areaCode,
		PhoneNumber:            phoneNumber,
		AreaCodeAndPhoneNumber: areaCode + " " + phoneNumber,
		Name:                   name,
		Birthday:               birthday,
		From:                   from,
		Services:               services,
		Gender:                 gender,
	}
	a.accountRepository.Create(&account)
	a.verificationCodeRepository.Delete(verificationCode)
	expiredAt := time.Now().AddDate(1, 0, 0).Unix()
	token, err := utils.SignJwt(
		utils.UintToString(account.ID()),
		expiredAt,
	)
	if err != nil {
		return "", nil
	}
	return token, nil
}

var ErrorAccountNotFound = errors.New("account not found")
var ErrorUnauthorized = errors.New("unauthorized")

func (a AccountService) CreateSession(
	areaCode,
	phoneNumber string,
	password *string,
	verificationCode *string) (string, error) {
	account := a.accountRepository.GetByPhoneNumber(areaCode, phoneNumber)
	if account == nil {
		return "", ErrorAccountNotFound
	}
	if password == nil && verificationCode == nil {
		return "", ErrorUnauthorized
	}
	if password != nil && !utils.PasswordsMatch(account.Password, *password, account.Salt) {
		return "", ErrorUnauthorized
	}
	if verificationCode != nil {
		a.verificationCodeRepository.DeleteExpired()
		code := a.verificationCodeRepository.GetByPhoneNumber(areaCode, phoneNumber)
		if code == nil {
			return "", ErrorUnauthorized
		}
		if code.Code != *verificationCode {
			return "", ErrorUnauthorized
		}
		a.verificationCodeRepository.Delete(code)
	}
	expiredAt := time.Now().AddDate(1, 0, 0).Unix()
	token, err := utils.SignJwt(
		utils.UintToString(account.ID()),
		expiredAt,
	)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a AccountService) GetAccountByPhoneNumber(areaCode, phoneNumber string) *Account {
	account := a.accountRepository.GetByPhoneNumber(areaCode, phoneNumber)
	if account == nil {
		return nil
	}
	return &Account{*account}
}

func NewAccount(account models.Account) Account {
	return Account{account}
}

type Account struct {
	entity daoModel.Account
}

func (a *Account) Entity() daoModel.Account {
	return a.entity
}

func (a *Account) ID() uint {
	return a.entity.ID()
}

func (a *Account) Points() int64 {
	return a.entity.Points
}

func (a *Account) SetPhoneNumber(areaCode, number string) *Account {
	a.entity.AreaCode = areaCode
	a.entity.PhoneNumber = number
	a.entity.AreaCodeAndPhoneNumber = areaCode + " " + number
	return a
}

func (a *Account) SetPassword(password string) *Account {
	hashed, salt := utils.HashWithSalt(password)
	a.entity.Password = hashed
	a.entity.Salt = salt
	return a
}

func (a *Account) SetGender(gender string) *Account {
	a.entity.Gender = gender
	return a
}

func (a *Account) SetBirthday(birthday time.Time) *Account {
	a.entity.Birthday = birthday
	return a
}

func (a *Account) SetName(name string) *Account {
	a.entity.Name = name
	return a
}

func (a *Account) SetServices(services []string) *Account {
	a.entity.Services = dutils.RemoveDuplication(services)
	return a
}

func (a *Account) Submit() *Account {
	accountRepository.Save(&a.entity)
	return a
}

func NewAuthService(inject *gorm.DB) AccountService {
	return AccountService{accountRepository: repositories.NewAccountRepository()}
}

func (a AccountService) GetAccount(id uint) *Account {
	account := a.accountRepository.GetById(id)
	if account == nil {
		return nil
	}
	return &Account{*account}
}

func (a AccountService) VerifyToken(token string) (Account, error) {
	auth := AuthorizeByJWT(token)
	if auth.Status != Authorized {
		return Account{}, errors.New("")
	}
	account := a.GetAccount(auth.AccountId)
	if account == nil {
		return Account{}, errors.New("")
	}
	return *account, nil
}

func (a AccountService) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := Authorize(c)
		if auth.Status == Authorized {
			account := a.GetAccount(auth.AccountId)
			if account != nil {
				c.Set("account", *account)
			}
		}
		c.Next()
	}
}
