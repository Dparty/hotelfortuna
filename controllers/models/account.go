package models

type UpdateAccountInfoRequest struct {
	Name     string   `json:"name"`
	Gender   string   `json:"gender"`
	Birthday int64    `json:"birthday"`
	Services []string `json:"services"`
}

type Account struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Gender      string      `json:"gender"`
	Birthday    int64       `json:"birthday"`
	PhoneNumber PhoneNumber `json:"phoneNumber"`
	Services    []string    `json:"services"`
	Points      int64       `json:"points"`
}

type PhoneNumber struct {
	AreaCode string `json:"areaCode"`
	Number   string `json:"number"`
}

type CreateAccountRequest struct {
	PhoneNumber      PhoneNumber `json:"phoneNumber"`
	Name             string      `json:"name"`
	Password         string      `json:"password"`
	VerificationCode string      `json:"verificationCode"`
	Gender           string      `json:"gender"`
	Birthday         int64       `json:"birthday"`
	From             string      `json:"from"`
	Services         []string    `json:"services"`
}

type CreateSessionRequest struct {
	PhoneNumber      PhoneNumber `json:"phoneNumber"`
	VerificationCode *string     `json:"verificationCode"`
	Password         *string     `json:"password"`
}

type Session struct {
	Token string `json:"token"`
}
