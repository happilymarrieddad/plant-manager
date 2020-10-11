-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS place_slots (
  id                   serial             NOT NULL PRIMARY key,
  place_id             int                NOT NULL,
  name                 text               NOT NULL,
  row_id               int                NOT NULL,
  column_id            int                NOT NULL,
  created_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(place_id, row_id, column_id)
);

COMMENT ON TABLE place_slots IS '
    place_slots represents a slot in a place

    Interesting columns:
        place_id - will be what place this container resides
        row_id - will be a row spot in a place
        column_id - will be a column spot in a place
';

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS place_slots;
