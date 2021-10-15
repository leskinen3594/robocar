package configs

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func ConnectDB() (*sqlx.DB, error) {
	// Setup db config
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.database"),
	)

	// Connect to database
	// db, err := sqlx.Open("postgres", "postgres://postgres:passwd123@localhost:5432/postgres?sslmode=disable")
	db, err := sqlx.Connect(viper.GetString("db.driver"), dsn)

	// Error handle if connection failed.
	// if err != nil {
	// 	log.Println(err)
	// }

	// Ping database
	// err = db.Ping()
	// if err != nil {
	// 	log.Panicln(err)
	// }

	return db, err
}
