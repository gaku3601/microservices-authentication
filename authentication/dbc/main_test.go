package dbc

import (
	"database/sql"
	"testing"
)

func TestDBConnect(t *testing.T) {
	//DB接続テスト(設定を読み込んでいないので失敗する。)
	err := DBConnect(func(db *sql.DB) {})

	if err != nil {
	} else {
		t.Log(err)
		t.Errorf("DBに接続できない際のエラーが正しく設定されていません。")
	}
}
