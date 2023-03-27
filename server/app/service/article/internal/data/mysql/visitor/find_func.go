package visitor

import (
	"context"
	"fmt"
	"yuumi/app/service/article/internal/data/mysql/githubuser"
	"yuumi/app/service/article/internal/data/mysql/schema"
)

func (model *Model) FindOneWithGithubID(ctx context.Context, id int64) (*schema.Visitor, error) {
	tx := model.GetDB().WithContext(ctx)
	visitor := schema.Visitor{}
	userModel := githubuser.New(tx)
	subQuery := tx.Select(userModel.Fields.ID).Where(fmt.Sprintf("`%s` = ?", userModel.Fields.GithubID), id).Table(userModel.TableName)
	result := tx.Where(fmt.Sprintf("`%s` = (?)", visitor.GetFields().ID), subQuery).First(&visitor)
	return &visitor, result.Error
}
