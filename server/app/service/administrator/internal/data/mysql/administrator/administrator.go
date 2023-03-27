package administrator

import (
	"yuumi/app/service/administrator/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func New(db *gorm.DB) *Model {
	v := schema.Administrator{}
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
	Fields    *schema.AdministratorFields
	Relations *schema.AdministratorRelations
}

func (m Model) GetDB() *gorm.DB {
	return m.client
}
