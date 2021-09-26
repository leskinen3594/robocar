package service

type APIkeyResponse struct {
	Username  string `json:"uname"`
	MacAdress string `json:"mac"`
}

type APIkeyService interface {
	GetUserFromKey(string) (*APIkeyResponse, error)
}
