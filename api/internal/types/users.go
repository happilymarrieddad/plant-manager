package types

import (
	pb "plant-manager/pb/go"
	"time"
)

type User struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	Email      string     `json:"email" xorm:"email" validate:"required"`
	FirstName  string     `json:"first_name" xorm:"first_name" validate:"required"`
	LastName   string     `json:"last_name" xorm:"last_name" validate:"required"`
	Password   string     `json:"password" xorm:"password" validate:"required"`
	CustomerID int64      `json:"customer_id" xorm:"customer_id" validate:"required"`
	CreatedAt  *time.Time `json:"created_at" xorm:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at" xorm:"updated_at"`
}

func (u *User) TableName() string {
	return `users`
}

func (u *User) ToProtobuf() *pb.User {
	return &pb.User{
		Id:         u.ID,
		Email:      u.Email,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		CustomerId: u.CustomerID,
	}
}
