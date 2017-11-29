package models

import (
	"database/sql"
	"testing"
)

//emailとpasswordがDBにUser登録されていることを確認する。
func TestFetchUserCheck(t *testing.T) {
	//userにemailとpasswordを渡し、格納されていることを確認する。
	user := NewUser("pro.gaku@gmail.com", "password")
	if user.email == "pro.gaku@gmail.com" {
	} else {
		t.Error("useremail格納エラー")
	}
	if user.password == "password" {
	} else {
		t.Error("userpass格納エラー")
	}

	//passwordをハッシュ化する
	hashpass1 := user.md5hash("password")
	if len(hashpass1) == 32 {
	} else {
		t.Log(user.md5hash("password"))
		t.Error("パスワードハッシュ化エラー")
	}
	hashpass2 := user.md5hash("password2")
	if hashpass1 != hashpass2 {
	} else {
		t.Error("正常にハッシュ化されていません。")
	}

	//DB接続を行う
	user.dbConnect(func(db *sql.DB) {
		err := db.Ping()
		if err == nil {
		} else {
			t.Error("DB接続エラー")
		}
	})
}
