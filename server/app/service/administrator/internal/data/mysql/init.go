package mysql

import (
	"yuumi/app/service/administrator/internal/data/mysql/schema"

	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	client = db
	db.AutoMigrate(&schema.Administrator{})
	db.AutoMigrate(&schema.Role{})
	db.AutoMigrate(&schema.Permission{})
	db.AutoMigrate(&schema.AdministratorRole{})
	db.AutoMigrate(&schema.RolePermission{})
	db.AutoMigrate(&schema.NavMenu{})
	db.AutoMigrate(&schema.RoleNavMenu{})
}
