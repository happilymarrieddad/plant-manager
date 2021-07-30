-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE plants ADD COLUMN plant_variety_id int REFERENCES plant_varieties(id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE plants DROP COLUMN plant_variety_id;
