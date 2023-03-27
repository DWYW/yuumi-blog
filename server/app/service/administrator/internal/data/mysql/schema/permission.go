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

type PermissionRelations struct {
	Roles string
}

type Permission struct {
	gorm.Model
	Name      string
	RpcMethod string
	Roles     []*Role `gorm:"many2many:role_permissions;"`
}

func (Permission) TableName() string {
	return "permissions"
}

func (Permission) GetFields() *PermissionFields {
	return &PermissionFields{
		ID:        "id",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: "deleted_at",
		Name:      "name",
		RpcMethod: "rpc_method",
	}
}

func (Permission) GetRelations() *PermissionRelations {
	return &PermissionRelations{
		Roles: "Roles",
	}
}
