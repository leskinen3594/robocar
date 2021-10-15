// Business Logic ; Adapter Service
package service

import (
	"database/sql"
	"errors"
	"goapi/api/models"
)

type userService struct {
	userRepo models.UserRepository
}

func NewUserService(userRepo models.UserRepository) UserService {
	return userService{userRepo: userRepo}
}

// custom user all
func (s userService) GetUsers() ([]UserResponse, error) {
	users, err := s.userRepo.GetUserAll()

	if err == sql.ErrNoRows {
		return nil, errors.New("users not found")
	}

	userResponses := []UserResponse{}
	for _, user := range users {
		userResp := UserResponse{
			UserID:   user.UserID,
			RobotID:  user.RobotID,
			Username: user.Username,
			Password: user.Password,
			Email:    user.Email,
			Phone:    user.PhoneString,
			// Is_Staff: user.Is_Staff,
		}
		userResponses = append(userResponses, userResp)
	}

	return userResponses, nil
}

// custom user by id
func (s userService) GetUser(id int) (*UserResponse, error) {
	user, err := s.userRepo.GetUserById(id)

	if err == sql.ErrNoRows {
		return nil, errors.New("this user id not found")
	}

	userResponse := UserResponse{
		UserID:   user.UserID,
		RobotID:  user.RobotID,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Phone:    user.PhoneString,
	}

	return &userResponse, nil
}
