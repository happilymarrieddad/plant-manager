package place_slots

import (
	"context"
	"plant-manager/internal/api/apiutils"
	"plant-manager/internal/permissions"
	pb "plant-manager/pb/go"
)

// grpcHandler auth routes
type grpcHandler struct {
}

// InitRoutes place slots routes
func InitRoutes() pb.V1PlaceSlotsServer {
	return &grpcHandler{}
}

func (h *grpcHandler) GetPlaceSlot(ctx context.Context, req *pb.GetPlaceSlotRequest) (reply *pb.GetPlaceSlotReply, err error) {
	reply = new(pb.GetPlaceSlotReply)
	user, err := apiutils.UserHasPermission(ctx, permissions.PermissionPlaceSlotsRead)
	if err != nil {
		return nil, err
	}

	slot, err := apiutils.RetrievePlaceSlotsModelFromContext(ctx).Get(
		req.GetId(),
		user.CustomerID,
	)
	if err != nil {
		return
	}

	reply.PlaceSlot = slot.ToProtobuf()

	return
}

func (h *grpcHandler) UpdatePlaceSlot(ctx context.Context, req *pb.UpdatePlaceSlotRequest) (reply *pb.EmptyReply, err error) {
	reply = new(pb.EmptyReply)
	if _, err = apiutils.UserHasPermission(ctx, permissions.PermissionPlaceSlotsUpdate); err != nil {
		return nil, err
	}

	if err = apiutils.RetrievePlaceSlotsModelFromContext(ctx).Update(
		req.GetId(),
		req.GetName(),
	); err != nil {
		return
	}

	return
}
