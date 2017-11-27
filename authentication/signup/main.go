package signup

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gaku3601/study-microservices/authentication/dbc"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//ユーザ登録
func SignUp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	user := new(User)

	error := decoder.Decode(&user)
	if error != nil {
		w.Write([]byte("json decode error " + error.Error() + "\n"))
	}

	//DB登録
	dbc.DBConnect(func(db *sql.DB) {
		//passwordのhash化
		bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		_, err := db.Exec("INSERT INTO users(email, password) VALUES($1, $2);", user.Email, bcryptPassword)
		if err != nil {
			w.Write([]byte("Signup DB insert error: " + err.Error() + "\n"))
		} else {
			w.Write([]byte("Signup OK\n"))
		}
	})
}
