package administrator

import (
	"yuumi/app/service/administrator/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func New(db *gorm.DB) *Model {
	return &Model{
		client:    db,
		TableName: "administrators",
		Feilds:    schema.Administrator{}.GetFeilds(),
	}
}

type Model struct {
	client    *gorm.DB
	TableName string
	Feilds    *schema.AdministratorFields
}

func (m Model) GetDB() *gorm.DB {
	return m.client
}
