package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gaku3601/microservices-authentication/authentication/models"
	"github.com/gaku3601/study-microservices/authentication/config"
	_ "github.com/lib/pq"

	mux "github.com/gorilla/mux.git"
)

func main() {
	readConfig("./config")

	r := mux.NewRouter()
	r.HandleFunc("/users/login", login).Methods("POST")
	r.HandleFunc("/users/signup", signup).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func login(w http.ResponseWriter, r *http.Request) {
	data := receptionData(r)
	user := models.NewUser(data["email"].(string), data["password"].(string))
	id, _ := user.FetchUser()

	if id == 0 {
		//送信処理
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"token":""}`)
		return
	}

	jwt := models.NewJwt("http://localhost:8001/consumers/gaku/jwt")
	token := jwt.CreateToken(models.Claims{
		"id": id,
	})

	//送信処理
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"token":"%s"}`, token)
}

func signup(w http.ResponseWriter, r *http.Request) {
}

//受信データ受け取り処理
func receptionData(r *http.Request) map[string]interface{} {
	body, _ := ioutil.ReadAll(r.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	return data.(map[string]interface{})
}

func readConfig(path string) {
	config.SetConfig(path)
}
