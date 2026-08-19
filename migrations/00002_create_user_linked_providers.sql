-- +goose Up
CREATE TABLE user_linked_providers (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    provider TEXT NOT NULL
        CHECK (provider IN ('google', 'email_password', 'apple', 'vk', 'github')),
    PRIMARY KEY (user_id, provider)
);

-- +goose Down
DROP TABLE user_linked_providers;
