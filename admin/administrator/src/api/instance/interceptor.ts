import { $storage } from "@/common"
import { ACCESS_TOKEN } from "@/constant/keys"
import { createMessage } from "yuumi-ui-vue"
import { YuumiRequest } from "yuumi-request"
import i18n from "@/i18n"
import { siginout } from "@/micro-application"
const { t } = i18n.global

export function bindRequestInterceptor(instance: YuumiRequest) {
  instance.interceptor.request((config: any) => {
    const token = $storage.local.getItem(ACCESS_TOKEN)

    const headers = config.headers || (config.headers = {})
    if (token) {
      headers.Authorization = token
    }

    const params = config.params || (config.params = {})
    params._version = process.env.__APP_VERSION__
    params._micro_name = process.env.__APP_MICRO_NAME__
    params._platform = "desktop"

    if (/get/i.test(config.method)) {
      params._time = Date.now()
    }

    return config
  }, (err: Error) => Promise.reject(err))
}

const CodeMsg: {[x: string]: string} = {
  0: t("MSG.REFUSED"),
  401: t("MSG.EXPIRED"),
  403: t("MSG.PERMISSION_DENIED")
}

const ErrorMsg: {[x: string]: string} = {
  "request:timeout": t("MSG.TIMEOUT"),
  "request:aborted": ""
}

export function bindResponseInterceptor(instance: YuumiRequest) {
  instance.interceptor.response(({ response }) => {
    if (typeof response === "string") {
      try {
        response = JSON.parse(response)
      } catch (err) {}
    }

    return response
  }, (err: any) => {
    const errcode = err.status || 500
    let message = ""

    try {
      const response = JSON.parse(err.response)
      message = CodeMsg[errcode] || response.message
    } catch (e) {
      if (err.status !== undefined) {
        message = CodeMsg[errcode]
      }
    }

    message = ErrorMsg[message] || message

    if (message) createMessage({ theme: "error", message, duration: 3000 })

    if (errcode === 401) {
      siginout()
    }

    return Promise.reject(err)
  })
}
