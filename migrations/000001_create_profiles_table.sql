-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS profiles
(
    id   SERIAL PRIMARY KEY NOT NULL,
    created_at timestamp(0) NOT NULL DEFAULT now(),
    updated_at timestamp(0) NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS profiles;
-- +goose StatementEnd
