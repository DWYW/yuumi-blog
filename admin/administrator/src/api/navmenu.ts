import { AdministratorServer, RequestOption } from "./instance"

export function getNavMenus(params?: { [key: string]: number|string }, options?: RequestOption) {
  return AdministratorServer.get("/navmenus", params, options)
}

export function bindRoles(id: string, role_ids: string[], options?: RequestOption) {
  return AdministratorServer.post(`/navmenu/${id}/roles/bind`, { role_ids }, options)
}

export function unbindRoles(id: string, role_ids: string[], options?: RequestOption) {
  return AdministratorServer.post(`/navmenu/${id}/roles/unbind`, { role_ids }, options)
}

export interface MenuData {
  name: string
  link_url?: string
  actived_rule?: string
  icon?: string
  weight?: string
  parent_id?: string
}

export function create(data: MenuData, options?: RequestOption) {
  return AdministratorServer.post("/navmenu", data, options)
}

export function updateWithId(id: string, data: MenuData, options?: RequestOption) {
  return AdministratorServer.put(`/navmenu/${id}`, data, options)
}

export function deleteWithId(id: string, options?: RequestOption) {
  return AdministratorServer.delete(`/navmenu/${id}`, null, options)
}
