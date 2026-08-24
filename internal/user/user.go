package user

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin/render"
	"github.com/google/uuid"
)

type UserEntity struct {
	ID          string
	DisplayName string `json:"display_name"`
	// Password		string `json:"-"`
	Email            string    `json:"email"`
	PhotoURL         string    `json:"photo_url"`
	PhoneNumber      string    `json:"phone_number"`
	EmailVerify      bool      `json:"email_verify"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	IdLinkedProvider []string  `json:"id_linked_provider"`
	SignInMethod     string    `json:"sign_in_method"`
	LastLoginAt      time.Time `json:"last_login_at"`
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
	// jwtToken, jwtError := newJwtToken()
	// if jwtError != nil {
	// 	return UserEntity{}, jwtError
	// }
	return UserEntity{
		newUserId.String(),
		jsonData["firstName"].(string) + " " + jsonData["lastName"].(string),
		// jsonData["password"].(string),
		jsonData["email"].(string),
		jsonData["photoUrl"].(string),
		jsonData["phoneNumber"].(string),
		false,
		time.Now(),
		time.Now(),
		[]string{jsonData["linkedProviders"].([]interface{})[0].(string), "originalBackend"},
		"originalBackend",
		time.Now(),
	}, nil
}

func CreateUserEntityFromFirebase(json render.JSON) (UserEntity, error) {
	jsonData := json.Data.(map[string]interface{})
	if jsonData == nil {
		return UserEntity{}, errors.New("invalid json data")
	}
	// jwtToken, jwtError := newJwtToken()
	// if jwtError != nil {
	// 	return UserEntity{}, jwtError
	// }
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
		time.Now(),
	}, nil
}
