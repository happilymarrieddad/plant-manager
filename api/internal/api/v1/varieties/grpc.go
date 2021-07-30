package varieties

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
func InitRoutes() pb.V1VarietiesServer {
	return &grpcHandler{}
}

// Find a specific place
func (h *grpcHandler) GetVariety(ctx context.Context, req *pb.GetVarietyRequest) (reply *pb.GetVarietyReply, err error) {
	reply = new(pb.GetVarietyReply)
	user, err := apiutils.UserHasPermission(ctx, permissions.PermissionVarietiesRead)
	if err != nil {
		return nil, err
	}

	variety, err := apiutils.RetrieveVarietiesModelFromContext(ctx).Get(req.GetId(), user.CustomerID)
	if err != nil {
		return
	}

	reply.Variety = variety.ToProtobuf()

	return
}

// Find all varieties
func (h *grpcHandler) FindVarieties(ctx context.Context, req *pb.FindVarietiesRequest) (reply *pb.FindVarietiesReply, err error) {
	reply = new(pb.FindVarietiesReply)
	user, err := apiutils.UserHasPermission(ctx, permissions.PermissionVarietiesRead)
	if err != nil {
		return nil, err
	}

	vars, err := apiutils.RetrieveVarietiesModelFromContext(ctx).
		Find(user.CustomerID, req.GetLimit(), req.GetOffset())
	if err != nil {
		return
	}

	for _, c := range vars {
		reply.Varieties = append(reply.Varieties, c.ToProtobuf())
	}

	return
}

// Create a place
func (h *grpcHandler) CreateVariety(ctx context.Context, req *pb.CreateVarietyRequest) (reply *pb.CreateVarietyReply, err error) {
	reply = new(pb.CreateVarietyReply)
	user, err := apiutils.UserHasPermission(ctx, permissions.PermissionVarietiesCreate)
	if err != nil {
		return nil, err
	}

	newVariety := &types.Variety{
		Name:        req.GetName(),
		PlantTypeID: req.GetPlantTypeId(),
		CustomerID:  user.CustomerID,
	}
	if err = apiutils.RetrieveVarietiesModelFromContext(ctx).Create(newVariety); err != nil {
		return
	}

	reply.Variety = newVariety.ToProtobuf()

	return
}

func (h *grpcHandler) UpdateVariety(ctx context.Context, req *pb.UpdateVarietyRequest) (reply *pb.EmptyReply, err error) {
	reply = new(pb.EmptyReply)
	if _, err = apiutils.UserHasPermission(ctx, permissions.PermissionVarietiesUpdate); err != nil {
		return nil, err
	}

	if err = apiutils.RetrieveVarietiesModelFromContext(ctx).Update(
		req.GetId(),
		req.GetName(),
	); err != nil {
		return
	}

	return
}

// Destroy Variety
func (h *grpcHandler) DestroyVariety(ctx context.Context, req *pb.DestroyVarietyRequest) (reply *pb.EmptyReply, err error) {
	reply = new(pb.EmptyReply)
	if _, err = apiutils.UserHasPermission(ctx, permissions.PermissionVarietiesDelete); err != nil {
		return nil, err
	}

	if err = apiutils.RetrieveVarietiesModelFromContext(ctx).Destroy(req.GetId()); err != nil {
		return
	}

	return
}
