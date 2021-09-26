package service

import (
	"database/sql"
	"errors"
	"goapi/api/models"
)

type apiService struct {
	apiRepo models.APIkeyRepository
}

func NewAPIkeyService(apiRepo models.APIkeyRepository) APIkeyService {
	return apiService{apiRepo: apiRepo}
}

// GET username, mac address from api key
func (s apiService) GetUserFromKey(api_key string) (*APIkeyResponse, error) {
	user_api, err := s.apiRepo.CheckAPIkey(api_key)

	if err == sql.ErrNoRows {
		return nil, errors.New("key not found mismatch")
	}

	apikeyResponse := APIkeyResponse{
		Username:  user_api.Username,
		MacAdress: user_api.MacAdress,
	}

	return &apikeyResponse, nil
}
