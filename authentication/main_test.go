package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//ログインテスト
func TestLogin(t *testing.T) {
	//パラメータの取得テスト
	stub := http.HandlerFunc(login)
	ts := httptest.NewServer(stub)
	defer ts.Close()

	t.Log(loginRequest(ts.URL))
	data := loginRequest(ts.URL).(map[string]interface{})
	if data["email"] != "pro.gaku@gmail.com" {
		t.Error("loginエラー")
	}
}

//json受け取りを実施し、正常にmap[string]interface{}で格納されるかテスト
func TestReceptionData(t *testing.T) {
	stub := http.HandlerFunc(login)
	ts := httptest.NewServer(stub)
	defer ts.Close()

	jsonStr := `{"email":"pro.gaku@gmail.com","password":"test"}`

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

//ログインリクエストを投げ、帰ってきたデータをintefaceで返却する。
func loginRequest(url string) interface{} {
	jsonStr := `{"email":"pro.gaku@gmail.com","password":"test"}`

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
