package githubuser

import (
	"yuumi/app/service/article/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func New(db *gorm.DB) *Model {
	return &Model{
		client:    db,
		TableName: "github_users",
		Fields:    schema.GithubUser{}.GetFields(),
	}
}

type Model struct {
	client    *gorm.DB
	TableName string
	Fields    *schema.GithubUserFields
}

func (m Model) GetDB() *gorm.DB {
	return m.client
}
