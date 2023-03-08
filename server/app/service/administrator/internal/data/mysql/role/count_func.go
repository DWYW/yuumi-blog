package role

import (
	"context"
	"fmt"
	"yuumi/app/service/administrator/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func (model *Model) Count(ctx context.Context, condition *FindCondition, exludeDeleted bool) (int64, error) {
	scopes := condition.Build()
	if exludeDeleted {
		fields := schema.Role{}.GetFeilds()
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("`%s` IS NULL", fields.DeletedAt))
		})
	}

	var count int64
	tx := model.GetDB().Table(model.TableName).WithContext(ctx)
	result := tx.Scopes(scopes...).Count(&count)
	return count, result.Error
}
