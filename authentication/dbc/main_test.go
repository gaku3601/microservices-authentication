package dbc

import (
	"database/sql"
	"testing"

	"github.com/gaku3601/study-microservices/authentication/config"
)

func TestDBConnect(t *testing.T) {
	config.SetConfig("../config")
	//DB接続テスト
	DBConnect(func(db *sql.DB) {
		err := db.Ping()
		if err == nil {
		} else {
			t.Log(err)
			t.Errorf("DB接続エラーです。config周りの設定がおかしいかもです。")
		}
	})
}
