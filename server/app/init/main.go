package main

import (
	"yuumi/app/init/task"
)

var data = []task.AdministratorData{
	{
		Name:     "root",
		Password: "password",
		Roles: []task.RoleData{
			{Name: "系统管理员", Type: "administrator:root", Children: []task.RoleData{
				{Name: "【read】系统管理员", Type: "administrator:root:read", Navs: []task.NavData{
					{Name: "系统管理", Icon: "line-setting", Weight: 0, Children: []task.NavData{
						{Name: "管理员管理", LinkUrl: "/micro/administrator/system/administrator/list", ActivedRule: "^/micro/administrator/system/administrator/list", Weight: 99},
						{Name: "角色管理", LinkUrl: "/micro/administrator/system/role/list", ActivedRule: "^/micro/administrator/system/role/list", Weight: 98},
						{Name: "权限管理", LinkUrl: "/micro/administrator/system/permission/list", ActivedRule: "^/micro/administrator/system/permission/list", Weight: 97},
						{Name: "菜单管理", LinkUrl: "/micro/administrator/system/navmenu/tree", ActivedRule: "^/micro/administrator/system/navmenu/tree", Weight: 96},
					}},
				}, Permissions: []task.PermissionData{
					{Name: "【管理员】获取详情", RpcMethod: "/interface.administrator.v1.AdministratorInterface/GetInfo"},
					{Name: "【管理员】获取分页列表", RpcMethod: "/interface.administrator.v1.AdministratorInterface/GetList"},
					{Name: "【角色】获取详情", RpcMethod: "/interface.administrator.v1.RoleInterface/GetInfo"},
					{Name: "【角色】获取全部", RpcMethod: "/interface.administrator.v1.RoleInterface/GetRoles"},
					{Name: "【角色】获取分页列表", RpcMethod: "/interface.administrator.v1.RoleInterface/GetList"},
					{Name: "【角色】根据管理员ID获取全部", RpcMethod: "/interface.administrator.v1.RoleInterface/GetRolesWithAdministratorID"},
					{Name: "【权限】获取分页列表", RpcMethod: "/interface.administrator.v1.PermissionInterface/GetList"},
					{Name: "【导航菜单】获取信息", RpcMethod: "/interface.administrator.v1.NavMenuInterface/GetInfo"},
					{Name: "【导航菜单】获取全部", RpcMethod: "/interface.administrator.v1.NavMenuInterface/GetNavMenus"},
					{Name: "【导航菜单】获取管理员自己的", RpcMethod: "/interface.administrator.v1.NavMenuInterface/GetNavMenusWithMine"},
					{Name: "【导航菜单】根据管理员ID获取全部", RpcMethod: "/interface.administrator.v1.NavMenuInterface/GetNavMenusWithAdministratorID"},
				}},
				{Name: "【write】系统管理员", Type: "administrator:root:write", Permissions: []task.PermissionData{
					{Name: "【管理员】创建", RpcMethod: "/interface.administrator.v1.AdministratorInterface/Create"},
					{Name: "【管理员】删除", RpcMethod: "/interface.administrator.v1.AdministratorInterface/Delete"},
					{Name: "【管理员】更新用户名", RpcMethod: "/interface.administrator.v1.AdministratorInterface/UpdateName"},
					{Name: "【管理员】更新用户密码", RpcMethod: "/interface.administrator.v1.AdministratorInterface/UpdatePassword"},
					{Name: "【管理员】追加角色", RpcMethod: "/interface.administrator.v1.AdministratorInterface/AppendRolesWithAdministratorID"},
					{Name: "【管理员】移除角色", RpcMethod: "/interface.administrator.v1.AdministratorInterface/DeleteRolesWithAdministratorID"},
					{Name: "【角色】创建", RpcMethod: "/interface.administrator.v1.RoleInterface/Create"},
					{Name: "【角色】删除", RpcMethod: "/interface.administrator.v1.RoleInterface/Delete"},
					{Name: "【角色】更新", RpcMethod: "/interface.administrator.v1.RoleInterface/Update"},
					{Name: "【角色】追加权限", RpcMethod: "/interface.administrator.v1.RoleInterface/AppendPermissionsWithRoleID"},
					{Name: "【角色】移除权限", RpcMethod: "/interface.administrator.v1.RoleInterface/DeletePermissionsWithRoleID"},
					{Name: "【权限】创建", RpcMethod: "/interface.administrator.v1.PermissionInterface/Create"},
					{Name: "【权限】删除", RpcMethod: "/interface.administrator.v1.PermissionInterface/Delete"},
					{Name: "【权限】更改", RpcMethod: "/interface.administrator.v1.PermissionInterface/Update"},
					{Name: "【导航菜单】创建", RpcMethod: "/interface.administrator.v1.NavMenuInterface/Create"},
					{Name: "【导航菜单】删除", RpcMethod: "/interface.administrator.v1.NavMenuInterface/Delete"},
					{Name: "【导航菜单】更新", RpcMethod: "/interface.administrator.v1.NavMenuInterface/Update"},
					{Name: "【导航菜单】绑定角色", RpcMethod: "/interface.administrator.v1.NavMenuInterface/BindWithRoleIDs"},
					{Name: "【导航菜单】解绑角色", RpcMethod: "/interface.administrator.v1.NavMenuInterface/UnbindWithRoleIDs"},
				}},
			}},
			{Name: "普通管理员", Type: "administrator:general", Children: []task.RoleData{
				{Name: "【read】普通管理员", Type: "administrator:general:read", Permissions: []task.PermissionData{
					{Name: "【导航菜单】获取自己全部", RpcMethod: "/interface.administrator.v1.NavMenuInterface/GetNavMenusWithMine"},
				}},
				{Name: "【write】普通管理员", Type: "administrator:general:write", Permissions: []task.PermissionData{}},
			}},
			{Name: "文章管理员", Type: "administrator:article", Children: []task.RoleData{
				{Name: "【read】文章管理员", Type: "administrator:article:read", Navs: []task.NavData{
					{Name: "文章管理", Icon: "line-catalog", Weight: 0, LinkUrl: "/micro/article/system/article/list", ActivedRule: "^/micro/article/system/article"},
				}, Permissions: []task.PermissionData{}},
				{Name: "【write】文章管理员", Type: "administrator:article:write", Permissions: []task.PermissionData{
					{Name: "【文章】创建", RpcMethod: "/interface.article.v1.ArticleInterface/Create"},
					{Name: "【文章】删除", RpcMethod: "/interface.article.v1.ArticleInterface/Delete"},
					{Name: "【文章】更新", RpcMethod: "/interface.article.v1.ArticleInterface/Update"},
				}},
			}},
		},
	},
}

func main() {
	task.Task{
		// administrator服务配置
		Host:           "127.0.0.1",
		Port:           9002,
		Administrators: data,
	}.Run()
}
