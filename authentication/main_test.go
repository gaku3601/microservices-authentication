package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gaku3601/microservices-authentication/authentication/models"
)

//ログインテスト
func TestLogin(t *testing.T) {
	models.Teardown()
	models.Setup()
	//パラメータの取得テスト
	stub := http.HandlerFunc(login)
	ts := httptest.NewServer(stub)
	defer ts.Close()

	data := loginRequest(ts.URL).(map[string]interface{})
	claims := models.DecryptionToken(data["token"].(string))
	if (*claims)["id"] != float64(1) {
		t.Error("loginエラー")
	}

	errordata := errorLoginRequest(ts.URL).(map[string]interface{})
	t.Log(errordata)
	if errordata["token"] != "" {
		t.Error("loginエラー")
	}

}

func TestSignup(t *testing.T) {
	models.Teardown()
	//パラメータの取得テスト
	stub := http.HandlerFunc(signup)
	ts := httptest.NewServer(stub)
	defer ts.Close()

	data := signupRequest(ts.URL).(map[string]interface{})

	if data["status"] != "ok" {
		t.Error("signup登録エラー")
	}

	//重複パターン
	data = signupRequest(ts.URL).(map[string]interface{})

	if data["status"] != "error" {
		t.Error("エラー時チェック")
	}

}

//json受け取りを実施し、正常にmap[string]interface{}で格納されるかテスト
func TestReceptionData(t *testing.T) {
	stub := http.HandlerFunc(login)
	ts := httptest.NewServer(stub)
	defer ts.Close()

	jsonStr := `{"email":"pro.gaku@gmail.com","password":"password"}`

	req, _ := http.NewRequest(
		"POST",
		ts.URL,
		bytes.NewBuffer([]byte(jsonStr)),
	)

	req.Header.Set("Content-Type", "application/json")
	data := receptionData(req)
	if data["email"] != "pro.gaku@gmail.com" {
		t.Error("request受け取りエラー")
	}
}

func signupRequest(url string) interface{} {
	jsonStr := `{"email":"pro.gaku@gmail.com","password":"password"}`

	req, _ := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var data interface{}
	json.Unmarshal(body, &data)

	return data
}

//ログインリクエストを投げ、帰ってきたデータをintefaceで返却する。
func loginRequest(url string) interface{} {
	jsonStr := `{"email":"pro.gaku@gmail.com","password":"password"}`

	req, _ := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var data interface{}
	json.Unmarshal(body, &data)

	return data
}

func errorLoginRequest(url string) interface{} {
	jsonStr := `{"email":"pro.gaku@gmail.com2","password":"password"}`

	req, _ := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var data interface{}
	json.Unmarshal(body, &data)

	return data
}
