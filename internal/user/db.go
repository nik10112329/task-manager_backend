package user

import (
	"context"
	"fmt"
	"learning-project/pkg/logging"
	"github.com/jackc/pgx/v5/pgxpool"
)


type db struct {
	pool   *pgxpool.Pool
	logger *logging.Logger
}

func NewStorage(pool *pgxpool.Pool, logger *logging.Logger) Storage {
	return &db{pool: pool, logger: logger}
}

func (d *db) FindOne(ctx context.Context, id string) (UserEntity, error) {
	var u UserEntity
	err := d.pool.QueryRow(ctx, `
        SELECT id, email, display_name, sign_in_method, id_linked_provider,
               photo_url, phone_number, email_verified, created_at, last_login_at
        FROM users WHERE id = $1
    `, id).Scan(&u.ID, &u.Email, &u.DisplayName, &u.SignInMethod, &u.IdLinkedProvider,
		&u.PhotoURL, &u.PhoneNumber, &u.EmailVerify, &u.CreatedAt, &u.LastLoginAt)
	if err != nil {
		// pgx.ErrNoRows → это не "ошибка сервера", а "пользователь не найден" —
		// на уровне service это два разных исхода, не заворачивайте их в один и тот же err
		return UserEntity{}, fmt.Errorf("find user by id: %w", err)
	}
	return u, nil
}

func (d *db) Create(ctx context.Context, user UserEntity) (string, error) {
	var u UserEntity
	err := d.pool.QueryRow(ctx, `
		INSERT INTO users (id, email, display_name, sign_in_method, id_linked_provider,
			photo_url, phone_number, email_verified, created_at, last_login_at) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`, user.ID, user.Email, user.DisplayName, user.SignInMethod, user.IdLinkedProvider,
		user.PhotoURL, user.PhoneNumber, user.EmailVerify, user.CreatedAt, user.LastLoginAt).Scan(&u.ID,
		&u.Email, &u.DisplayName, &u.SignInMethod, &u.IdLinkedProvider,
		&u.PhotoURL, &u.PhoneNumber, &u.EmailVerify, &u.CreatedAt, &u.LastLoginAt)
	if err != nil {
		return "", fmt.Errorf("create user: %w", err)
	}
	return u.ID, nil
}

func (d *db) Update(ctx context.Context, user UserEntity) error {
	_, err := d.pool.Exec(ctx, `
		UPDATE users SET email = $2, display_name = $3, sign_in_method = $4, 
		id_linked_provider = $5, photo_url = $6, phone_number = $7, email_verified = $8, 
		last_login_at = $9
		WHERE id = $1
	`, user.ID, user.Email, user.DisplayName, user.SignInMethod, user.IdLinkedProvider,
		user.PhotoURL, user.PhoneNumber, user.EmailVerify, user.LastLoginAt)
	if err != nil {
		return fmt.Errorf("update user: %w", err)
	}
	return nil
}
func (d *db) Delete(ctx context.Context, id string) error {
	_, err := d.pool.Exec(ctx, `
		DELETE FROM users WHERE id = $1
	`, id)

	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}
	return nil
}