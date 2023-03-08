package transform

import (
	"yuumi/app/service/administrator/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func TransformRolesWithIds(ids ...int64) []*schema.Role {
	var roles []*schema.Role
	for _, id := range ids {
		roles = append(roles, &schema.Role{Model: gorm.Model{ID: uint(id)}})
	}
	return roles
}
