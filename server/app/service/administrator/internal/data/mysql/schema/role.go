package schema

import "gorm.io/gorm"

type RoleFields struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Name      string
	Type      string
	ParentID  string
}

type Role struct {
	gorm.Model
	ParentID       uint
	Name           string
	Type           string
	Administrators []*Administrator `gorm:"many2many:administrator_roles;"`
	Permissions    []*Permission    `gorm:"many2many:role_permissions;"`
	NavMenus       []*NavMenu       `gorm:"many2many:role_nav_menus;"`
}

func (Role) GetFeilds() *RoleFields {
	return &RoleFields{
		ID:        "id",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: "deleted_at",
		Name:      "name",
		Type:      "type",
		ParentID:  "parent_id",
	}
}
