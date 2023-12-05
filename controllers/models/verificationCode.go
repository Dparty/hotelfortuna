package models

type CreateVerificationCodeRequest struct {
	Purpose     string       `json:"purpose"`
	PhoneNumber *PhoneNumber `json:"phoneNumber"`
}
