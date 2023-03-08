package administrator

import (
	"context"
	"yuumi/app/service/administrator/internal/data/mysql/schema"
	"yuumi/app/service/administrator/internal/data/mysql/transform"
	"yuumi/pkg/crypto"

	"gorm.io/gorm"
)

func (m *Model) UpdateWithID(ctx context.Context, id int64, v interface{}) (*schema.Administrator, error) {
	administrator := schema.Administrator{Model: gorm.Model{ID: uint(id)}}

	result := m.GetDB().WithContext(ctx).Model(&administrator).Omit(m.Feilds.Salt).Updates(v)
	if result.Error != nil {
		return nil, result.Error
	}

	return m.FindOne(ctx, &FindCondition{ID: id})
}

func (m *Model) UpdatePassword(ctx context.Context, id int64, password string, salt string) (*schema.Administrator, error) {
	return m.UpdateWithID(ctx, id, schema.Administrator{Password: crypto.HmacSha256(salt, password)})
}

func (model *Model) AppendRolesWithAdministratorID(ctx context.Context, administratorId int64, roleIds []int64) error {
	administrator := schema.Administrator{Model: gorm.Model{ID: uint(administratorId)}}
	roles := transform.TransformRolesWithIds(roleIds...)
	tx := model.GetDB().WithContext(ctx)
	err := tx.Model(&administrator).Association(schema.Relation.Roles).Append(roles)
	return err
}

func (model *Model) DeleteRolesWithAdministratorID(ctx context.Context, administratorId int64, roleIds []int64) error {
	administrator := schema.Administrator{Model: gorm.Model{ID: uint(administratorId)}}
	roles := transform.TransformRolesWithIds(roleIds...)
	tx := model.GetDB().WithContext(ctx)
	err := tx.Model(&administrator).Association(schema.Relation.Roles).Delete(roles)
	return err
}
