package permission

import (
	"context"
	"yuumi/app/service/administrator/internal/data/mysql/schema"

	"gorm.io/gorm"
)

type FindWithRoleCondition struct {
	RoleID        int64
	FindCondition *FindCondition
}

// 查找角色的所有权限
func (model *Model) FindWithRole(ctx context.Context, condition *FindWithRoleCondition) ([]*schema.Permission, error) {
	var permissions []*schema.Permission
	role := schema.Role{Model: gorm.Model{ID: uint(condition.RoleID)}}

	tx := model.GetDB().WithContext(ctx).Model(&role)
	if condition.FindCondition != nil {
		tx = tx.Scopes(condition.FindCondition.Build()...)
	}

	err := tx.Association(role.GetRelations().Permissions).Find(&permissions)
	return permissions, err
}
