package models

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"

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

func (u *user) dbConnect(fn func(db *sql.DB)) {
	db, _ := sql.Open("postgres", "user=postgres host=localhost dbname=auth_db port=5432 sslmode=disable")
	fn(db)
}
