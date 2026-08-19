package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/google/uuid"
)

type JwtToken struct {
	Id        string
	Platform  string
	Token     string
	UpdatedAt time.Time
}

func newJwtToken() (JwtToken, error) {
	tokenId, idError := uuid.NewV7()
	if idError != nil {
		return JwtToken{}, idError
	}
	token, tokenError := jwt.New(jwt.SigningMethodRS512).SigningString()
	if tokenError != nil {
		return JwtToken{}, tokenError
	}
	return JwtToken{
		Id:        tokenId.String(),
		Platform:  "originalBackend",
		Token:     token,
		UpdatedAt: time.Now(),
	}, nil
}
func updateJwt(oldToken JwtToken) (JwtToken, error) {

	token, tokenError := jwt.New(jwt.SigningMethodRS512).SigningString()
	if tokenError != nil {
		return JwtToken{}, tokenError
	}
	return JwtToken{
		Id:        oldToken.Id,
		Platform:  oldToken.Platform,
		Token:     token,
		UpdatedAt: time.Now(),
	}, nil
}
