package types

import (
	pb "plant-manager/pb/go"
	"time"
)

type Variety struct {
	ID          int64      `json:"id" xorm:"'id' pk autoincr"`
	Name        string     `json:"name" xorm:"name" validate:"required"`
	CustomerID  int64      `json:"customer_id" xorm:"customer_id"`
	PlantTypeID int64      `json:"plant_type_id" xorm:"plant_type_id"`
	CreatedAt   *time.Time `json:"created_at" xorm:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" xorm:"updated_at"`
}

func (u *Variety) TableName() string {
	return `plant_varieties`
}

func (u *Variety) ToProtobuf() *pb.Variety {
	return &pb.Variety{
		Id:          u.ID,
		Name:        u.Name,
		CustomerId:  u.CustomerID,
		PlantTypeId: u.PlantTypeID,
	}
}

func VarietiesToProtobuf(slots []*Variety) (ret []*pb.Variety) {
	for _, p := range slots {
		if p != nil {
			ret = append(ret, p.ToProtobuf())
		}
	}
	return
}

func ProtobufToVariety(slot *pb.Variety) (ret *Variety) {
	return &Variety{
		ID:         slot.GetId(),
		Name:       slot.GetName(),
		CustomerID: slot.GetCustomerId(),
	}
}

func ProtobufToVarieties(slots []*pb.Variety) (ret []*Variety) {
	for _, p := range slots {
		ret = append(ret, ProtobufToVariety(p))
	}
	return
}
