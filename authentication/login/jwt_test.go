package login

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchCreateToken(t *testing.T) {
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
	ts := httptest.NewServer(tokenKeyStub)
	defer ts.Close()
	token, _ := fetchCreateToken("1", ts.URL)

	if token != "" {
	} else {
		t.Error("fetchCreateToken:token発行に失敗しています。")
	}

	//jwt key発行サーバが存在しない場合のテスト
	_, err := fetchCreateToken("1", "")

	if err != nil {
	} else {
		t.Errorf("jwt key発行サーバが存在しない場合のテスト:%s", err)
	}
}
