package navmenu

import (
	"context"
	"yuumi/app/service/administrator/internal/data/mysql/schema"
	"yuumi/app/service/administrator/internal/data/mysql/transform"

	"gorm.io/gorm"
)

func (model *Model) BindWithRoleIDs(ctx context.Context, id int64, roleIds []int64) error {
	menu := schema.NavMenu{Model: gorm.Model{ID: uint(id)}}

	roles := transform.TransformRolesWithIds(roleIds...)
	tx := model.GetDB().WithContext(ctx)
	err := tx.Model(&menu).Association(model.Relations.Roles).Append(roles)
	return err
}

func (model *Model) UnbindWithRoleIDs(ctx context.Context, id int64, roleIds []int64) error {
	menu := schema.NavMenu{Model: gorm.Model{ID: uint(id)}}

	roles := transform.TransformRolesWithIds(roleIds...)
	tx := model.GetDB().WithContext(ctx)
	err := tx.Model(&menu).Association(model.Relations.Roles).Delete(roles)
	return err
}
