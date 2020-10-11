package models

import (
	"plant-manager/internal/types"
	"plant-manager/internal/utils"

	"xorm.io/xorm"
)

type Places interface {
	Get(id int64) (place *types.Place, err error)
	Find() ([]*types.Place, error)
	Create(place *types.Place) error
	Destroy(id int64) error
}

func NewPlace(db *xorm.Engine) Places {
	return &places{db}
}

type places struct {
	db *xorm.Engine
}

func (c *places) Get(id int64) (place *types.Place, err error) {
	place = new(types.Place)
	has, err := c.db.ID(id).Get(place)
	if !has {
		return nil, ErrNotFound
	}

	return
}

func (c *places) Find() (places []*types.Place, err error) {

	if err = c.db.Find(&places); err != nil {
		return
	}

	return
}

func (c *places) Create(place *types.Place) (err error) {
	if err = types.Validate(place); err != nil {
		return
	}

	place.CreatedAt = utils.TimeNow()
	place.UpdatedAt = utils.TimeNow()

	if _, err = c.db.Insert(place); err != nil {
		return
	}

	return
}

func (c *places) Destroy(id int64) (err error) {
	if _, err = c.db.ID(id).Delete(&types.Place{}); err != nil {
		return
	}

	return
}
