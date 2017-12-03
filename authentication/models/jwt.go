package models

import (
	"encoding/json"
	"net/http"

	jwtgo "github.com/dgrijalva/jwt-go"
)

type Jwt struct {
	Created_at  int    `json:"created_at"`
	Id          string `json:"id"`
	Algorithm   string `json:"algorithm"`
	Key         string `json:"key"`
	Secret      string `json:"secret"`
	Consumer_id string `json:"consumer_id"`
}

func NewJwt(url string) *Jwt {
	j := new(Jwt)
	j.fetchJwtKey(url)
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

type Claims map[string]interface{}

func (c Claims) Valid() error {
	return nil
}

func (j *Jwt) CreateToken(claims Claims) string {
	//issはkong認証で必ず必要であるため、必須で追加する
	claims["iss"] = j.Key
	tokendata := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	token, _ := tokendata.SignedString([]byte(j.Secret))

	return token
}

func DecryptionToken(token string) *jwtgo.MapClaims {
	p, _ := jwtgo.Parse(token, nil)
	data := p.Claims.(jwtgo.MapClaims)
	return &data
}
