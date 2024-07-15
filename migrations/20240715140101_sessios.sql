-- +goose Up
-- +goose StatementBegin
CREATE TABLE sessions (
                          id UUID PRIMARY KEY,
                          seller_id UUID NOT NULL REFERENCES sellers(id),
                          created_at TIMESTAMPTZ NOT NULL,
                          expired_at TIMESTAMPTZ NOT NULL
);;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sessions;
-- +goose StatementEnd
