-- +goose Up
    ALTER TABLE users ADD api_key VARCHAR(64) UNIQUE;


-- +goose Down
ALTER TABLE users DROP COLUMN api_key;
