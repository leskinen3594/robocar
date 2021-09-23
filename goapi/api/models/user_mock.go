package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type UserRepositoryMock struct {
	UserRepositoryMock []UserMock `json:"Users"`
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

func UserRepoMock() {
	file, _ := ioutil.ReadFile("./goapi/resource/mock/user.json")

	data := UserRepositoryMock{}

	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(data.UserRepositoryMock); i++ {
		fmt.Println("Product Id: ", data.UserRepositoryMock[i].UserID)
		fmt.Println("Quantity: ", data.UserRepositoryMock[i].Username)
	}
}
