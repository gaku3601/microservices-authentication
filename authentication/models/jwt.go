package models

import (
	"encoding/json"
	"net/http"
)

type Jwt struct {
	Created_at  int    `json:"created_at"`
	Id          string `json:"id"`
	Algorithm   string `json:"algorithm"`
	Key         string `json:"key"`
	Secret      string `json:"secret"`
	Consumer_id string `json:"consumer_id"`
}

func NewJwt() *Jwt {
	j := new(Jwt)
	return j
}

func (j *Jwt) fetchJwtKey(url string) {
	req, _ := http.NewRequest(
		"POST",
		url,
		nil,
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(j)
}
