package types

import (
	pb "plant-manager/pb/go"
	"time"
)

// TODO: add address and longitude/latitude
type Place struct {
	ID        int64      `json:"id" xorm:"'id' pk autoincr"`
	Name      string     `json:"name" xorm:"name" validate:"required"`
	CreatedAt *time.Time `json:"createdAt" xorm:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" xorm:"updated_at"`
}

func (p *Place) TableName() string {
	return "places"
}

func (p *Place) ToProtobuf() *pb.Place {
	return &pb.Place{
		Id:   p.ID,
		Name: p.Name,
	}
}
