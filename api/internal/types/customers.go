package types

import (
	pb "plant-manager/pb/go"
	"time"
)

type Customer struct {
	ID        int64      `json:"id" xorm:"'id' pk autoincr"`
	Name      string     `json:"name" xorm:"name" validate:"required"`
	CreatedAt *time.Time `json:"created_at" xorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" xorm:"updated_at"`
}

func (u *Customer) TableName() string {
	return `customers`
}

func (u *Customer) ToProtobuf() *pb.Customer {
	return &pb.Customer{
		Id:   u.ID,
		Name: u.Name,
	}
}
