package apiutils

import (
	"context"
	"plant-manager/internal/models"
	"plant-manager/internal/types"
)

func RetrieveCustomersModelFromContext(ctx context.Context) models.Customers {
	if c, ok := ctx.Value(globalCustomersKey).(models.Customers); ok {
		return c
	}
	return nil
}

func RetrieveUsersModelFromContext(ctx context.Context) models.Users {
	if c, ok := ctx.Value(globalUsersKey).(models.Users); ok {
		return c
	}
	return nil
}

func RetrievePlacesModelFromContext(ctx context.Context) models.Places {
	if c, ok := ctx.Value(globalPlacesKey).(models.Places); ok {
		return c
	}
	return nil
}

func RetrievePermissionsModelFromContext(ctx context.Context) models.Permissions {
	if c, ok := ctx.Value(globalPermissionsKey).(models.Permissions); ok {
		return c
	}
	return nil
}

func RetrieveUserFromContext(ctx context.Context) *types.User {
	if u, ok := ctx.Value(storageUserKey).(*types.User); ok {
		return u
	}
	return nil
}
