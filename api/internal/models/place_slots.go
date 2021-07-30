package models

import (
	"plant-manager/internal/types"

	"xorm.io/xorm"
)

type PlaceSlots interface {
	Get(id, customerID int64) (*types.PlaceSlot, error)
	Update(id int64, name string) error
}

func NewPlaceSlots(db *xorm.Engine) PlaceSlots {
	return &placeSlots{db}
}

type placeSlots struct {
	db *xorm.Engine
}

func (p *placeSlots) Get(id, customerID int64) (placeSlot *types.PlaceSlot, err error) {
	placeSlot = new(types.PlaceSlot)

	if _, err = p.db.ID(id).Where("customer_id = ?", customerID).Get(placeSlot); err != nil {
		return
	}

	return
}

func (p *placeSlots) Update(id int64, name string) (err error) {
	if _, err = p.db.ID(id).Update(&types.PlaceSlot{Name: name}); err != nil {
		return
	}

	return
}
