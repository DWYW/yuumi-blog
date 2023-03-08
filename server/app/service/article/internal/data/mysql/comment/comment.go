package comment

import (
	"yuumi/app/service/article/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func New(db *gorm.DB) *Model {
	return &Model{
		client:    db,
		TableName: "comments",
		Feilds:    schema.Comment{}.GetFeilds(),
	}
}

type Model struct {
	client    *gorm.DB
	TableName string
	Feilds    *schema.CommentFields
}

func (m Model) GetDB() *gorm.DB {
	return m.client
}
