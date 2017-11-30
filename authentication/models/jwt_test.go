package models

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewJwt(t *testing.T) {
	jwt := NewJwt()
	if jwt == nil {
		t.Error("jwtモデルの生成エラー")
	}
}

//JwtKeyを取得する
func TestFetchJwtKey(t *testing.T) {
	ts := jwtServerStub()
	defer ts.Close()

	jwt := NewJwt()
	jwt.fetchJwtKey(ts.URL)

	if jwt.Key != "RmzcPktBjNbnsGdZPwLioOmdThCjFGIO" {
		t.Log(jwt)
		t.Error("JwtKey取得エラー")
	}
}

func jwtServerStub() *httptest.Server {
	//jwt key発行用Stub(kongで出力想定)
	tokenKeyStub := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
			{
				"created_at":1510615420000,
				"id":"79051f7f-ff24-4d69-8e1d-d83146bc9ec7",
				"algorithm":"HS256",
				"key":"RmzcPktBjNbnsGdZPwLioOmdThCjFGIO",
				"secret":"wKtru3BuCiT9vFFki77cg5DE2rt6a4if",
				"consumer_id":"0a086d40-dafd-43e2-94dc-835d1b96c92b"
			}
		`)
	})
	return httptest.NewServer(tokenKeyStub)
}
