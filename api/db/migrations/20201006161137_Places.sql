-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS places (
  id                   serial             NOT NULL PRIMARY key,
  name                 text               NOT NULL,
  rows                 int                NOT NULL,
  columns              int                NOT NULL,
  customer_id          int                NOT NULL REFERENCES customers (id),
  created_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(name, customer_id)
);

COMMENT ON TABLE places IS '
    places represents a named entity, and is unique by name
    
    Example:
        Garden Plot 1
        Greenhouse 1

    Interesting columns:
        rows - which rows are available starting with 0
        columns - which columns are available starting with 0
';

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS places;
