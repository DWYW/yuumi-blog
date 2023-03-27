package article

import (
	"yuumi/app/service/article/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func New(db *gorm.DB) *Model {
	v := schema.Article{}

	return &Model{
		client:    db,
		TableName: v.TableName(),
		Fields:    v.GetFields(),
	}
}

type Model struct {
	client    *gorm.DB
	TableName string
	Fields    *schema.ArticleFields
}

func (m Model) GetDB() *gorm.DB {
	return m.client
}
