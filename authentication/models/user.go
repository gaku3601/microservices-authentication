package models

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"

	_ "github.com/lib/pq"
)

type (
	User interface{}
	user struct {
		email    string `json:email`
		password string `json:password`
	}
)

func NewUser(email string, password string) *user {
	return &user{email: email, password: password}
}

func (u *user) md5hash(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (u *user) InsertUser() error {
	var err = errors.New("")
	u.dbConnect(func(db *sql.DB) {
		_, err = db.Exec("INSERT INTO Users(email, password) VALUES($1, $2);", u.email, u.md5hash(u.password))
	})
	if err != nil {
		if err.Error() == `pq: 重複キーが一意性制約"users_email_key"に違反しています` {
			return errors.New("email重複エラー")
		}
	}

	return err
}

func (u *user) FetchUser() (int, error) {
	var id = 0
	var err = errors.New("")
	u.dbConnect(func(db *sql.DB) {
		err = db.QueryRow("SELECT ID FROM USERS WHERE EMAIL = $1 AND PASSWORD = $2;", u.email, u.md5hash(u.password)).Scan(&id)
	})
	return id, err
}

func (u *user) dbConnect(fn func(db *sql.DB)) {
	db, _ := sql.Open("postgres", "user=postgres host=localhost dbname=auth_db port=5432 sslmode=disable")
	fn(db)
}
