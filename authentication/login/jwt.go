package login

import (
	"encoding/json"
	"errors"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtKey struct {
	Created_at  int    `json:"created_at"`
	Id          string `json:"id"`
	Algorithm   string `json:"algorithm"`
	Key         string `json:"key"`
	Secret      string `json:"secret"`
	Consumer_id string `json:"consumer_id"`
	UserID      string
}

func fetchCreateToken(userID string, url string) (string, error) {
	req, _ := http.NewRequest(
		"POST",
		url,
		nil,
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	jwtKey := new(JwtKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("Jwt Key発行用のサーバが指定されていません。")
	}
	json.NewDecoder(resp.Body).Decode(&jwtKey)

	jwtKey.UserID = userID
	defer resp.Body.Close()

	return createTokenString(jwtKey), nil
}

func createTokenString(jwtKey *JwtKey) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  jwtKey.UserID,
		"iss": jwtKey.Key,
	})
	tokenString, _ := token.SignedString([]byte(jwtKey.Secret))

	return tokenString
}
