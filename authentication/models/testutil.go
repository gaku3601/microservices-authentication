package models

import "database/sql"

func Setup() {
	//userを作成
	user := NewUser("pro.gaku@gmail.com", "password")
	user.dbConnect(func(db *sql.DB) {
		db.Exec("INSERT INTO Users(id, email, password) VALUES(1, $1, $2);", user.email, user.md5hash(user.password))
	})
}

func Teardown() {
	user := NewUser("", "")
	user.dbConnect(func(db *sql.DB) {
		db.Exec("DELETE FROM USERS;")
	})
}
