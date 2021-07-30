package models

import (
	"errors"
	perms "plant-manager/internal/permissions"
	"plant-manager/internal/types"
	"plant-manager/internal/utils"

	"xorm.io/xorm"
)

type Permissions interface {
	Create(permission *types.Permission) error
	UserHasPermission(userID int64, perm perms.Permission) error
	UserPermissions(userID int64) (perms []*types.Permission, err error)
}

func NewPermission(db *xorm.Engine) Permissions {
	return &permissions{db}
}

type permissions struct {
	db *xorm.Engine
}

func (p *permissions) Create(perm *types.Permission) (err error) {
	if err = types.Validate(perm); err != nil {
		return
	}

	perm.CreatedAt = utils.TimeNow()
	perm.UpdatedAt = utils.TimeNow()

	if _, err = p.db.Insert(perm); err != nil {
		return
	}

	return
}

func (p *permissions) UserHasPermission(userID int64, perm perms.Permission) error {
	has, err := p.db.Where("name = ?", perm).And("user_id = ?", userID).Get(&types.Permission{})
	if err != nil {
		return err
	} else if !has {
		return errors.New("unauthorized")
	}
	return nil
}

func (p *permissions) UserPermissions(userID int64) (perms []*types.Permission, err error) {
	if err = p.db.Where("user_id = ?", userID).Find(&perms); err != nil {
		return
	}

	return
}
