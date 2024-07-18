package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var REFRESH_TOKEN_SECRET = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
var ACCESS_TOKEN_SECRET = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ID       uint64 `json:"id"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(username string, email string, id uint64, role string) (string, error) {
	Exp := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Username: username,
		Email:    email,
		ID:       id,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(Exp),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessTokenSTR, err := accessToken.SignedString(ACCESS_TOKEN_SECRET)

	if err != nil {
		return "", err
	}

	return accessTokenSTR, nil
}
func GenerateRefreshToken(username string, email string, id uint64, role string) (string, error) {

	Exp := time.Now().Add(24 * time.Hour * 7)

	claims := &Claims{
		Username: username,
		Email:    email,
		ID:       id,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(Exp),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshTokenSTR, err := refreshToken.SignedString(REFRESH_TOKEN_SECRET)

	return refreshTokenSTR, err
}

func ValidateToken(tokenString string, tokenType string) (*Claims, error) {
	claims := &Claims{}

	if tokenType == "Access Token" {
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return ACCESS_TOKEN_SECRET, nil
		})

		if err != nil {
			return nil, err
		}

		if !token.Valid {
			return nil, errors.New("invalid token")
		}

		if  claims.ExpiresAt.Before(time.Now()) {
			return nil, errors.New("token expired")
		}

		return claims, nil
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return REFRESH_TOKEN_SECRET, nil
	})

	  if err != nil || !token.Valid {
        return nil, err
    }

	return claims, nil
}
