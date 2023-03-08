package visitor

import (
	"context"
	"fmt"
	"yuumi/app/service/article/internal/data/mysql/schema"

	"gorm.io/gorm"
)

type DeleteCondition struct {
	ID int64
}

func (condition *DeleteCondition) Build() []func(*gorm.DB) *gorm.DB {
	var scopes = make([]func(*gorm.DB) *gorm.DB, 0)
	fields := schema.Visitor{}.GetFeilds()

	if condition.ID > 0 {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("`%s` = ?", fields.ID), condition.ID)
		})
	}

	return scopes
}

func (model *Model) DeleteWithID(ctx context.Context, id int64) (*schema.Visitor, error) {
	dest := schema.Visitor{}
	condition := DeleteCondition{ID: id}
	result := model.GetDB().WithContext(ctx).Scopes(condition.Build()...).Delete(&dest)
	return &dest, result.Error
}
