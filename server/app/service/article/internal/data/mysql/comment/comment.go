package comment

import (
	"yuumi/app/service/article/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func New(db *gorm.DB) *Model {
	v := schema.Comment{}
	return &Model{
		client:    db,
		TableName: v.TableName(),
		Fields:    v.GetFields(),
		Relations: v.GetRelations(),
	}
}

type Model struct {
	client    *gorm.DB
	TableName string
	Fields    *schema.CommentFields
	Relations *schema.CommentRelations
}

func (m Model) GetDB() *gorm.DB {
	return m.client
}
