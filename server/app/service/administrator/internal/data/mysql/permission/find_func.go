package permission

import (
	"context"
	"fmt"
	"yuumi/app/service/administrator/internal/data/mysql/schema"

	"gorm.io/gorm"
)

type FindCondition struct {
	ID        int64
	RpcMethod string
	Sort      *FindConditionSort
	Preload   *FindConditionPreload
}

type FindConditionSort struct {
	ID int64
}

type FindConditionPreload struct {
	Roles bool
}

func (condition *FindCondition) Build() []func(*gorm.DB) *gorm.DB {
	var scopes = make([]func(*gorm.DB) *gorm.DB, 0)

	fields := schema.Permission{}.GetFeilds()
	if condition.ID > 0 {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("`%s` = ?", fields.ID), condition.ID)
		})
	}

	if condition.RpcMethod != "" {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("`%s` = ?", fields.RpcMethod), condition.RpcMethod)
		})
	}

	if condition.Sort != nil && condition.Sort.ID > 0 {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Order(fmt.Sprintf("%s asc", fields.ID))
		})
	} else {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Order(fmt.Sprintf("%s desc", fields.ID))
		})
	}

	return scopes
}

func (model *Model) FindOne(ctx context.Context, condition *FindCondition) (*schema.Permission, error) {
	dest := schema.Permission{}
	result := model.GetDB().WithContext(ctx).Scopes(condition.Build()...).First(&dest)
	return &dest, result.Error
}

func (model *Model) Find(ctx context.Context, condition *FindCondition, scopes ...func(*gorm.DB) *gorm.DB) ([]*schema.Permission, error) {
	var dest []*schema.Permission
	tx := model.GetDB().WithContext(ctx)
	if condition.Preload != nil {
		if condition.Preload.Roles {
			tx = tx.Preload(schema.Relation.Roles)
		}
	}
	result := tx.Scopes(scopes...).Scopes(condition.Build()...).Find(&dest)
	return dest, result.Error
}

func (model *Model) FindByID(ctx context.Context, id int64) (*schema.Permission, error) {
	return model.FindOne(ctx, &FindCondition{ID: id})
}

type FindListCondition struct {
	Page          int64
	PageSize      int64
	FindCondition *FindCondition
}

func (model *Model) FindList(ctx context.Context, condition *FindListCondition) ([]*schema.Permission, error) {
	skipNum := (condition.Page - 1) * condition.PageSize
	return model.Find(ctx, condition.FindCondition, func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(skipNum)).Limit(int(condition.PageSize))
	})
}

type FindWithRoleCondition struct {
	RoleID        int64
	FindCondition *FindCondition
}

func (model *Model) FindWithRole(ctx context.Context, condition *FindWithRoleCondition) ([]*schema.Permission, error) {
	var permissions []*schema.Permission
	tx := model.GetDB().WithContext(ctx)
	tx = tx.Model(&schema.Role{Model: gorm.Model{ID: uint(condition.RoleID)}})
	if condition.FindCondition != nil {
		tx = tx.Scopes(condition.FindCondition.Build()...)
	}
	err := tx.Association(schema.Relation.Permissions).Find(&permissions)
	return permissions, err
}
