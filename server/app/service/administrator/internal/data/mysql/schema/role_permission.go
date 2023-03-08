package schema

type RolePermissionFields struct {
	RoleID       string
	PermissionID string
}

type RolePermission struct {
	RoleID       uint
	PermissionID uint
}

func (RolePermission) GetFeilds() *RolePermissionFields {
	return &RolePermissionFields{
		RoleID:       "role_id",
		PermissionID: "permission_id",
	}
}
