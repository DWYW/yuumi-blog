package transform

import (
	v1 "yuumi/api/service/administrator/v1"
	"yuumi/app/service/administrator/internal/data/mysql/schema"
)

func TransformAdministratorData(v *schema.Administrator) *v1.AdministratorData {
	var roles = make([]*v1.AdministratorRoleData, len(v.Roles))
	for i, role := range v.Roles {
		roles[i] = TransformAdministratorRoleData(role)
	}

	return &v1.AdministratorData{
		Id:        int64(v.ID),
		Name:      v.Name,
		CreatedAt: v.CreatedAt.String(),
		UpdatedAt: v.UpdatedAt.String(),
		Roles:     roles,
	}
}

func TransformAdministratorsData(v []*schema.Administrator) []*v1.AdministratorData {
	var administrators = make([]*v1.AdministratorData, len(v))
	for i, administrator := range v {
		administrators[i] = TransformAdministratorData(administrator)
	}
	return administrators
}

func TransformAdministratorRoleData(v *schema.Role) *v1.AdministratorRoleData {
	return &v1.AdministratorRoleData{
		Id:        int64(v.ID),
		Name:      v.Name,
		CreatedAt: v.CreatedAt.String(),
		UpdatedAt: v.UpdatedAt.String(),
	}
}

func TransformRolesData(v []*schema.Role) []*v1.RoleData {
	var roles = make([]*v1.RoleData, len(v))
	for i, role := range v {
		roles[i] = TransformRoleData(role)
	}
	return roles
}

func TransformRoleData(v *schema.Role) *v1.RoleData {
	return &v1.RoleData{
		Id:        int64(v.ID),
		Name:      v.Name,
		Type:      v.Type,
		ParentId:  int64(v.ParentID),
		CreatedAt: v.CreatedAt.String(),
		UpdatedAt: v.UpdatedAt.String(),
	}
}

func TransformPermissionsData(v []*schema.Permission) []*v1.PermissionData {
	var permissions = make([]*v1.PermissionData, len(v))
	for i, permission := range v {
		permissions[i] = TransformPermissionData(permission)
	}
	return permissions
}

func TransformPermissionData(v *schema.Permission) *v1.PermissionData {
	var roles = make([]*v1.AdministratorRoleData, len(v.Roles))
	for i, role := range v.Roles {
		roles[i] = TransformAdministratorRoleData(role)
	}

	return &v1.PermissionData{
		Id:        int64(v.ID),
		Name:      v.Name,
		RpcMethod: v.RpcMethod,
		CreatedAt: v.CreatedAt.String(),
		UpdatedAt: v.UpdatedAt.String(),
		Roles:     roles,
	}
}

func TransformNavMenusData(v []*schema.NavMenu) []*v1.NavMenuData {
	var menus = make([]*v1.NavMenuData, len(v))
	for i, menu := range v {
		menus[i] = TransformNavMenuData(menu)
	}
	return menus
}

func TransformNavMenuData(v *schema.NavMenu) *v1.NavMenuData {
	return &v1.NavMenuData{
		Id:          int64(v.ID),
		CreatedAt:   v.CreatedAt.String(),
		UpdatedAt:   v.UpdatedAt.String(),
		ParentId:    int64(v.ParentID),
		Name:        v.Name,
		LinkUrl:     v.LinkUrl,
		ActivedRule: v.ActivedRule,
		Weight:      v.Weight,
		Icon:        v.Icon,
		Roles:       TransformRolesData(v.Roles),
	}
}
