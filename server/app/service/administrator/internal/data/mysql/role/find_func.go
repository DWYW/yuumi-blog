package role

import (
	"context"
	"yuumi/app/service/administrator/internal/data/mysql/schema"

	"gorm.io/gorm"
)

type FindWithAdministratorCondition struct {
	AdministratorID int64
	FindCondition   *FindCondition
}

// 查找管理员拥有的角色
func (model *Model) FindWithAdministrator(ctx context.Context, condition *FindWithAdministratorCondition) ([]*schema.Role, error) {
	var roles []*schema.Role
	administrator := schema.Administrator{Model: gorm.Model{ID: uint(condition.AdministratorID)}}
	tx := model.GetDB().WithContext(ctx)
	tx = tx.Model(&administrator)
	if condition.FindCondition != nil {
		tx = tx.Scopes(condition.FindCondition.Build()...)
	}
	err := tx.Association(administrator.GetRelations().Roles).Find(&roles)
	return roles, err
}

type FindWithPermissionCondition struct {
	PermissionID  int64
	FindCondition *FindCondition
}

// 查找权限所属的角色
func (model *Model) FindWithPermission(ctx context.Context, condition *FindWithPermissionCondition) ([]*schema.Role, error) {
	var roles []*schema.Role
	permission := schema.Permission{Model: gorm.Model{ID: uint(condition.PermissionID)}}
	tx := model.GetDB().WithContext(ctx)
	tx = tx.Model(&permission)
	if condition.FindCondition != nil {
		tx = tx.Scopes(condition.FindCondition.Build()...)
	}
	err := tx.Association(permission.GetRelations().Roles).Find(&roles)
	return roles, err
}

// 通过父级查找
func (model *Model) FindWithParentIds(ctx context.Context, ids ...int64) ([]*schema.Role, error) {
	roles, err := model.Find(ctx, &FindCondition{
		Where: &FindConditionWhere{ParentIDs: ids},
	})
	if err != nil {
		return nil, err
	}

	// 递归查询
	if len(roles) > 0 {
		ids = make([]int64, len(roles))
		for i, role := range roles {
			ids[i] = int64(role.ID)
		}

		dest, err := model.FindWithParentIds(ctx, ids...)

		if err == nil {
			roles = append(roles, dest...)
		}
	}

	return roles, nil
}
