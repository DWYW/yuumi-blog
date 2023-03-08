import { AdministratorServer, RequestOption } from "./instance"

export function getList(params?: { [key: string]: number|string }, options?: RequestOption) {
  return AdministratorServer.get("/permissions/list", params, options)
}

export interface PermissionData {
  name: string
  rpc_method: string
}

export function create(data: PermissionData, options?: RequestOption) {
  return AdministratorServer.post("/permission", data, options)
}

export function deleteWithId(id: string, options?: RequestOption) {
  return AdministratorServer.delete(`/permission/${id}`, null, options)
}

export function updateWithId(id: string, data: PermissionData, options?: RequestOption) {
  return AdministratorServer.put(`/permission/${id}`, data, options)
}
