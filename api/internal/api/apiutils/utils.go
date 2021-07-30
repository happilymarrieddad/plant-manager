package apiutils

import (
	"context"
	"plant-manager/internal/permissions"
	"plant-manager/internal/types"
)

func UserHasPermission(ctx context.Context, perm permissions.Permission) (*types.User, error) {
	user := RetrieveUserFromContext(ctx)
	return user, RetrievePermissionsModelFromContext(ctx).
		UserHasPermission(user.ID, perm)
}
