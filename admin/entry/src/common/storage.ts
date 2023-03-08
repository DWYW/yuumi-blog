interface StorageHelperOption {
  type: "localStorage"|"sessionStorage"
  hyphen?: string
  keyPrefix?: string
}

interface StorageHelperConfig extends StorageHelperOption {
  hyphen: string
  keyPrefix: string
}

export default class StorageHelper {
  config: StorageHelperConfig;
  storage: Storage;

  constructor(options: StorageHelperOption) {
    this.config = Object.assign({
      hyphen: ":",
      keyPrefix: "",
      clearBlackList: []
    }, options)

    this.storage = (window as Window)[options.type]
  }

  fullKey(key: string) {
    const { keyPrefix, hyphen } = this.config
    return `${keyPrefix}${hyphen}${key}`
  }

  getItem(key: string) {
    key = this.fullKey(key)
    let value: any = this.storage.getItem(key)
    if (value) {
      value = JSON.parse(value)
    }

    if (value?.__once__) {
      value = value.data
      this.storage.removeItem(key)
    }

    return value
  }

  setItem(key: string, data: string|number|string|object, once?: boolean) {
    key = this.fullKey(key)
    const value = JSON.stringify(once ? { __once__: true, data } : data)
    this.storage.setItem(key, value)
  }

  removeItem(key: string) {
    key = this.fullKey(key)
    this.storage.removeItem(key)
  }

  clear() {
    Object.keys(this.storage).forEach((key: string) => {
      if (!/ignore$/i.test(key) && key.indexOf(this.config.keyPrefix) >= 0) {
        this.storage.removeItem(key)
      }
    })
  }
}
