package user

import (
	"context"
)

type Storage interface {
	Create(ctx context.Context, user UserEntity) (string, error)
	FindOne(ctx context.Context, id string) (UserEntity, error)
	Update(ctx context.Context, user UserEntity) error
	Delete(ctx context.Context, id string) error
}
