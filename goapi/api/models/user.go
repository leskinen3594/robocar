// Hexagonal Architecture ; Port Repository
package models

import "database/sql"

// Entity
type User struct {
	UserID   int            `db:"uid"`
	RobotID  int            `db:"urbt_id"`
	Username string         `db:"uname"`
	Password string         `db:"passwd"`
	Email    string         `db:"email"`
	Phone    sql.NullString `db:"phone"`
	Is_Staff int            `db:"is_staff"`

	// handle null value
	PhoneString string
}

type UserRepository interface {
	GetUserAll() ([]User, error)
	GetUserById(int) (*User, error)
	CreateUser(User) (*User, error)
}
