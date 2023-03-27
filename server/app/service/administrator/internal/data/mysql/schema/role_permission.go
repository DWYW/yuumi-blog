package schema

type RolePermissionFields struct {
	RoleID       string
	PermissionID string
}

type RolePermission struct {
	RoleID       uint
	PermissionID uint
}

func (RolePermission) TableName() string {
	return "role_permissions"
}

func (RolePermission) GetFields() *RolePermissionFields {
	return &RolePermissionFields{
		RoleID:       "role_id",
		PermissionID: "permission_id",
	}
}
