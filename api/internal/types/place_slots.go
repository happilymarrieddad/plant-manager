package types

import (
	pb "plant-manager/pb/go"
	"time"
)

type PlaceSlot struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	Name       string     `json:"name" xorm:"name" validate:"required"`
	PlaceID    int64      `json:"place_id" xorm:"place_id" validate:"required"`
	RowID      int64      `json:"row_id" xorm:"row_id" validate:"required"`
	ColumnID   int64      `json:"column_id" xorm:"column_id" validate:"required"`
	CustomerID int64      `json:"customer_id" xorm:"customer_id"`
	CreatedAt  *time.Time `json:"created_at" xorm:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at" xorm:"updated_at"`
}

func (p *PlaceSlot) TableName() string {
	return "place_slots"
}

func (p *PlaceSlot) ToProtobuf() *pb.PlaceSlot {
	return &pb.PlaceSlot{
		Id:         p.ID,
		Name:       p.Name,
		PlaceId:    p.PlaceID,
		RowId:      p.RowID,
		ColumnId:   p.ColumnID,
		CustomerId: p.CustomerID,
	}
}

func PlaceSlotsToProtobuf(slots []*PlaceSlot) (ret []*pb.PlaceSlot) {
	for _, p := range slots {
		ret = append(ret, p.ToProtobuf())
	}
	return
}

func ProtobufToPlaceSlot(slot *pb.PlaceSlot) (ret *PlaceSlot) {
	return &PlaceSlot{
		ID:         slot.GetId(),
		Name:       slot.GetName(),
		PlaceID:    slot.GetPlaceId(),
		RowID:      slot.GetRowId(),
		ColumnID:   slot.GetColumnId(),
		CustomerID: slot.GetCustomerId(),
	}
}

func ProtobufToPlaceSlots(slots []*pb.PlaceSlot) (ret []*PlaceSlot) {
	for _, p := range slots {
		ret = append(ret, ProtobufToPlaceSlot(p))
	}
	return
}
