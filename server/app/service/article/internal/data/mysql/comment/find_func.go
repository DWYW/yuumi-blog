package comment

import (
	"context"
	"fmt"
	"yuumi/app/service/article/internal/data/mysql/schema"

	"gorm.io/gorm"
)

type FindCondition struct {
	ID        int64
	ArticleID int64
	CreatorID int64
	Sort      *FindConditionSort
	Preload   *FindConditionPreload
}

type FindConditionSort struct {
	ID int64
}

type FindConditionPreload struct {
	Article bool
	Creator bool
}

func (condition *FindCondition) Build() []func(*gorm.DB) *gorm.DB {
	var scopes = make([]func(*gorm.DB) *gorm.DB, 0)
	fields := schema.Comment{}.GetFeilds()

	if condition.ID > 0 {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("`%s` = ?", fields.ID), condition.ID)
		})
	}

	if condition.ArticleID > 0 {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("`%s` = ?", fields.ArticleID), condition.ArticleID)
		})
	}

	if condition.CreatorID > 0 {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("`%s` = ?", fields.CreatorID), condition.CreatorID)
		})
	}

	if condition.Sort != nil {
		scopes = append(scopes, condition.BuildOrder()...)
	}

	return scopes
}

func (condition *FindCondition) BuildOrder() []func(*gorm.DB) *gorm.DB {
	var scopes = make([]func(*gorm.DB) *gorm.DB, 0)
	fields := schema.Comment{}.GetFeilds()

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

func (model *Model) FindByID(ctx context.Context, id int64) (*schema.Comment, error) {
	return model.FindOne(ctx, &FindCondition{ID: id})
}

func (model *Model) FindOne(ctx context.Context, condition *FindCondition) (*schema.Comment, error) {
	dest := schema.Comment{}
	result := model.GetDB().WithContext(ctx).Scopes(condition.Build()...).First(&dest)
	return &dest, result.Error
}

type FindListCondition struct {
	Page          int64
	PageSize      int64
	FindCondition *FindCondition
}

func (model *Model) FindList(ctx context.Context, condition *FindListCondition) ([]*schema.Comment, error) {
	skipNum := (condition.Page - 1) * condition.PageSize
	return model.Find(ctx, condition.FindCondition, func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(skipNum)).Limit(int(condition.PageSize))
	})
}

func (model *Model) Find(ctx context.Context, condition *FindCondition, scopes ...func(*gorm.DB) *gorm.DB) ([]*schema.Comment, error) {
	var dest []*schema.Comment
	tx := model.GetDB().WithContext(ctx)

	if condition.Preload != nil {
		relations := schema.Comment{}.GetRelations()

		if condition.Preload.Creator {
			tx = tx.Preload(relations.Creator)
		}

		if condition.Preload.Article {
			tx = tx.Preload(relations.Article)
		}
	}

	result := tx.Scopes(scopes...).Scopes(condition.Build()...).Find(&dest)
	return dest, result.Error
}
