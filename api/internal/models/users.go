package models

import (
	"plant-manager/internal/types"
	"plant-manager/internal/utils"

	"xorm.io/xorm"
)

type Users interface {
	Get(id int64) (user *types.User, err error)
	Find(customerID int64) (users []*types.User, err error)
	FindByEmail(email string) (user *types.User, err error)
	Create(user *types.User) error
}

func NewUser(db *xorm.Engine) Users {
	return &users{db}
}

type users struct {
	db *xorm.Engine
}

func (u *users) Get(id int64) (user *types.User, err error) {
	user = new(types.User)
	has, err := u.db.ID(id).Get(user)
	if err != nil {
		return
	}
	if !has {
		return nil, ErrNotFound
	}

	return
}

func (c *users) Find(customerID int64) (users []*types.User, err error) {
	if err = c.db.Where("customer_id = ?", customerID).Find(&users); err != nil {
		return
	}

	return
}

func (u *users) FindByEmail(email string) (user *types.User, err error) {
	user = new(types.User)
	found, err := u.db.Where("email = ?", email).Get(user)
	if err != nil {
		return
	}
	if !found {
		return nil, ErrNotFound
	}

	return
}

func (u *users) Create(user *types.User) (err error) {
	if err = types.Validate(user); err != nil {
		return
	}

	user.CreatedAt = utils.TimeNow()
	user.UpdatedAt = utils.TimeNow()

	if _, err = u.db.Insert(user); err != nil {
		return
	}

	return
}
