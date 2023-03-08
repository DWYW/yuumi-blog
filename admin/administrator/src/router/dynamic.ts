export const DYNAMIC_ROUTES = [{
  path: "system/administrator/list",
  name: "AdministratorList",
  component: () => import("../pages/administrator/list/index.vue"),
  meta: {}
}, {
  path: "system/navmenu/tree",
  name: "NavMenuTree",
  component: () => import("../pages/navmenu/tree/index.vue"),
  meta: {}
}, {
  path: "system/permission/list",
  name: "PermissionList",
  component: () => import("../pages/permission/list/index.vue"),
  meta: {}
}, {
  path: "system/role/list",
  name: "RoleList",
  component: () => import("../pages/role/list/index.vue"),
  meta: {}
}]
