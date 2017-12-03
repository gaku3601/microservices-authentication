package models

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestNewJwt(t *testing.T) {
	ts := jwtServerStub()
	defer ts.Close()
	jwt := NewJwt(ts.URL)
	if jwt.Created_at != 1510615420000 {
		t.Error("jwtモデルの生成エラー")
	}
}

func TestFetchJwtKey(t *testing.T) {
	ts := jwtServerStub()
	defer ts.Close()

	jwt := new(Jwt)
	jwt.fetchJwtKey(ts.URL)

	if jwt.Key != "RmzcPktBjNbnsGdZPwLioOmdThCjFGIO" {
		t.Log(jwt)
		t.Error("JwtKey取得エラー")
	}
}

func TestCreateToken(t *testing.T) {
	ts := jwtServerStub()
	defer ts.Close()

	jwt := NewJwt(ts.URL)

	token := jwt.CreateToken(claimsTestData())
	t.Log(token)
	//jwtの場合、確実にドットが二つ入っているはずなので、それを利用しテストする。
	arr := strings.Split(token, ".")
	if len(arr) != 3 {
		t.Error("token生成エラー")
	}
}

func TestDecryptionToken(t *testing.T) {
	ts := jwtServerStub()
	defer ts.Close()

	jwt := NewJwt(ts.URL)
	token := jwt.CreateToken(claimsTestData())

	data := DecryptionToken(token)

	t.Log((*data)["id"])
	t.Log(reflect.TypeOf((*data)["id"]))
	if (*data)["id"] != float64(1) {
		t.Error("jwtデコードエラー")
	}
	if (*data)["name"] != "gaku" {
		t.Error("jwtデコードエラー")
	}
}

func TestValid(t *testing.T) {
	//Valid() errorはjwt-goのClaimsインターフェースを利用するために作成したもの。
	//そのため、何もせず、nilを返却する。
	claims := new(Claims)
	if claims.Valid() != nil {
		t.Error("Valid()エラー")
	}
}

func claimsTestData() Claims {
	return Claims{
		"id":   1,
		"name": "gaku",
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
