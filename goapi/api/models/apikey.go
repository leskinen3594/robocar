package models

// Entity
type APIkey struct {
	// APIkeys table
	ApiID  int    `db:"api_id"`
	ApiKey string `db:"api_key"`

	// Query ; uname, mac_addr WHERE api_key LIKE '?'
	Username  string `db:"uname"`
	MacAdress string `db:"mac_addr"`
}

type APIkeyRepository interface {
	CheckAPIkey(string) (*APIkey, error)
}
