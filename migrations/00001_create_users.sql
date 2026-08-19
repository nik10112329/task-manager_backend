-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    display_name VARCHAR,
    email VARCHAR NOT NULL UNIQUE,
    sign_in_method TEXT NOT NULL  
        CHECK (sign_in_method IN ('google', 'email_password', 'apple', 'vk', 'github')),
    photo_url VARCHAR,
    phone_number VARCHAR,
    email_verified BOOLEAN NOT NULL default FALSE,
    id_linked_provider VARCHAR,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    last_login_at TIMESTAMPTZ
);
-- +goose Down
DROP TABLE users;
