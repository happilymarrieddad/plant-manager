package plant_types

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
func InitRoutes() pb.V1PlantTypesServer {
	return &grpcHandler{}
}

// Find a specific plant_types
func (h *grpcHandler) GetPlantType(ctx context.Context, req *pb.GetPlantTypeRequest) (reply *pb.GetPlantTypeReply, err error) {
	reply = new(pb.GetPlantTypeReply)
	user, err := apiutils.UserHasPermission(ctx, permissions.PermissionPlantTypesRead)
	if err != nil {
		return nil, err
	}

	plantType, err := apiutils.RetrievePlantTypesModelFromContext(ctx).Get(req.GetId(), user.CustomerID)
	if err != nil {
		return
	}

	reply.PlantType = plantType.ToProtobuf()

	return
}

// Find all plant_typess
func (h *grpcHandler) FindPlantTypes(ctx context.Context, req *pb.FindPlantTypesRequest) (reply *pb.FindPlantTypesReply, err error) {
	reply = new(pb.FindPlantTypesReply)
	user, err := apiutils.UserHasPermission(ctx, permissions.PermissionPlantTypesRead)
	if err != nil {
		return nil, err
	}

	plantTypes, err := apiutils.RetrievePlantTypesModelFromContext(ctx).Find(user.CustomerID, req.GetLimit(), req.GetOffset())
	if err != nil {
		return
	}

	for _, c := range plantTypes {
		reply.PlantTypes = append(reply.PlantTypes, c.ToProtobuf())
	}

	return
}

// Create a plant_types
func (h *grpcHandler) CreatePlantType(ctx context.Context, req *pb.CreatePlantTypeRequest) (reply *pb.CreatePlantTypeReply, err error) {
	reply = new(pb.CreatePlantTypeReply)
	user, err := apiutils.UserHasPermission(ctx, permissions.PermissionPlantTypesCreate)
	if err != nil {
		return nil, err
	}

	newPlantType := &types.PlantType{
		Name:       req.GetName(),
		CustomerID: user.ID,
	}
	if err = apiutils.RetrievePlantTypesModelFromContext(ctx).Create(newPlantType); err != nil {
		return
	}

	reply.PlantType = newPlantType.ToProtobuf()

	return
}

// Destroy PlantType
func (h *grpcHandler) DestroyPlantType(ctx context.Context, req *pb.DestroyPlantTypeRequest) (reply *pb.EmptyReply, err error) {
	reply = new(pb.EmptyReply)
	if _, err = apiutils.UserHasPermission(ctx, permissions.PermissionPlantTypesDelete); err != nil {
		return nil, err
	}

	if err = apiutils.RetrievePlantTypesModelFromContext(ctx).Destroy(req.GetId()); err != nil {
		return
	}

	return
}

func (h *grpcHandler) UpdatePlantType(ctx context.Context, req *pb.UpdatePlantTypeRequest) (reply *pb.EmptyReply, err error) {
	reply = new(pb.EmptyReply)
	if _, err = apiutils.UserHasPermission(ctx, permissions.PermissionPlantTypesUpdate); err != nil {
		return nil, err
	}

	if err = apiutils.RetrievePlantTypesModelFromContext(ctx).Update(
		req.GetId(),
		req.GetName(),
	); err != nil {
		return
	}

	return
}
