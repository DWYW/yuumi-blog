import { PassportServer, RequestOption } from "./instance"

export function signinWithPassword(data: any, options?: RequestOption) {
  return PassportServer.post("/authorize/administrator/password", data, options)
}
