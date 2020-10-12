-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- TODO: Need to potentially add when planted and other properties
CREATE TABLE IF NOT EXISTS plants (
  id                   serial             NOT NULL PRIMARY key,
  name                 text               NOT NULL,
  container_slot_id    int                NOT NULL REFERENCES container_slots (id),
  plant_type_id        int                NOT NULL REFERENCES plant_types (id),
  customer_id          int                NOT NULL REFERENCES customers (id),
  created_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(name, container_slot_id, plant_type_id, customer_id)
);

COMMENT ON TABLE plants IS '
    plants represents a plant

    Example:
        Blueberry
';

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS plants;
