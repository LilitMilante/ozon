-- +goose Up
-- +goose StatementBegin
CREATE TABLE sellers
(
    id UUID PRIMARY KEY,
    full_name TEXT NOT NULL,
    login TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sellers;
-- +goose StatementEnd
