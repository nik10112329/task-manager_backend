package db

import (
	"context"
	"database/sql"
	"learning-project/internal/user"
	"learning-project/pkg/logging"
)

type db struct {
	dbTable *sql.DB
	logger  *logging.Logger
}

func (d *db) Create(ctx context.Context, user user.UserEntity) (string, error) {

	panic("unimplemented")
}

func (d *db) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (d *db) FindOne(ctx context.Context, id string) (user.UserEntity, error) {
	panic("unimplemented")
}

func (d *db) Update(ctx context.Context, user user.UserEntity) error {
	panic("unimplemented")
}

func NewStorage(collection string, logger *logging.Logger) user.Storage {
	return &db{}
}
