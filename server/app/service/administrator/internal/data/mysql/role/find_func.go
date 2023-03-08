package role

import (
	"context"
	"fmt"
	"yuumi/app/service/administrator/internal/data/mysql/schema"

	"gorm.io/gorm"
)

type FindCondition struct {
	ID        int64
	Type      string
	ParentIDs []int64
	Sort      *FindConditionSort
	Preload   *FindConditionPreload
}

type FindConditionPreload struct {
	Permissions bool
}

type FindConditionSort struct {
	ID int64
}

func (condition *FindCondition) Build() []func(*gorm.DB) *gorm.DB {
	var scopes = make([]func(*gorm.DB) *gorm.DB, 0)

	fields := schema.Role{}.GetFeilds()
	if condition.ID > 0 {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("`%s` = ?", fields.ID), condition.ID)
		})
	}

	if condition.Type != "" {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("`%s` = ?", fields.Type), condition.Type)
		})
	}

	if condition.ParentIDs != nil {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("`%s` in ?", fields.ParentID), condition.ParentIDs)
		})
	}

	if condition.Preload != nil && condition.Preload.Permissions {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Preload(schema.Relation.Permissions)
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

func (model *Model) FindOne(ctx context.Context, condition *FindCondition) (*schema.Role, error) {
	dest := schema.Role{}
	result := model.GetDB().WithContext(ctx).Scopes(condition.Build()...).First(&dest)
	return &dest, result.Error
}

func (model *Model) Find(ctx context.Context, condition *FindCondition, scopes ...func(*gorm.DB) *gorm.DB) ([]*schema.Role, error) {
	var dest []*schema.Role
	tx := model.GetDB().WithContext(ctx)
	result := tx.Scopes(scopes...).Scopes(condition.Build()...).Find(&dest)
	return dest, result.Error
}

func (model *Model) FindByID(ctx context.Context, id int64) (*schema.Role, error) {
	return model.FindOne(ctx, &FindCondition{ID: id})
}

type FindListCondition struct {
	Page          int64
	PageSize      int64
	FindCondition *FindCondition
}

func (model *Model) FindList(ctx context.Context, condition *FindListCondition) ([]*schema.Role, error) {
	skipNum := (condition.Page - 1) * condition.PageSize
	return model.Find(ctx, condition.FindCondition, func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(skipNum)).Limit(int(condition.PageSize))
	})
}

type FindWithAdministratorCondition struct {
	AdministratorID int64
	FindCondition   *FindCondition
}

func (model *Model) FindWithAdministrator(ctx context.Context, condition *FindWithAdministratorCondition) ([]*schema.Role, error) {
	var roles []*schema.Role
	tx := model.GetDB().WithContext(ctx)
	tx = tx.Model(&schema.Administrator{Model: gorm.Model{ID: uint(condition.AdministratorID)}})
	if condition.FindCondition != nil {
		tx = tx.Scopes(condition.FindCondition.Build()...)
	}
	err := tx.Association(schema.Relation.Roles).Find(&roles)
	return roles, err
}

type FindWithPermissionCondition struct {
	PermissionID  int64
	FindCondition *FindCondition
}

func (model *Model) FindWithPermission(ctx context.Context, condition *FindWithPermissionCondition) ([]*schema.Role, error) {
	var roles []*schema.Role
	tx := model.GetDB().WithContext(ctx)
	tx = tx.Model(&schema.Permission{Model: gorm.Model{ID: uint(condition.PermissionID)}})
	if condition.FindCondition != nil {
		tx = tx.Scopes(condition.FindCondition.Build()...)
	}
	err := tx.Association(schema.Relation.Roles).Find(&roles)
	return roles, err
}

func (model *Model) FindWithParentIds(ctx context.Context, ids ...int64) ([]*schema.Role, error) {
	roles, err := model.Find(ctx, &FindCondition{ParentIDs: ids})
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
