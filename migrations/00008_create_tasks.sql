-- +goose Up
CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    name VARCHAR NOT NULL,
    status TEXT 
        CHECK(status IN ('todo', 'in_progress', 'on_review', 
        'done', 'blocked', 'cancelled'))
        NOT NULL DEFAULT 'todo',
    owner_id UUID REFERENCES users(id) ON DELETE SET NULL,
    archived BOOLEAN NOT NULL DEFAULT FALSE,
    archived_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    start_date TIMESTAMPTZ,
    end_date TIMESTAMPTZ,
    description TEXT
);
-- +goose Down
DROP TABLE tasks;