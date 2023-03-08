package navmenu

import (
	"yuumi/app/service/administrator/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func New(db *gorm.DB) *Model {
	return &Model{
		client:    db,
		TableName: "nav_menu",
		Feilds:    schema.NavMenu{}.GetFeilds(),
	}
}

type Model struct {
	client    *gorm.DB
	TableName string
	Feilds    *schema.NavMenuFields
}

func (m Model) GetDB() *gorm.DB {
	return m.client
}
