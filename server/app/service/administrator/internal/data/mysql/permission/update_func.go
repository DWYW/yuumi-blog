package permission

import (
	"context"
	"yuumi/app/service/administrator/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func (model *Model) UpdateWithID(ctx context.Context, id int64, v interface{}) (*schema.Permission, error) {
	permission := schema.Permission{Model: gorm.Model{ID: uint(id)}}

	result := model.GetDB().WithContext(ctx).Model(&permission).Updates(v)
	if result.Error != nil {
		return nil, result.Error
	}

	return model.FindOne(ctx, &FindCondition{ID: id})
}
