-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- TODO: Need to potentially add when planted and other properties
CREATE TABLE IF NOT EXISTS plant_varieties (
  id                   serial             NOT NULL PRIMARY key,
  name                 text               NOT NULL,
  plant_type_id        int                NOT NULL REFERENCES plant_types (id),
  customer_id          int                NOT NULL REFERENCES customers (id),
  created_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(name, plant_type_id, customer_id)
);

COMMENT ON TABLE plant_varieties IS '
    plant_varieties represents a plant variety for a plant type

    Example:
        Hardyblue
        Jersey
        Legacy
';

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS plant_varieties;
