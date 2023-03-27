package navmenu

import (
	"context"
	"yuumi/app/service/administrator/internal/data/mysql/schema"
)

type FindWithRoleIDsCondition struct {
	RoleIDS       []int64
	FindCondition *FindCondition
}

func (model *Model) FindWithRoleIDs(ctx context.Context, condition *FindWithRoleIDsCondition) ([]*schema.NavMenu, error) {
	var dest []*schema.NavMenu

	roles := make([]schema.Role, len(condition.RoleIDS))
	for i, item := range condition.RoleIDS {
		roles[i].ID = uint(item)
	}

	tx := model.GetDB().WithContext(ctx)
	tx = tx.Model(roles)
	if condition.FindCondition != nil {
		tx = tx.Scopes(condition.FindCondition.Build()...)
	}

	relations := schema.Role{}.GetRelations()
	err := tx.Association(relations.NavMenus).Find(&dest)
	return dest, err
}
