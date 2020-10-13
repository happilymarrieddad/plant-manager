package models

import (
	"plant-manager/internal/types"
	"plant-manager/internal/utils"

	"xorm.io/xorm"
)

type Varieties interface {
	Get(id, customerID int64) (plantType *types.Variety, err error)
	Find(customerID, limit, offset int64) ([]*types.Variety, error)
	Create(plantType *types.Variety) error
	Update(id int64, name string) error
	Destroy(id int64) error
}

func NewVariety(db *xorm.Engine) Varieties {
	return &varieties{db}
}

type varieties struct {
	db *xorm.Engine
}

func (c *varieties) Get(id, customerID int64) (plantType *types.Variety, err error) {
	plantType = new(types.Variety)
	has, err := c.db.ID(id).Where("customer_id = ?", customerID).Get(plantType)
	if !has {
		return nil, ErrNotFound
	}

	return
}

func (c *varieties) Find(customerID, limit, offset int64) (varieties []*types.Variety, err error) {

	if err = c.db.Where("customer_id = ?", customerID).
		Limit(int(limit), int(offset)).
		Find(&varieties); err != nil {
		return
	}

	return
}

func (c *varieties) Create(plantType *types.Variety) (err error) {
	if err = types.Validate(plantType); err != nil {
		return
	}

	plantType.CreatedAt = utils.TimeNow()
	plantType.UpdatedAt = utils.TimeNow()

	if _, err = c.db.Insert(plantType); err != nil {
		return
	}

	return
}

func (c *varieties) Destroy(id int64) (err error) {
	if _, err = c.db.ID(id).Delete(&types.Variety{}); err != nil {
		return
	}

	return
}

func (p *varieties) Update(id int64, name string) (err error) {
	if _, err = p.db.ID(id).Update(&types.Variety{Name: name}); err != nil {
		return
	}

	return
}
