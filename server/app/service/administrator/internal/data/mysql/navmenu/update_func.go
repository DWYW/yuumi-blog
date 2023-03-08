package navmenu

import (
	"context"
	"yuumi/app/service/administrator/internal/data/mysql/schema"
	"yuumi/app/service/administrator/internal/data/mysql/transform"

	"gorm.io/gorm"
)

func (model *Model) UpdateWithID(ctx context.Context, id int64, v interface{}) (*schema.NavMenu, error) {
	menu := schema.NavMenu{Model: gorm.Model{ID: uint(id)}}

	result := model.GetDB().WithContext(ctx).Model(&menu).Updates(v)
	if result.Error != nil {
		return nil, result.Error
	}

	return model.FindOne(ctx, &FindCondition{ID: id})
}

func (model *Model) BindWithRoleIDs(ctx context.Context, id int64, roleIds []int64) error {
	menu := schema.NavMenu{Model: gorm.Model{ID: uint(id)}}
	roles := transform.TransformRolesWithIds(roleIds...)
	tx := model.GetDB().WithContext(ctx)
	err := tx.Model(&menu).Association(schema.Relation.Roles).Append(roles)
	return err
}

func (model *Model) UnbindWithRoleIDs(ctx context.Context, id int64, roleIds []int64) error {
	menu := schema.NavMenu{Model: gorm.Model{ID: uint(id)}}
	roles := transform.TransformRolesWithIds(roleIds...)
	tx := model.GetDB().WithContext(ctx)
	err := tx.Model(&menu).Association(schema.Relation.Roles).Delete(roles)
	return err
}
