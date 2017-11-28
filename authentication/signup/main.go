package signup

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

//ユーザ登録
func SignUp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	user := new(User)

	decoder.Decode(&user)
	if user.Email == "" || user.Password == "" {
		errorResponse := new(ErrorResponse)
		errorResponse.Error = "ユーザまたはパスワードの送信エラーです。"
		res, _ := json.Marshal(errorResponse)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
		return
	}

	//DB登録
	/*
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
	*/
}
