import { AdministratorServer, RequestOption } from "./instance"

export function getMineNavMenus(params?: { [key: string]: number|string }, options?: RequestOption) {
  return AdministratorServer.get("/navmenus/mine", params, options)
}
