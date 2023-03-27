package permission

import (
	"yuumi/app/service/administrator/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func New(db *gorm.DB) *Model {
	v := schema.Permission{}
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
	Fields    *schema.PermissionFields
	Relations *schema.PermissionRelations
}

func (m Model) GetDB() *gorm.DB {
	return m.client
}
