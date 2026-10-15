package user

import (
	"context"
)

type Storage interface {
	Create(ctx context.Context, user UserEntity) (string, error)
	FindOneById(ctx context.Context, id string) (UserEntity, error)
	FindOneByParams(ctx context.Context,id string, email string, phoneNumber string) (UserEntity, error)
	Update(ctx context.Context, user UserEntity) error
	Delete(ctx context.Context, id string) error
}
