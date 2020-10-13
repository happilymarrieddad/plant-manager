package customers

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
func InitRoutes() pb.V1CustomersServer {
	return &grpcHandler{}
}

// Find a specific customer
func (h *grpcHandler) GetCustomer(ctx context.Context, req *pb.GetCustomerRequest) (reply *pb.GetCustomerReply, err error) {
	reply = new(pb.GetCustomerReply)
	if _, err = apiutils.UserHasPermission(ctx, permissions.PermissionCustomersRead); err != nil {
		return nil, err
	}

	cust, err := apiutils.RetrieveCustomersModelFromContext(ctx).Get(req.GetId())
	if err != nil {
		return
	}

	reply.Customer = cust.ToProtobuf()

	return
}

// Find all customers
func (h *grpcHandler) FindCustomers(ctx context.Context, req *pb.FindCustomersRequest) (reply *pb.FindCustomersReply, err error) {
	reply = new(pb.FindCustomersReply)
	if _, err = apiutils.UserHasPermission(ctx, permissions.PermissionCustomersRead); err != nil {
		return nil, err
	}

	custs, err := apiutils.RetrieveCustomersModelFromContext(ctx).Find()
	if err != nil {
		return
	}

	for _, c := range custs {
		reply.Customers = append(reply.Customers, c.ToProtobuf())
	}

	return
}

// Create a customer
func (h *grpcHandler) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (reply *pb.CreateCustomerReply, err error) {
	reply = new(pb.CreateCustomerReply)
	if _, err = apiutils.UserHasPermission(ctx, permissions.PermissionSystemUser); err != nil {
		return nil, err
	}

	newCustomer := &types.Customer{
		Name: req.GetName(),
	}
	if err = apiutils.RetrieveCustomersModelFromContext(ctx).Create(newCustomer); err != nil {
		return
	}

	reply.Customer = newCustomer.ToProtobuf()

	return
}

// Destroy Customer
func (h *grpcHandler) DestroyCustomer(ctx context.Context, req *pb.DestroyCustomerRequest) (reply *pb.EmptyReply, err error) {
	reply = new(pb.EmptyReply)
	if _, err = apiutils.UserHasPermission(ctx, permissions.PermissionSystemUser); err != nil {
		return nil, err
	}

	if err = apiutils.RetrieveCustomersModelFromContext(ctx).Destroy(req.GetId()); err != nil {
		return
	}

	return
}
