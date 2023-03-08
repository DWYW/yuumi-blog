package visitor

import (
	"context"
	"fmt"
	"yuumi/app/service/article/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func (model *Model) UpdateWithID(ctx context.Context, id int64, v interface{}) (*schema.Visitor, error) {
	visitor := schema.Visitor{Model: gorm.Model{ID: uint(id)}}

	result := model.GetDB().WithContext(ctx).Model(&visitor).Updates(v)
	if result.Error != nil {
		return nil, result.Error
	}

	return model.FindOne(ctx, &FindCondition{ID: id})
}

func (model *Model) UpdateWithIDAndUpdateGithubUser(ctx context.Context, id int64, v interface{}, u interface{}) (*schema.Visitor, error) {
	err := model.GetDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		visitor := schema.Visitor{Model: gorm.Model{ID: uint(id)}}
		if err := tx.Model(&visitor).Updates(v).Error; err != nil {
			return nil
		}

		user := schema.GithubUser{}
		if err := tx.Model(&user).Where(
			fmt.Sprintf("`%s` = ?", user.GetFeilds().VisitorID),
			visitor.ID,
		).Updates(u).Error; err != nil {
			return nil
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return model.FindOne(ctx, &FindCondition{ID: id})
}
