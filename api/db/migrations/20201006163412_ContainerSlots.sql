-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS container_slots (
  id                   serial             NOT NULL PRIMARY key,
  name                 text               NOT NULL,
  container_id             int                NOT NULL REFERENCES containers (id),
  row_id               int                NOT NULL,
  column_id            int                NOT NULL,
  customer_id          int                NOT NULL REFERENCES customers (id),
  created_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(container_id, row_id, column_id, customer_id)
);

COMMENT ON TABLE container_slots IS '
    container_slots represents a slot in a container

    Interesting columns:
        container_id - will be what container this container resides
        row_id - will be a row spot in a container
        column_id - will be a column spot in a container
';

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS container_slots;
