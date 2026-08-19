package models

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin/render"
	"github.com/google/uuid"
)

type UserEntity struct {
	ID              string
	DisplayName     string
	Email           string
	PhotoURL        string
	PhoneNumber     string
	EmailVerify     bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	LinkedProviders []string
	SignInMethod    string
	JwtToken        JwtToken
}

func NewUser() UserEntity {
	return UserEntity{}
}

func CreateUserEntityFromJson(json render.JSON) (UserEntity, error) {
	jsonData := json.Data.(map[string]interface{})
	newUserId, errorIdGenerate := uuid.NewV7()
	if errorIdGenerate != nil {
		return UserEntity{}, errorIdGenerate
	}
	jwtToken, jwtError := newJwtToken()
	if jwtError != nil {
		return UserEntity{}, jwtError
	}
	return UserEntity{
		newUserId.String(),
		jsonData["firstName"].(string) + " " + jsonData["lastName"].(string),
		jsonData["email"].(string),
		jsonData["photoUrl"].(string),
		jsonData["phoneNumber"].(string),
		false,
		time.Now(),
		time.Now(),
		[]string{jsonData["linkedProviders"].([]interface{})[0].(string), "originalBackend"},
		"originalBackend",
		jwtToken,
	}, nil
}

func CreateUserEntityFromFirebase(json render.JSON) (UserEntity, error) {
	jsonData := json.Data.(map[string]interface{})
	if jsonData == nil {
		return UserEntity{}, errors.New("invalid json data")
	}
	jwtToken, jwtError := newJwtToken()
	if jwtError != nil {
		return UserEntity{}, jwtError
	}
	return UserEntity{
		jsonData["id"].(string),
		jsonData["displayName"].(string),
		jsonData["email"].(string),
		jsonData["photoUrl"].(string),
		jsonData["phoneNumber"].(string),
		jsonData["emailVerify"].(bool),
		time.Now(),
		time.Now(),
		[]string{jsonData["linkedProviders"].([]interface{})[0].(string), "originalBackend"},
		jsonData["signInMethod"].(string),
		jwtToken,
	}, nil
}
