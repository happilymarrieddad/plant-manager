package users

import (
	"context"
	"errors"
	"fmt"
	"plant-manager/internal/api/apiutils"
	"plant-manager/internal/permissions"
	"plant-manager/internal/types"
	"plant-manager/internal/utils"
	pb "plant-manager/pb/go"
)

// grpcHandler auth routes
type grpcHandler struct {
}

// InitRoutes auth routes
func InitRoutes() pb.V1UsersServer {
	return &grpcHandler{}
}

func (g *grpcHandler) Login(ctx context.Context, req *pb.LoginRequest) (reply *pb.LoginReply, err error) {
	reply = new(pb.LoginReply)
	usersModel := apiutils.RetrieveUsersModelFromContext(ctx)

	user, err := usersModel.FindByEmail(req.GetEmail())
	if err != nil {
		return
	}

	if valid := types.IsPasswordValid(user.Password, req.GetPassword()); !valid {
		return nil, errors.New("invalid credentials")
	}

	claims := utils.GetNewClaims("user", map[string]interface{}{
		"user": user,
	})

	token, err := utils.GetSignedToken(claims)
	if err != nil {
		return
	}

	perms, err := apiutils.RetrievePermissionsModelFromContext(ctx).UserPermissions(user.ID)
	if err != nil {
		return
	}

	reply.Token = token
	for _, p := range perms {
		reply.Permissions = append(reply.Permissions, p.ToProtobuf())
	}

	return
}

func (g *grpcHandler) VerifyJWT(
	ctx context.Context, req *pb.VerifyJWTRequest,
) (reply *pb.VerifyJWTReply, err error) {
	user, err := utils.GetDataFromToken(req.GetJWT())
	if err != nil {
		fmt.Printf("VerifyJWT failed. Err: %s\n", err.Error())
		return &pb.VerifyJWTReply{
			Valid: false,
		}, nil
	}

	perms, err := apiutils.RetrievePermissionsModelFromContext(ctx).UserPermissions(user.ID)
	if err != nil {
		fmt.Printf("UserPermissions failed. Err: %s\n", err.Error())
		return &pb.VerifyJWTReply{
			Valid: false,
		}, nil
	}

	perm, routeHasPermission := permissions.PermissionRoutes[req.GetRoute()]
	var userHasPermission bool
	for _, p := range perms {
		if p.Name == string(perm) {
			userHasPermission = true
			break
		}
	}

	reply = &pb.VerifyJWTReply{
		Valid: utils.IsJWTValid(req.GetJWT()),
	}
	for _, p := range perms {
		reply.Permissions = append(reply.Permissions, p.ToProtobuf())
	}
	// If the route is protected and the user does not have permission then deny access
	if routeHasPermission && !userHasPermission {
		reply.HasPermission = false
		return
	}

	reply.HasPermission = true

	return
}

func (g *grpcHandler) AddUser(ctx context.Context, req *pb.AddUserRequest) (reply *pb.AddUserReply, err error) {
	reply = new(pb.AddUserReply)
	return reply, errors.New("unauthorized")
}

func (g *grpcHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (reply *pb.GetUserReply, err error) {
	reply = new(pb.GetUserReply)
	u, err := apiutils.UserHasPermission(ctx, permissions.PermissionUsersRead)
	if err != nil {
		return nil, err
	}

	user, err := apiutils.RetrieveUsersModelFromContext(ctx).Get(req.GetId(), u.CustomerID)
	if err != nil {
		return
	}

	reply.User = user.ToProtobuf()

	return
}

// Find all users
func (g *grpcHandler) FindUsers(ctx context.Context, req *pb.FindUsersRequest) (reply *pb.FindUsersReply, err error) {
	reply = new(pb.FindUsersReply)
	user, err := apiutils.UserHasPermission(ctx, permissions.PermissionUsersRead)
	if err != nil {
		return nil, err
	}

	custs, err := apiutils.RetrieveUsersModelFromContext(ctx).Find(user.CustomerID)
	if err != nil {
		return
	}

	for _, c := range custs {
		reply.Users = append(reply.Users, c.ToProtobuf())
	}

	return
}
