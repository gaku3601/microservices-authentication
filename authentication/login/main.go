package login

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gaku3601/study-microservices/authentication/dbc"
	viper "github.com/spf13/viper.git"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FetchUser struct {
	ID string `json:"id"`
	User
}

type Response struct {
	Token string `json:"token"`
}

//ログイン認証。ログイン完了後、JWTトークンを返却する
func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	error := decoder.Decode(&user)
	if error != nil {
		w.Write([]byte("json decode error " + error.Error() + "\n"))
	}

	dbc.DBConnect(func(db *sql.DB) {
		fetchUser := new(FetchUser)
		err := db.QueryRow("SELECT id,email,password FROM users where email = $1;", user.Email).Scan(&fetchUser.ID, &fetchUser.Email, &fetchUser.Password)
		if err != nil {
			w.Write([]byte("emailが登録されていません。:" + err.Error() + "\n"))
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(fetchUser.Password), []byte(user.Password))
		if err != nil {
			w.Write([]byte("email,passwordが違います。:" + err.Error() + "\n"))
			return
		}

		//認証処理
		res := new(Response)
		res.Token, err = fetchCreateToken(fetchUser.ID, viper.GetString("TokenKeyServerURL"))
		if err != nil {
			fmt.Println("jwt認証サーバが検出できませんでした。")
			os.Exit(1)
		}
		//返却
		json.NewEncoder(w).Encode(res)
	})
}
