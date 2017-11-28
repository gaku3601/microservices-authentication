package dbc

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
	viper "github.com/spf13/viper.git"
)

func DBConnect(fn func(db *sql.DB)) error {
	db, _ := sql.Open("postgres", "user="+viper.GetString("database.user")+" host="+viper.GetString("database.host")+" dbname="+viper.GetString("database.dbname")+" port="+viper.GetString("database.port")+" sslmode="+viper.GetString("database.sslmode"))
	defer db.Close()
	err := db.Ping()

	if err != nil {
		return errors.New("DB接続ができていません。")
	}

	fn(db)
	return nil
}
