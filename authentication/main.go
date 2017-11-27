package main

import (
	"net/http"

	"github.com/gaku3601/study-microservices/authentication/config"
	"github.com/gaku3601/study-microservices/authentication/login"
	"github.com/gaku3601/study-microservices/authentication/signup"
	_ "github.com/lib/pq"

	mux "github.com/gorilla/mux.git"
)

func main() {
	readConfig("./config")

	r := mux.NewRouter()
	r.HandleFunc("/users/login", login.Login).Methods("POST")
	r.HandleFunc("/users/signup", signup.SignUp).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func readConfig(path string) {
	config.SetConfig(path)
}
