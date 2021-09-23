// Hexagonal Architecture ; Port Service
// Business Layer ; custom response data
package service

type UserResponse struct {
	UserID   int    `json:"uid"`
	RobotID  int    `json:"rbt_id"`
	Username string `json:"uname"`
	Password string `json:"passwd"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	// Is_Staff int    `json:"is_staff"`
}

type UserService interface {
	GetUsers() ([]UserResponse, error)
	GetUser(int) (*UserResponse, error)
}
