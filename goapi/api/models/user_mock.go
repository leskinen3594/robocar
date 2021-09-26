package models

type UserRepositoryMock struct {
	UserRepositoryMock []UserMock `json:"users"`
}

type UserMock struct {
	UserID   int    `json:"uid"`
	RobotID  int    `json:"rbt_id"`
	Username string `json:"uname"`
	Password string `json:"passwd"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Is_Staff int    `json:"is_staff"`
}

// Import data from json file
