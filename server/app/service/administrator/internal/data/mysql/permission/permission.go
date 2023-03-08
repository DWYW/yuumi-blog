package permission

import (
	"yuumi/app/service/administrator/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func New(db *gorm.DB) *Model {
	return &Model{
		client:    db,
		TableName: "permissions",
		Feilds:    schema.Permission{}.GetFeilds(),
	}
}

type Model struct {
	client    *gorm.DB
	TableName string
	Feilds    *schema.PermissionFields
}

func (m Model) GetDB() *gorm.DB {
	return m.client
}
