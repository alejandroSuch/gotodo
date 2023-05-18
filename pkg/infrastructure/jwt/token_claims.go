package jwt

import "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	UserID string `json:"userId"`
	Name   string `json:"name"`
	jwt.StandardClaims
}
