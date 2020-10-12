package types

import (
	pb "plant-manager/pb/go"
	"time"
)

// TODO: add address and longitude/latitude
type Place struct {
	ID         int64        `json:"id" xorm:"'id' pk autoincr"`
	Name       string       `json:"name" xorm:"name" validate:"required"`
	Rows       int64        `json:"rows" xorm:"rows" validate:"required"`
	Columns    int64        `json:"columns" xorm:"columns" validate:"required"`
	CustomerID int64        `json:"customer_d" xorm:"customer_id"`
	Slots      []*PlaceSlot `json:"slots" xorm:"slots"`
	CreatedAt  *time.Time   `json:"created_at" xorm:"created_at"`
	UpdatedAt  *time.Time   `json:"updated_at" xorm:"updated_at"`
}

func (p *Place) TableName() string {
	return "places"
}

func (p *Place) ToProtobuf() *pb.Place {
	return &pb.Place{
		Id:         p.ID,
		Name:       p.Name,
		Rows:       p.Rows,
		Columns:    p.Columns,
		CustomerId: p.CustomerID,
		Slots:      PlaceSlotsToProtobuf(p.Slots),
	}
}
