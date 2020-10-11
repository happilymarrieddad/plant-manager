package places

import (
	"context"
	"plant-manager/internal/api/apiutils"
	"plant-manager/internal/permissions"
	"plant-manager/internal/types"
	pb "plant-manager/pb/go"
)

// grpcHandler auth routes
type grpcHandler struct {
}

// InitRoutes auth routes
func InitRoutes() pb.V1PlacesServer {
	return &grpcHandler{}
}

// Find a specific place
func (h *grpcHandler) GetPlace(ctx context.Context, req *pb.GetPlaceRequest) (reply *pb.GetPlaceReply, err error) {
	reply = new(pb.GetPlaceReply)
	if err = apiutils.UserHasPermission(ctx, permissions.PermissionPlacesRead); err != nil {
		return nil, err
	}

	cust, err := apiutils.RetrievePlacesModelFromContext(ctx).Get(req.GetId())
	if err != nil {
		return
	}

	reply.Place = cust.ToProtobuf()

	return
}

// Find all places
func (h *grpcHandler) FindPlaces(ctx context.Context, req *pb.FindPlacesRequest) (reply *pb.FindPlacesReply, err error) {
	reply = new(pb.FindPlacesReply)
	if err = apiutils.UserHasPermission(ctx, permissions.PermissionPlacesRead); err != nil {
		return nil, err
	}

	custs, err := apiutils.RetrievePlacesModelFromContext(ctx).Find()
	if err != nil {
		return
	}

	for _, c := range custs {
		reply.Places = append(reply.Places, c.ToProtobuf())
	}

	return
}

// Create a place
func (h *grpcHandler) CreatePlace(ctx context.Context, req *pb.CreatePlaceRequest) (reply *pb.CreatePlaceReply, err error) {
	reply = new(pb.CreatePlaceReply)
	if err = apiutils.UserHasPermission(ctx, permissions.PermissionPlacesCreate); err != nil {
		return nil, err
	}

	newPlace := &types.Place{
		Name: req.GetName(),
	}
	if err = apiutils.RetrievePlacesModelFromContext(ctx).Create(newPlace); err != nil {
		return
	}

	reply.Place = newPlace.ToProtobuf()

	return
}

// Destroy Place
func (h *grpcHandler) DestroyPlace(ctx context.Context, req *pb.DestroyPlaceRequest) (reply *pb.EmptyReply, err error) {
	reply = new(pb.EmptyReply)
	if err = apiutils.UserHasPermission(ctx, permissions.PermissionPlacesDelete); err != nil {
		return nil, err
	}

	if err = apiutils.RetrievePlacesModelFromContext(ctx).Destroy(req.GetId()); err != nil {
		return
	}

	return
}
