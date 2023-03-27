package transaction

import (
	"yuumi/app/service/administrator/internal/data/mysql/administrator"
	"yuumi/app/service/administrator/internal/data/mysql/navmenu"
	"yuumi/app/service/administrator/internal/data/mysql/permission"
	"yuumi/app/service/administrator/internal/data/mysql/role"

	"gorm.io/gorm"
)

type Transaction struct {
	DB               *gorm.DB
	GetAdministrator func(databases ...*gorm.DB) *administrator.Model
	GetRole          func(databases ...*gorm.DB) *role.Model
	GetPermission    func(databases ...*gorm.DB) *permission.Model
	GetNavmenu       func(databases ...*gorm.DB) *navmenu.Model
}
