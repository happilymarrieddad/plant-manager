-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS permissions (
  id                   serial             NOT NULL PRIMARY key,
  name                 text               NOT NULL,
  user_id              int                NOT NULL,
  created_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(name)
);

COMMENT ON TABLE permissions IS '
    permissions represents the ability for a user to do something

    Example:
        Example
';

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS permissions;
