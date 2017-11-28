package main

import (
	"net/http"

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
	http.Error(w, "エラー", 1)
}

func signup(w http.ResponseWriter, r *http.Request) {
}

func readConfig(path string) {
	config.SetConfig(path)
}
