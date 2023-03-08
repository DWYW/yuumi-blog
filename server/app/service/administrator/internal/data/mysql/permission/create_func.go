package permission

import (
	"context"
	"yuumi/app/service/administrator/internal/data/mysql/schema"
)

func (model *Model) CreateOne(ctx context.Context, v *schema.Permission) (*schema.Permission, error) {
	result := model.GetDB().WithContext(ctx).Create(v)
	if err := result.Error; err != nil {
		return nil, err
	}

	return model.FindOne(ctx, &FindCondition{ID: int64(v.ID)})
}
