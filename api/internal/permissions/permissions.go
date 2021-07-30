package permissions

type Permission string

func (p Permission) String() string {
	return string(p)
}

var AllPermissions []Permission
var PermissionRoutes map[string]Permission

var (
	PermissionSystemUser Permission = "system.user"
)

func init() {
	AllPermissions = append(AllPermissions, PermissionCustomersRead)
	AllPermissions = append(AllPermissions, PermissionCustomersCreate)
	AllPermissions = append(AllPermissions, PermissionCustomersDelete)

	AllPermissions = append(AllPermissions, PermissionPlacesRead)
	AllPermissions = append(AllPermissions, PermissionPlacesCreate)
	AllPermissions = append(AllPermissions, PermissionPlacesUpdate)
	AllPermissions = append(AllPermissions, PermissionPlacesDelete)

	AllPermissions = append(AllPermissions, PermissionPlaceSlotsRead)
	AllPermissions = append(AllPermissions, PermissionPlaceSlotsUpdate)

	AllPermissions = append(AllPermissions, PermissionUsersRead)
	AllPermissions = append(AllPermissions, PermissionUsersCreate)
	AllPermissions = append(AllPermissions, PermissionUsersDelete)

	AllPermissions = append(AllPermissions, PermissionPlantTypesRead)
	AllPermissions = append(AllPermissions, PermissionPlantTypesCreate)
	AllPermissions = append(AllPermissions, PermissionPlantTypesDelete)
	AllPermissions = append(AllPermissions, PermissionPlantTypesUpdate)

	AllPermissions = append(AllPermissions, PermissionVarietiesRead)
	AllPermissions = append(AllPermissions, PermissionVarietiesCreate)
	AllPermissions = append(AllPermissions, PermissionVarietiesDelete)
	AllPermissions = append(AllPermissions, PermissionVarietiesUpdate)

	// These are the permissions that are required for specific routes
	PermissionRoutes = make(map[string]Permission)

	PermissionRoutes["/customers"] = PermissionCustomersRead
	PermissionRoutes["/customers/create"] = PermissionCustomersCreate

	PermissionRoutes["/places"] = PermissionPlacesRead
	PermissionRoutes["/places/create"] = PermissionPlacesCreate

	PermissionRoutes["/plant_types"] = PermissionPlantTypesRead
	PermissionRoutes["/plant_types/create"] = PermissionPlantTypesCreate

	PermissionRoutes["/users"] = PermissionUsersRead
	PermissionRoutes["/users/create"] = PermissionUsersCreate
}
