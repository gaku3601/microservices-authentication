package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	//パラメータの取得テスト
	stub := http.HandlerFunc(login)
	ts := httptest.NewServer(stub)
	defer ts.Close()

	jsonStr := `{"email":"test","password":"aaaa"}`

	req, _ := http.NewRequest(
		"POST",
		ts.URL,
		bytes.NewBuffer([]byte(jsonStr)),
	)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
	t.Error(string(body))

	/*
		if data.(map[string]interface{})["email"] != nil {
		} else {
			t.Error("emailの取得エラー")
		}
	*/
}
