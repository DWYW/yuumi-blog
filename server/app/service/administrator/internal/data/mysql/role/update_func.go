package role

import (
	"context"
	"yuumi/app/service/administrator/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func (model *Model) UpdateWithID(ctx context.Context, id int64, v interface{}) (*schema.Role, error) {
	role := schema.Role{Model: gorm.Model{ID: uint(id)}}

	result := model.GetDB().WithContext(ctx).Model(&role).Updates(v)
	if result.Error != nil {
		return nil, result.Error
	}

	return model.FindOne(ctx, &FindCondition{ID: id})
}

func createPermissionsWithIds(ids []int64) []*schema.Permission {
	var permissions []*schema.Permission
	for _, id := range ids {
		permissions = append(permissions, &schema.Permission{Model: gorm.Model{ID: uint(id)}})
	}
	return permissions
}

func (model *Model) AppendPermissionsWithRoleID(ctx context.Context, roleId int64, permissionIds []int64) error {
	role := schema.Role{Model: gorm.Model{ID: uint(roleId)}}
	permissions := createPermissionsWithIds(permissionIds)
	tx := model.GetDB().WithContext(ctx)
	err := tx.Model(&role).Association(schema.Relation.Permissions).Append(permissions)
	return err
}

func (model *Model) DeletePermissionsWithRoleID(ctx context.Context, roleId int64, permissionIds []int64) error {
	role := schema.Role{Model: gorm.Model{ID: uint(roleId)}}
	permissions := createPermissionsWithIds(permissionIds)
	tx := model.GetDB().WithContext(ctx)
	err := tx.Model(&role).Association(schema.Relation.Permissions).Delete(permissions)
	return err
}
