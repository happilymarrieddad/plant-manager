package apiutils

import (
	"context"
	"plant-manager/internal/permissions"
)

func UserHasPermission(ctx context.Context, perm permissions.Permission) error {
	return RetrievePermissionsModelFromContext(ctx).
		UserHasPermission(RetrieveUserFromContext(ctx).ID, perm)
}
