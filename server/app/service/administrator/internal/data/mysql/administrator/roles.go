package administrator

import (
	"context"
	"yuumi/app/service/administrator/internal/data/mysql/schema"
	"yuumi/app/service/administrator/internal/data/mysql/transform"

	"gorm.io/gorm"
)

func (model *Model) AppendRolesWithAdministratorID(ctx context.Context, administratorId int64, roleIds []int64) error {
	administrator := schema.Administrator{Model: gorm.Model{ID: uint(administratorId)}}
	roles := transform.TransformRolesWithIds(roleIds...)
	tx := model.GetDB().WithContext(ctx)
	err := tx.Model(&administrator).Association(model.Relations.Roles).Append(roles)
	return err
}

func (model *Model) DeleteRolesWithAdministratorID(ctx context.Context, administratorId int64, roleIds []int64) error {
	administrator := schema.Administrator{Model: gorm.Model{ID: uint(administratorId)}}
	roles := transform.TransformRolesWithIds(roleIds...)
	tx := model.GetDB().WithContext(ctx)
	err := tx.Model(&administrator).Association(model.Relations.Roles).Delete(roles)
	return err
}
