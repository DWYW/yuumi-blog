export default class Validator {
  isDefine(input: any) {
    return input !== null && input !== undefined
  }

  isString(input: any) {
    return typeof input === "string"
  }

  isBoolean(input: any) {
    return typeof input === "boolean"
  }

  isNumber(input: any) {
    return typeof input === "number"
  }

  isNaN(input: any) {
    return Number.isNaN ? Number.isNaN(input) : this.isNumber(input) && isNaN(input)
  }

  isBigint(input: any) {
    return typeof input === "bigint"
  }

  isSymbol(input: any) {
    return typeof input === "symbol"
  }

  isFunction(input: any) {
    return typeof input === "function"
  }

  isObject(input: any) {
    return typeof input === "object" && input === Object(input)
  }

  isArray(input: any) {
    return Object.prototype.toString.call(input) === "[object Array]"
  }

  isDate(input: any) {
    return Object.prototype.toString.call(input) === "[object Date]"
  }

  isRegExp(input: any) {
    return Object.prototype.toString.call(input) === "[object RegExp]"
  }

  isPromise(input: any) {
    return Object.prototype.toString.call(input) === "[object Promise]"
  }

  isFormData(input: any) {
    return Object.prototype.toString.call(input) === "[object FormData]"
  }

  isEmpty(input: any) {
    if (!this.isDefine(input)) return true
    if (this.isString(input)) return /^\s*$/.test(input)
    if (this.isNaN(input)) return true
    if (this.isArray(input)) return input.length === 0
    if (this.isDate(input)) return input.toString() === "Invalid Date"

    if (this.isObject(input)) {
      for (const key in input) {
        if (key) return false
      }

      return true
    }

    return false
  }

  isMobileNumber(value: string): boolean {
    return /^1[3-9]\d{9}$/.test(value)
  }

  isTelNumber(value: string): boolean {
    return /^0\d{2,3}-?[2-9]\d{6,7}$/.test(value)
  }

  isPhoneNumber(value: string): boolean {
    return this.isMobileNumber(value) || this.isTelNumber(value)
  }

  __PROVINCES__ = [
    11, 12, 13, 14, 15,
    21, 22, 23,
    31, 32, 33, 34, 35, 36, 37,
    41, 42, 43, 44, 45, 46,
    50, 51, 52, 53, 54,
    61, 62, 63, 64, 65,
    71,
    81, 82
  ]

  isOldIdentityCardNumber(input: string): boolean {
    if (!/^\d{15}$/.test(input)) return false
    if (!~this.__PROVINCES__.indexOf(Number(input.slice(0, 2)))) return false
    if (Number(input.slice(8, 10)) > 12) return false
    if (Number(input.slice(10, 12)) > 31) return false
    return true
  }

  isIdentityCardNumber(input: string): boolean {
    if (!/^\d{17}(\d|X)$/.test(input)) return false
    if (!~this.__PROVINCES__.indexOf(Number(input.slice(0, 2)))) return false
    if (Number(input.slice(10, 12)) > 12) return false
    if (Number(input.slice(12, 14)) > 31) return false

    const max = input.length - 1
    let sumary = 0

    for (let i = 0; i < max; i++) {
      const Wi = Math.pow(2, max - i) % 11
      const Ai = Number(input.slice(i, i + 1))
      sumary += Ai * Wi
    }

    const remainder = sumary % 11
    const validationCode = ["1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"]
    return validationCode[remainder] === input.slice(max)
  }
}
