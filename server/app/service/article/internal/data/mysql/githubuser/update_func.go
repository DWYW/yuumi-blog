package githubuser

import (
	"context"
	"yuumi/app/service/article/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func (model *Model) UpdateWithID(ctx context.Context, id int64, v interface{}) (*schema.GithubUser, error) {
	user := schema.GithubUser{Model: gorm.Model{ID: uint(id)}}

	result := model.GetDB().WithContext(ctx).Model(&user).Updates(v)
	if result.Error != nil {
		return nil, result.Error
	}

	return model.FindOne(ctx, &FindCondition{ID: id})
}
