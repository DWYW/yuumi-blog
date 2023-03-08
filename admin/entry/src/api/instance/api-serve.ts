import { YuumiRequest } from "yuumi-request"
import { bindRequestInterceptor, bindResponseInterceptor } from "./interceptor"

export const AdministratorServer = new YuumiRequest({
  baseURI: process.env.VUE_APP_SERVICE_ADMINISTRATOR,
  headers: {
    "Content-Type": "application/json"
  }
})
bindRequestInterceptor(AdministratorServer)
bindResponseInterceptor(AdministratorServer)

export const PassportServer = new YuumiRequest({
  baseURI: process.env.VUE_APP_SERVICE_PASSPORT,
  headers: {
    "Content-Type": "application/json"
  }
})
bindRequestInterceptor(PassportServer)
bindResponseInterceptor(PassportServer)
