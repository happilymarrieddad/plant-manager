-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS containers (
  id                   serial             NOT NULL PRIMARY key,
  name                 text               NOT NULL,
  place_slot_id        int                NOT NULL REFERENCES place_slots (id),
  rows                 int                NOT NULL,
  columns              int                NOT NULL,
  customer_id          int                NOT NULL REFERENCES customers (id),
  created_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(name, place_slot_id, customer_id)
);

COMMENT ON TABLE containers IS '
    containers represents a named entity, and is unique by name
    
    Example:
        Ground Spot 1
        5 Gallon Plastic Pot

    Interesting columns:
        place_id - will be what place this container resides
';

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS containers;
