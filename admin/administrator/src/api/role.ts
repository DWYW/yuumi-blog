import { AdministratorServer, RequestOption } from "./instance"

export interface RoleData {
  name: string
  type: string
  parent_id?: string
}

export function create(data: RoleData, options?: RequestOption) {
  return AdministratorServer.post("/role", data, options)
}

export function deleteWithId(id: string, options?: RequestOption) {
  return AdministratorServer.delete(`/role/${id}`, null, options)
}

export function updateWithId(id: string, data: RoleData, options?: RequestOption) {
  return AdministratorServer.put(`/role/${id}`, data, options)
}

export function getList(params?: { [key: string]: number|string }, options?: RequestOption) {
  return AdministratorServer.get("/roles/list", params, options)
}

export function getRoles(params?: { [key: string]: number|string }, options?: RequestOption) {
  return AdministratorServer.get("/roles", params, options)
}

export function getRolesWithAdministratorId(administratorId: string, params?: { [key: string]: number|string }, options?: RequestOption) {
  return AdministratorServer.get(`/roles/administrator/${administratorId}`, params, options)
}

export function appendPermissions(roleId: string, permisiionIds: string[], options?: RequestOption) {
  return AdministratorServer.post(`/role/${roleId}/permissions/append`, { permission_ids: permisiionIds }, options)
}

export function deletePermissions(roleId: string, permisiionIds: string[], options?: RequestOption) {
  return AdministratorServer.post(`/role/${roleId}/permissions/delete`, { permission_ids: permisiionIds }, options)
}
