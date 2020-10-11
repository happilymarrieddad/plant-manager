-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS customers (
  id                   serial             NOT NULL PRIMARY key,
  name                 text               NOT NULL,
  created_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(name)
);

COMMENT ON TABLE customers IS '
    customers represents a named entity, and is unique by name
    
    Example:
        Example
';

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS customers;
