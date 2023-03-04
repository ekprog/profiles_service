-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS profiles
(
    user_id    bigint unique not null,
    name       VARCHAR(255)  NOT NULL,
    photo      VARCHAR(255)  NOT NULL,
    created_at timestamp(0)  NOT NULL DEFAULT now(),
    updated_at timestamp(0)  NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS profiles;
-- +goose StatementEnd
