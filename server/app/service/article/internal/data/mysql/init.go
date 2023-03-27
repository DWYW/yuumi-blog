package mysql

import (
	"yuumi/app/service/article/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	client = db
	db.AutoMigrate(&schema.Article{})
	db.AutoMigrate(&schema.Visitor{})
	db.AutoMigrate(&schema.GithubUser{})
	db.AutoMigrate(&schema.Comment{})
}
