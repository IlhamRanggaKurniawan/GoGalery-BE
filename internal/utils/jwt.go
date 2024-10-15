package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var REFRESH_TOKEN_SECRET = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
var ACCESS_TOKEN_SECRET = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))

type Claims struct {
	Id         uint64 `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	ProfileUrl string `json:"profileUrl"`
	Bio        string `json:"bio"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(username string, email string, id uint64, role string, profilePicture *string, userBio *string) (string, error) {
	Exp := time.Now().Add(24 * time.Hour * 7)

	claims := &Claims{
		Username:   username,
		Email:      email,
		Id:         id,
		Role:       role,
		ProfileUrl: "",
		Bio:        "",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(Exp),
		},
	}

	if profilePicture != nil {
		claims.ProfileUrl = *profilePicture
	}

	if userBio != nil {
		claims.Bio = *userBio
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessTokenSTR, err := accessToken.SignedString(ACCESS_TOKEN_SECRET)

	if err != nil {
		return "", err
	}

	return accessTokenSTR, nil
}
func GenerateRefreshToken(username string, email string, id uint64, role string, profilePicture *string, userBio *string) (string, error) {

	Exp := time.Now().Add(24 * time.Hour * 7)

	claims := &Claims{
		Username:   username,
		Email:      email,
		Id:         id,
		Role:       role,
		ProfileUrl: "",
		Bio:        "",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(Exp),
		},
	}

	if profilePicture != nil {
		claims.ProfileUrl = *profilePicture
	}

	if userBio != nil {
		claims.Bio = *userBio
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshTokenSTR, err := refreshToken.SignedString(REFRESH_TOKEN_SECRET)

	return refreshTokenSTR, err
}

func ValidateToken(tokenString string, tokenType string) (*Claims, error) {
	claims := &Claims{}
	var secretKey []byte

	if tokenType == "Access Token" {
		secretKey = ACCESS_TOKEN_SECRET
	} else if tokenType == "Refresh Token" {
		secretKey = REFRESH_TOKEN_SECRET
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	return claims, nil
}

func DecodeAccessToken(r *http.Request) (*Claims, error) {

	cookie, err := r.Cookie("AccessToken")
    if err != nil {
        return nil, fmt.Errorf("could not find access token in cookies: %v", err)
    }

    token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return ACCESS_TOKEN_SECRET, nil
    })

    if err != nil || !token.Valid {
        return nil, fmt.Errorf("invalid token: %v", err)
    }

    claims, ok := token.Claims.(*Claims)
    if !ok {
        return nil, fmt.Errorf("failed to parse claims")
    }

    return claims, nil
}
