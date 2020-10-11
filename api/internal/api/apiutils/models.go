package apiutils

import (
	"context"
	"errors"
	"plant-manager/internal/models"
	"plant-manager/internal/utils"
	"reflect"

	"google.golang.org/grpc"
	"xorm.io/xorm"
)

const (
	globalUsersKey       = "v1models:users"
	globalCustomersKey   = "v1models:customers"
	globalPermissionsKey = "v1models:permissions"
	storageUserKey       = "storage:user"
)

func ModelInjector(db *xorm.Engine) grpc.UnaryServerInterceptor {
	return grpc.UnaryServerInterceptor(func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {

		// TODO: Do this in a separate handler at some point
		ctx = context.WithValue(
			ctx,
			globalUsersKey,
			models.NewUser(db),
		)
		ctx = context.WithValue(
			ctx,
			globalCustomersKey,
			models.NewCustomer(db),
		)
		ctx = context.WithValue(
			ctx,
			globalPermissionsKey,
			models.NewPermission(db),
		)

		// BEFORE the request
		v := reflect.Indirect(reflect.ValueOf(req))
		vField := reflect.Indirect(v.FieldByName("JWT"))

		// If there is NO jwt field on the request, then just continue
		if !vField.IsValid() {
			return handler(ctx, req)
		}

		jwtToken := vField.String()

		user, err := utils.GetDataFromToken(jwtToken)
		if err != nil {
			return nil, errors.New("unauthorized")
		}

		ctx = context.WithValue(
			ctx,
			storageUserKey,
			user,
		)

		h, err := handler(ctx, req)

		return h, err
	})
}
