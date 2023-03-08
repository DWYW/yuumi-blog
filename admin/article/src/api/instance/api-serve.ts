import { YuumiRequest } from "yuumi-request"
import { bindRequestInterceptor, bindResponseInterceptor } from "./interceptor"

export const ArticleServer = new YuumiRequest({
  baseURI: process.env.VUE_APP_SERVICE_ARTICLE,
  headers: {
    "Content-Type": "application/json"
  }
})

bindRequestInterceptor(ArticleServer)
bindResponseInterceptor(ArticleServer)
