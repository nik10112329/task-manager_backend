-- +goose Up
CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    is_read BOOLEAN NOT NULL DEFAULT FALSE,
    date TIMESTAMPTZ NOT NULL,
    text TEXT NOT NULL,
    title VARCHAR NOT NULL,
    image_url VARCHAR,
    deep_link VARCHAR
);

-- +goose Down
DROP TABLE notifications;
