// Hexagonal Architecture ; Adapter Repository
package models

import (
	"github.com/jmoiron/sqlx"
)

type userRepositoryDB struct {
	db *sqlx.DB
}

func NewUserRepositoryDB(db *sqlx.DB) UserRepository {
	return userRepositoryDB{db: db}
}

// Get user all
func (r userRepositoryDB) GetUserAll() ([]User, error) {
	users := []User{}
	qry := "SELECT uid, urbt_id, uname, passwd, email, phone, is_staff FROM Users"
	err := r.db.Select(&users, qry)

	if err != nil {
		return nil, err
	}

	users = handleNull(users)

	return users, nil
}

// Get user by id
func (r userRepositoryDB) GetUserById(id int) (*User, error) {
	user := User{}
	qry := "SELECT uid, urbt_id, uname, passwd, email, phone, is_staff FROM Users WHERE uid=?"
	err := r.db.Get(&user, qry, id)

	if err != nil {
		return nil, err
	}

	user = handleNullOneRow(user)

	return &user, nil
}

// Create new user
func (r userRepositoryDB) CreateUser(user User) (*User, error) {
	return nil, nil
}

// convert null to zero type value
func handleNull(users []User) []User {
	for index := range users {
		// Phone
		if users[index].Phone.Valid {
			users[index].PhoneString = users[index].Phone.String
		} else {
			users[index].PhoneString = "" // Zero type value
		}
	}

	return users
}

func handleNullOneRow(usr User) User {
	// Phone
	if usr.Phone.Valid {
		usr.PhoneString = usr.Phone.String
	} else {
		usr.PhoneString = "" // Zero type value
	}

	return usr
}
