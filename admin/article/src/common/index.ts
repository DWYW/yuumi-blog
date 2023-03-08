import StorageHelper from "./storage"
import Validator from "./validator"

const env = (process.env.NODE_ENV).toUpperCase()

export const $storage = {
  local: new StorageHelper({
    type: "localStorage",
    keyPrefix: `${env}_${process.env.VUE_APP_STORAGE_PREFIX}`
  }),
  session: new StorageHelper({
    type: "sessionStorage",
    keyPrefix: `${env}_${process.env.VUE_APP_STORAGE_PREFIX}`
  })
}

export * as $day from "./day"
export * as $helper from "./helper"
export const $validator = new Validator()
