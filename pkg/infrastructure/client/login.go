package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gotodo/pkg/infrastructure/config"
	jwt2 "gotodo/pkg/infrastructure/jwt"
	"net/http"
)

type Login struct {
	url        string
	saveConfig config.SaveConfig
}

func NewLoginClient(
	baseUrl string,
	saveConfig config.SaveConfig,
) Login {
	return Login{
		url:        fmt.Sprintf("%s/login", baseUrl),
		saveConfig: saveConfig,
	}
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l Login) Execute(credentials Credentials) (*jwt2.TokenClaims, error) {
	payload, err := json.Marshal(credentials)
	if err != nil {
		return nil, err
	}

	var apiResponse map[string]string
	err = DoRequest(Request{
		Method:      http.MethodPost,
		Url:         l.url,
		RequestBody: bytes.NewBuffer(payload),
		ApiResponse: &apiResponse,
	})
	if err != nil {
		return nil, err
	}

	tokenString := apiResponse["token"]
	claims := &jwt2.TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, nil)
	claims, _ = token.Claims.(*jwt2.TokenClaims)

	l.saveConfig(config.Config{
		Token:    tokenString,
		UserName: claims.Name,
	})

	return claims, nil
}
