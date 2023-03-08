package visitor

import (
	"yuumi/app/service/article/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func New(db *gorm.DB) *Model {
	return &Model{
		client:    db,
		TableName: "visitors",
		Feilds:    schema.Visitor{}.GetFeilds(),
	}
}

type Model struct {
	client    *gorm.DB
	TableName string
	Feilds    *schema.VisitorFields
}

func (m Model) GetDB() *gorm.DB {
	return m.client
}
