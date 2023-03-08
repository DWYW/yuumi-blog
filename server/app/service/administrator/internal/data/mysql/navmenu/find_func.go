package navmenu

import (
	"context"
	"fmt"
	"yuumi/app/service/administrator/internal/data/mysql/schema"

	"gorm.io/gorm"
)

type FindCondition struct {
	ID      int64
	Name    string
	Sort    *FindConditionSort
	Preload *FindConditionPreload
}

type FindConditionSort struct {
	Weight int64
}

type FindConditionPreload struct {
	Roles bool
}

func (condition *FindCondition) Build() []func(*gorm.DB) *gorm.DB {
	fields := schema.NavMenu{}.GetFeilds()
	var scopes = make([]func(*gorm.DB) *gorm.DB, 0)

	if condition.ID > 0 {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("`%s` = ?", fields.ID), condition.ID)
		})
	}

	if condition.Name != "" {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("`%s` = ?", fields.Name), condition.Name)
		})
	}

	if condition.Sort != nil {
		scopes = append(scopes, condition.Sort.Build()...)
	}

	return scopes
}

func (condition *FindConditionSort) Build() []func(*gorm.DB) *gorm.DB {
	fields := schema.NavMenu{}.GetFeilds()
	var scopes = make([]func(*gorm.DB) *gorm.DB, 0)

	if condition.Weight != 0 {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			var desc string
			if condition.Weight > 0 {
				desc = fmt.Sprintf("`%s`", fields.Weight)
			} else {
				desc = fmt.Sprintf("`%s` desc", fields.Weight)
			}
			return db.Order(desc)
		})
	}

	return scopes
}

func (model *Model) FindOne(ctx context.Context, condition *FindCondition) (*schema.NavMenu, error) {
	dest := schema.NavMenu{}
	result := model.GetDB().WithContext(ctx).Scopes(condition.Build()...).First(&dest)
	return &dest, result.Error
}

func (model *Model) Find(ctx context.Context, condition *FindCondition, scopes ...func(*gorm.DB) *gorm.DB) ([]*schema.NavMenu, error) {
	var dest []*schema.NavMenu
	tx := model.GetDB().WithContext(ctx)
	if condition.Preload != nil {
		if condition.Preload.Roles {
			tx = tx.Preload(schema.Relation.Roles)
		}
	}
	result := tx.Scopes(scopes...).Scopes(condition.Build()...).Find(&dest)
	return dest, result.Error
}

func (model *Model) FindByID(ctx context.Context, id int64) (*schema.NavMenu, error) {
	return model.FindOne(ctx, &FindCondition{ID: id})
}

type FindListCondition struct {
	Page          int64
	PageSize      int64
	FindCondition *FindCondition
}

func (model *Model) FindList(ctx context.Context, condition *FindListCondition) ([]*schema.NavMenu, error) {
	skipNum := (condition.Page - 1) * condition.PageSize
	return model.Find(ctx, condition.FindCondition, func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(skipNum)).Limit(int(condition.PageSize))
	})
}

type FindWithRoleIDsCondition struct {
	RoleIDS       []int64
	FindCondition *FindCondition
}

func (model *Model) FindWithRoleIDs(ctx context.Context, condition *FindWithRoleIDsCondition) ([]*schema.NavMenu, error) {
	var dest []*schema.NavMenu
	roles := make([]schema.Role, len(condition.RoleIDS))
	for i, item := range condition.RoleIDS {
		roles[i].ID = uint(item)
	}

	tx := model.GetDB().WithContext(ctx)
	tx = tx.Model(roles)
	if condition.FindCondition != nil {
		tx = tx.Scopes(condition.FindCondition.Build()...)
	}
	err := tx.Association(schema.Relation.NavMenus).Find(&dest)
	return dest, err
}
