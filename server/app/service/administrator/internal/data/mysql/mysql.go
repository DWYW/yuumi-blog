package mysql

import (
	"yuumi/app/service/administrator/internal/data/mysql/administrator"
	"yuumi/app/service/administrator/internal/data/mysql/navmenu"
	"yuumi/app/service/administrator/internal/data/mysql/permission"
	"yuumi/app/service/administrator/internal/data/mysql/role"
	"yuumi/app/service/administrator/internal/data/mysql/transaction"

	"gorm.io/gorm"
)

func GetClient() *gorm.DB {
	return client
}

func getDefaultDB(databases ...*gorm.DB) *gorm.DB {
	var db *gorm.DB
	for _, item := range databases {
		if item != nil {
			db = item
			break
		}
	}

	if db == nil {
		db = client
	}

	return db
}

func GetAdministrator(databases ...*gorm.DB) *administrator.Model {
	return administrator.New(getDefaultDB(databases...))
}

func GetRole(databases ...*gorm.DB) *role.Model {
	return role.New(getDefaultDB(databases...))
}

func GetPermission(databases ...*gorm.DB) *permission.Model {
	return permission.New(getDefaultDB(databases...))
}

func GetNavmenu(databases ...*gorm.DB) *navmenu.Model {
	return navmenu.New(getDefaultDB(databases...))
}

func GetTransaction(databases ...*gorm.DB) *transaction.Transaction {
	db := getDefaultDB(databases...)
	return &transaction.Transaction{
		DB:               db,
		GetAdministrator: GetAdministrator,
		GetRole:          GetRole,
		GetPermission:    GetPermission,
		GetNavmenu:       GetNavmenu,
	}
}
