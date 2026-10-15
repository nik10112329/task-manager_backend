package user

import (
	"time"
)

type UserEntity struct {
	ID               string     `json:"id"`
	DisplayName      string     `json:"display_name"`
	Email            string     `json:"email"`
	PhotoURL         string     `json:"photo_url"`
	PhoneNumber      string     `json:"phone_number"`
	EmailVerify      bool       `json:"email_verify"`
	CreatedAt        time.Time  `json:"created_at"`
	IDLinkedProvider string     `json:"id_linked_provider"`
	LinkedProviders  []string   `json:"linked_providers"`
	SignInMethod     string     `json:"sign_in_method"`
	LastLoginAt      *time.Time `json:"last_login_at,omitempty"`
}

func NewUser() UserEntity {
	return UserEntity{}
}
