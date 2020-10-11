-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS plant_types (
  id                   serial             NOT NULL PRIMARY key,
  name                 text               NOT NULL,
  created_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(name)
);

COMMENT ON TABLE plant_types IS '
    plant_types represents a named plant type

    Example:
        Blueberry
';

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS plant_types;
