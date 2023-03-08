package schema

import "gorm.io/gorm"

type PermissionFields struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Name      string
	RpcMethod string
}

type Permission struct {
	gorm.Model
	Name      string
	RpcMethod string
	Roles     []*Role `gorm:"many2many:role_permissions;"`
}

func (Permission) GetFeilds() *PermissionFields {
	return &PermissionFields{
		ID:        "id",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: "deleted_at",
		Name:      "name",
		RpcMethod: "rpc_method",
	}
}
