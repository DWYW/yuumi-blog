package visitor

import (
	"context"
	"fmt"
	"yuumi/app/service/article/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func (model *Model) UpdateWithIDAndUpdateGithubUser(ctx context.Context, id int64, user interface{}, githubUser interface{}) (*schema.Visitor, error) {
	err := model.GetDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		visitor := schema.Visitor{Model: gorm.Model{ID: uint(id)}}
		if err := tx.Model(&visitor).Updates(user).Error; err != nil {
			return nil
		}

		user := schema.GithubUser{}
		if err := tx.Model(&user).Where(
			fmt.Sprintf("`%s` = ?", user.GetFields().VisitorID),
			visitor.ID,
		).Updates(githubUser).Error; err != nil {
			return nil
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return model.FindOne(ctx, &FindCondition{Where: &FindConditionWhere{ID: id}})
}
