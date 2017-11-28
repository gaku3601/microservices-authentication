package signup

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignUp(t *testing.T) {
	var signup = http.HandlerFunc(SignUp)
	ts := httptest.NewServer(signup)
	defer ts.Close()

	jsonStr := `{"emaila":"test","password":"aaaa"}`

	req, _ := http.NewRequest(
		"POST",
		ts.URL,
		bytes.NewBuffer([]byte(jsonStr)),
	)

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	var data interface{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &data)
	if data.(map[string]interface{}) != nil {
	} else {
		t.Error("デコードエラーです。")
	}
}
