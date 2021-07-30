package types

import (
	pb "plant-manager/pb/go"
	"time"
)

type PlantType struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	Name       string     `json:"name" xorm:"name" validate:"required"`
	CustomerID int64      `json:"customer_id" xorm:"customer_id"`
	Varieties  []*Variety `json:"varieties" xorm:"varieties"`
	CreatedAt  *time.Time `json:"created_at" xorm:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at" xorm:"updated_at"`
}

func (u *PlantType) TableName() string {
	return `plant_types`
}

func (u *PlantType) ToProtobuf() *pb.PlantType {
	return &pb.PlantType{
		Id:         u.ID,
		Name:       u.Name,
		CustomerId: u.CustomerID,
		Varieties:  VarietiesToProtobuf(u.Varieties),
	}
}
