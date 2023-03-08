package githubuser

import (
	"context"
	"fmt"
	"yuumi/app/service/article/internal/data/mysql/schema"

	"gorm.io/gorm"
)

type FindCondition struct {
	ID   int64
	Sort *FindConditionSort
}

type FindConditionSort struct {
	ID int64
}

func (condition *FindCondition) Build() []func(*gorm.DB) *gorm.DB {
	var scopes = make([]func(*gorm.DB) *gorm.DB, 0)
	fields := schema.GithubUser{}.GetFeilds()

	if condition.ID > 0 {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("`%s` = ?", fields.ID), condition.ID)
		})
	}

	if condition.Sort != nil {
		scopes = append(scopes, condition.BuildOrder()...)
	}

	return scopes
}

func (condition *FindCondition) BuildOrder() []func(*gorm.DB) *gorm.DB {
	var scopes = make([]func(*gorm.DB) *gorm.DB, 0)
	fields := schema.GithubUser{}.GetFeilds()

	if condition.Sort.ID > 0 {
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

func (model *Model) FindByID(ctx context.Context, id int64) (*schema.GithubUser, error) {
	return model.FindOne(ctx, &FindCondition{ID: id})
}

func (model *Model) FindOne(ctx context.Context, condition *FindCondition) (*schema.GithubUser, error) {
	dest := schema.GithubUser{}
	result := model.GetDB().WithContext(ctx).Scopes(condition.Build()...).First(&dest)
	return &dest, result.Error
}

type FindListCondition struct {
	Page          int64
	PageSize      int64
	FindCondition *FindCondition
}

func (model *Model) FindList(ctx context.Context, condition *FindListCondition) ([]*schema.GithubUser, error) {
	skipNum := (condition.Page - 1) * condition.PageSize
	return model.Find(ctx, condition.FindCondition, func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(skipNum)).Limit(int(condition.PageSize))
	})
}

func (model *Model) Find(ctx context.Context, condition *FindCondition, scopes ...func(*gorm.DB) *gorm.DB) ([]*schema.GithubUser, error) {
	var dest []*schema.GithubUser
	tx := model.GetDB().WithContext(ctx)
	result := tx.Scopes(scopes...).Scopes(condition.Build()...).Find(&dest)
	return dest, result.Error
}
