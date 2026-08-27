// internal/user/service.go
package user

import (
	"context"
	"fmt"
)

var allowedSignInMethods = map[string]bool{
	"google": true, "email_password": true, "apple": true, "vk": true, "github": true,
}

type Service interface {
	Register(ctx context.Context, req CreateUserRequest) (UserEntity, error)
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) Register(ctx context.Context, req CreateUserRequest) (UserEntity, error) {
	if req.Email == "" {
		return UserEntity{}, fmt.Errorf("email is required")
	}
	if !allowedSignInMethods[req.SignInMethod] {
		return UserEntity{}, fmt.Errorf("unsupported sign_in_method: %s", req.SignInMethod)
	}

	user := UserEntity{
		Email:            req.Email,
		DisplayName:      req.DisplayName,
		PhotoURL:         req.PhotoURL,
		PhoneNumber:      req.PhoneNumber,
		SignInMethod:     req.SignInMethod,
		IDLinkedProvider: req.IDLinkedProvider,
		LinkedProviders:  req.LinkedProviders,
	}

	id, err := s.storage.Create(ctx, user)
	if err != nil {
		return UserEntity{}, fmt.Errorf("register user: %w", err)
	}

	return s.storage.FindOne(ctx, id)
}
