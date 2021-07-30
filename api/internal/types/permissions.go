package types

import (
	pb "plant-manager/pb/go"
	"time"
)

type Permission struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	Name       string     `json:"name" xorm:"name" validate:"required"`
	UserID     int64      `json:"user_id" xorm:"user_id" validate:"required"`
	CustomerID int64      `json:"customer_id" xorm:"customer_id" validate:"required"`
	CreatedAt  *time.Time `json:"created_at" xorm:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at" xorm:"updated_at"`
}

func (u *Permission) TableName() string {
	return `permissions`
}

func (u *Permission) ToProtobuf() *pb.Permission {
	return &pb.Permission{
		Id:     u.ID,
		Name:   u.Name,
		UserId: u.UserID,
	}
}
