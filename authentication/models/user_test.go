package models

import (
	"database/sql"
	"testing"
)

//emailとpasswordがDBにUser登録されていることを確認する。
func TestNewUser(t *testing.T) {
	//userにemailとpasswordを渡し、格納されていることを確認する。
	user := NewUser("pro.gaku@gmail.com", "password")
	if user.email != "pro.gaku@gmail.com" {
		t.Error("useremail格納エラー")
	}

	if user.password != "password" {
		t.Error("userpass格納エラー")
	}

}

func TestInsertUser(t *testing.T) {
	Teardown()
	//insert処理
	user := NewUser("pro.gaku@gmail.com", "passpass")
	user.InsertUser()
	//fetch処理
	email, password := fetchUser(user)
	t.Log(email)
	if email != "pro.gaku@gmail.com" {
		t.Error("ユーザ登録エラー:email")
	}
	if password != user.md5hash("passpass") {
		t.Error("ユーザ登録エラー:password")
	}

	//emailが重複し、登録する場合エラーを返却する。
	err := user.InsertUser()
	t.Log(err)

	if err.Error() != "email重複エラー" {
		t.Error("email重複エラーチェック")
	}
}

func TestMD5hash(t *testing.T) {
	//passwordをハッシュ化する
	user := NewUser("pro.gaku@gmail.com", "password")
	hashdata := user.md5hash("password")

	if len(hashdata) != 32 {
		t.Log(user.md5hash("password"))
		t.Error("パスワードハッシュ化エラー")
	}
	if user.md5hash("password2") == hashdata {
		t.Error("正常にハッシュ化されていません。")
	}
}

func TestDBConnect(t *testing.T) {
	//DB接続を行う
	user := NewUser("pro.gaku@gmail.com", "password")
	user.dbConnect(func(db *sql.DB) {
		if db.Ping() != nil {
			t.Error("DB接続エラー")
		}
	})
}

func TestFetchUser(t *testing.T) {
	Teardown()
	Setup()
	//DBからUser情報を取得する。
	user := NewUser("pro.gaku@gmail.com", "password")
	userID, err := user.FetchUser()
	if userID != 1 {
		t.Errorf("DBよりUserを取得する際に不具合があります。:%s", err)
	}

	user = NewUser("pro.gaku@gmail.com2", "password")
	userID, err = user.FetchUser()
	if err == nil {
		t.Errorf("FetchUser()でエラーを返却できていません。%s", err)
	}
}

func fetchUser(u *user) (string, string) {
	email := ""
	password := ""
	u.dbConnect(func(db *sql.DB) {
		db.QueryRow("SELECT EMAIL,PASSWORD FROM USERS WHERE EMAIL = $1 AND PASSWORD = $2;", u.email, u.md5hash(u.password)).Scan(&email, &password)
	})
	return email, password
}
