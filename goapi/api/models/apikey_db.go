package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type apiRepositoryDB struct {
	db *sqlx.DB
}

func NewAPIkeyRepositoryDB(db *sqlx.DB) APIkeyRepository {
	return apiRepositoryDB{db: db}
}

// Find api key
func (r apiRepositoryDB) CheckAPIkey(api_key, uname string) (*APIkey, error) {
	key := APIkey{}
	qry := fmt.Sprintf(`SELECT APIkeys.api_id, APIkeys.api_key, Users.uname, Robots.mac_addr FROM Users
						INNER JOIN Robots ON Users.urbt_id = Robots.rbt_id
						INNER JOIN APIkeys ON Users.uid = APIkeys.api_uid
						WHERE APIkeys.api_key LIKE '%s'
						AND Users.uname LIKE '%s'`, api_key, uname)
	err := r.db.Get(&key, qry)

	if err != nil {
		return nil, err
	}

	return &key, nil
}
