package models

import (
	"crypto/md5"
	"encoding/hex"
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
