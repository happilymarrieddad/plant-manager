package models

import (
	"plant-manager/internal/types"
	"plant-manager/internal/utils"

	"xorm.io/xorm"
)

type Customers interface {
	Get(id int64) (cust *types.Customer, err error)
	Find() ([]*types.Customer, error)
	Create(cust *types.Customer) error
	Destroy(id int64) error
}

func NewCustomer(db *xorm.Engine) Customers {
	return &customers{db}
}

type customers struct {
	db *xorm.Engine
}

func (c *customers) Get(id int64) (cust *types.Customer, err error) {
	cust = new(types.Customer)
	has, err := c.db.ID(id).Get(cust)
	if !has {
		return nil, ErrNotFound
	}

	return
}

func (c *customers) Find() (custs []*types.Customer, err error) {

	if err = c.db.Find(&custs); err != nil {
		return
	}

	return
}

func (c *customers) Create(cust *types.Customer) (err error) {
	if err = types.Validate(cust); err != nil {
		return
	}

	cust.CreatedAt = utils.TimeNow()
	cust.UpdatedAt = utils.TimeNow()

	if _, err = c.db.Insert(cust); err != nil {
		return
	}

	return
}

func (c *customers) Destroy(id int64) (err error) {
	if _, err = c.db.ID(id).Delete(&types.Customer{}); err != nil {
		return
	}

	return
}
