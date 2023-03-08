package administrator

import (
	"context"
	"fmt"
	"yuumi/app/service/administrator/internal/data/mysql/schema"

	"gorm.io/gorm"
)

type FindCondition struct {
	ID      int64
	Name    string
	Preload *FindConditionPreload
	Sort    *FindConditionSort
}

type FindConditionPreload struct {
	Roles bool
}

type FindConditionSort struct {
	ID int64
}

func (condition *FindCondition) Build() []func(*gorm.DB) *gorm.DB {
	var scopes = make([]func(*gorm.DB) *gorm.DB, 0)
	fields := schema.Administrator{}.GetFeilds()

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

func (model *Model) FindByID(ctx context.Context, id int64) (*schema.Administrator, error) {
	return model.FindOne(ctx, &FindCondition{ID: id})
}

func (model *Model) FindOne(ctx context.Context, condition *FindCondition) (*schema.Administrator, error) {
	dest := schema.Administrator{}
	result := model.GetDB().WithContext(ctx).Scopes(condition.Build()...).First(&dest)
	return &dest, result.Error
}

type FindListCondition struct {
	Page          int64
	PageSize      int64
	FindCondition *FindCondition
}

func (model *Model) FindList(ctx context.Context, condition *FindListCondition) ([]*schema.Administrator, error) {
	skipNum := (condition.Page - 1) * condition.PageSize
	return model.Find(ctx, condition.FindCondition, func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(skipNum)).Limit(int(condition.PageSize))
	})
}

func (model *Model) Find(ctx context.Context, condition *FindCondition, scopes ...func(*gorm.DB) *gorm.DB) ([]*schema.Administrator, error) {
	var dest []*schema.Administrator
	tx := model.GetDB().WithContext(ctx)

	if condition.Preload != nil {
		if condition.Preload.Roles {
			tx = tx.Preload(schema.Relation.Roles)
		}
	}

	result := tx.Scopes(scopes...).Scopes(condition.Build()...).Find(&dest)
	return dest, result.Error
}
