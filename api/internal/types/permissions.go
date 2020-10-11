package types

import (
	pb "plant-manager/pb/go"
	"time"
)

type Permission struct {
	ID        int64      `json:"id" xorm:"'id' pk autoincr"`
	Name      string     `json:"name" xorm:"name" validate:"required"`
	UserID    int64      `json:"userId" xorm:"user_id" validate:"required"`
	CreatedAt *time.Time `json:"createdAt" xorm:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" xorm:"updated_at"`
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
