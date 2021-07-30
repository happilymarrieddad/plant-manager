-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS users (
  id                   serial             NOT NULL PRIMARY key,
  email                text               NOT NULL,
  first_name           text               NOT NULL,
  last_name            text               NOT NULL,
  password             text               NOT NULL,
  customer_id          int                NOT NULL REFERENCES customers (id),
  created_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(email)
);

COMMENT ON TABLE users IS '
    users represents a user in the system
';

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS users;
