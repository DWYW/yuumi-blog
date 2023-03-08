import { AdministratorServer, RequestOption } from "./instance"

export function getList(params?: { [key: string]: number|string }, options?: RequestOption) {
  return AdministratorServer.get("/administrators/list", params, options)
}

export function create(data: {
  name: string
  password: string
}, options?: RequestOption) {
  return AdministratorServer.post("/administrator", data, options)
}

export function deleteWithId(id: string, options?: RequestOption) {
  return AdministratorServer.delete(`/administrator/${id}`, null, options)
}

export function updateName(id: string, name: string, options?: RequestOption) {
  return AdministratorServer.put(`/administrator/${id}/name`, { name }, options)
}

export function appendRoles(administratorId: string, roleIds: string[], options?: RequestOption) {
  return AdministratorServer.post(`/administrator/${administratorId}/roles/append`, { role_ids: roleIds }, options)
}

export function removeRoles(administratorId: string, roleIds: string[], options?: RequestOption) {
  return AdministratorServer.post(`/administrator/${administratorId}/roles/delete`, { role_ids: roleIds }, options)
}
