package models

import (
	"errors"
	"plant-manager/internal/types"
	"plant-manager/internal/utils"

	"xorm.io/xorm"
)

type PlantTypes interface {
	Get(id, customerID int64) (plantType *types.PlantType, err error)
	Find(customerID, limit, offset int64) ([]*types.PlantType, error)
	Create(plantType *types.PlantType) error
	Update(id int64, name string) error
	Destroy(id int64) error
}

func NewPlantType(db *xorm.Engine) PlantTypes {
	return &plantTypes{db}
}

type plantTypes struct {
	db *xorm.Engine
}

func (c *plantTypes) Get(id, customerID int64) (plantType *types.PlantType, err error) {
	plantType = new(types.PlantType)

	res, err := c.db.Query(`
		SELECT
			plant_types.id AS id,
			plant_types.name AS name,
			plant_types.customer_id AS customer_id,
			json_agg(plant_varieties)::jsonb AS varieties
		FROM plant_types
		LEFT JOIN plant_varieties ON plant_varieties.plant_type_id = plant_types.id
		WHERE plant_types.id = $1
		AND plant_types.customer_id = $2
		GROUP BY plant_types.id
	`, id, customerID)
	if err != nil {
		return
	}

	for _, r := range res {
		plantType, err = plantTypeFromRow(r)
		if err != nil {
			return
		}
		break
	}

	if plantType == nil || plantType.ID != id || plantType.CustomerID != customerID {
		return nil, errors.New("unable to find object")
	}

	return
}

func (c *plantTypes) Find(customerID, limit, offset int64) (plantTypes []*types.PlantType, err error) {

	res, err := c.db.Query(`
		SELECT
			plant_types.id AS id,
			plant_types.name AS name,
			plant_types.customer_id AS customer_id,
			json_agg(plant_varieties)::jsonb AS varieties
		FROM plant_types
		LEFT JOIN plant_varieties ON plant_varieties.plant_type_id = plant_types.id
		WHERE plant_types.customer_id = $1
		GROUP BY plant_types.id
		LIMIT $2
		OFFSET $3
	`, customerID, limit, offset)
	if err != nil {
		return
	}

	for _, r := range res {
		var plantType *types.PlantType
		plantType, err = plantTypeFromRow(r)
		if err != nil {
			return
		}

		plantTypes = append(plantTypes, plantType)
	}

	return
}

func (c *plantTypes) Create(plantType *types.PlantType) (err error) {
	if err = types.Validate(plantType); err != nil {
		return
	}

	plantType.CreatedAt = utils.TimeNow()
	plantType.UpdatedAt = utils.TimeNow()

	if _, err = c.db.Cols(
		"id",
		"name",
		"customer_id",
	).Insert(plantType); err != nil {
		return
	}

	return
}

func (c *plantTypes) Destroy(id int64) (err error) {
	if _, err = c.db.ID(id).Delete(&types.PlantType{}); err != nil {
		return
	}

	return
}

func (p *plantTypes) Update(id int64, name string) (err error) {
	if _, err = p.db.ID(id).Update(&types.PlantType{Name: name}); err != nil {
		return
	}

	return
}
