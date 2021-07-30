package models

import (
	"errors"
	"plant-manager/internal/types"
	"plant-manager/internal/utils"

	"xorm.io/xorm"
)

// Places - places to store plants
type Places interface {
	Get(id, customerID int64) (place *types.Place, err error)
	Find(customerID, limit, offset int64) ([]*types.Place, error)
	Create(place *types.Place) error
	Update(id int64, name string) error
	UpdateSlot(id int64, name string) error
	Destroy(id int64) error
}

// NewPlace - create new place
func NewPlace(db *xorm.Engine) Places {
	return &places{db}
}

type places struct {
	db *xorm.Engine
}

func (p *places) Get(id, customerID int64) (place *types.Place, err error) {
	place = new(types.Place)

	res, err := p.db.Query(`
		SELECT
			places.id AS id,
			places.name AS name,
			places.rows AS rows,
			places.columns AS columns,
			places.customer_id AS customer_id,
			json_agg(place_slots)::jsonb AS slots
		FROM places
		LEFT JOIN place_slots ON place_slots.place_id = places.id
		WHERE places.id = $1
		AND places.customer_id = $2
		GROUP BY places.id
		LIMIT 1
	`, id, customerID)
	if err != nil {
		return
	}

	for _, r := range res {
		place, err = placeFromRow(r)
		if err != nil {
			return
		}
		break
	}

	if place == nil || place.ID != id || place.CustomerID != customerID {
		return nil, errors.New("unable to find object")
	}

	return
}

func (p *places) Find(customerID, limit, offset int64) (places []*types.Place, err error) {
	if limit == 0 {
		limit = 25
	}
	if offset < 0 {
		offset = 0
	}

	res, err := p.db.Query(`
		SELECT
			places.id AS id,
			places.name AS name,
			places.rows AS rows,
			places.columns AS columns,
			places.customer_id AS customer_id,
			json_agg(place_slots)::jsonb AS slots
		FROM places
		LEFT JOIN place_slots ON place_slots.place_id = places.id
		WHERE places.customer_id = $1
		GROUP BY places.id
		LIMIT $2
		OFFSET $3
	`, customerID, limit, offset)
	if err != nil {
		return
	}

	for _, r := range res {
		place := new(types.Place)

		place, err = placeFromRow(r)

		places = append(places, place)
	}

	return
}

func (p *places) Create(place *types.Place) (err error) {
	if err = types.Validate(place); err != nil {
		return
	}

	t := utils.TimeNow()

	place.CreatedAt = t
	place.UpdatedAt = t

	if _, err = p.db.Cols(
		"id",
		"name",
		"rows",
		"columns",
		"customer_id",
	).Insert(place); err != nil {
		return
	}

	var placeSlots []*types.PlaceSlot
	for rowID := int64(0); rowID < place.Rows; rowID++ {
		for columnID := int64(0); columnID < place.Columns; columnID++ {
			placeSlots = append(placeSlots, &types.PlaceSlot{
				PlaceID:    place.ID,
				RowID:      rowID,
				ColumnID:   columnID,
				CustomerID: place.CustomerID,
				CreatedAt:  t,
				UpdatedAt:  t,
			})
		}
	}

	if _, err = p.db.Insert(placeSlots); err != nil {
		return
	}

	place.Slots = placeSlots

	return
}

func (p *places) Update(id int64, name string) (err error) {
	if _, err = p.db.ID(id).Update(&types.Place{Name: name}); err != nil {
		return
	}

	return
}

func (p *places) UpdateSlot(id int64, name string) (err error) {
	if _, err = p.db.ID(id).Update(&types.PlaceSlot{Name: name}); err != nil {
		return
	}

	return
}

func (p *places) Destroy(id int64) (err error) {
	if _, err = p.db.Where("place_id = ?", id).Delete(&types.PlaceSlot{}); err != nil {
		return
	}

	if _, err = p.db.ID(id).Delete(&types.Place{}); err != nil {
		return
	}

	return
}
