interface DateAttrs {
  year: number; // 年
  month: number; // 月
  date: number; // 日
  hour: number; // 时
  minute: number; // 分
  second: number; // 秒
  day: number; // 星期
  monthDays: number; // 本月的天数
  stamp: number; // 时间戳
}

function numberStringify(number: number): string {
  return number < 10 ? `0${number}` : number.toString()
}

function dateParse(input: Date): DateAttrs|null {
  if ((input instanceof Date) === false) return null

  const year = input.getFullYear()
  const month = input.getMonth()
  const date = input.getDate()
  const hour = input.getHours()
  const minute = input.getMinutes()
  const second = input.getSeconds()
  const day = input.getDay()
  const monthDays = new Date(year, month + 1, 0, 0, 0, 0).getDate()
  const stamp = input.getTime()

  return {
    year,
    month,
    date,
    hour,
    minute,
    second,
    day,
    monthDays,
    stamp
  }
}

export const useDay =() => {
  function format(input: Date, rule = "YYYY/MM/DD"): string {
    const attrs = dateParse(input)
    if (!attrs) return ""

    let _string = rule.replace(/(^\s+)|(\s+$)/g, "")

    const formats: [RegExp, number|string][] = [
      [/Y{4,}/, attrs.year],
      [/Y{2,}/, attrs.year.toString().substr(2)],
      [/M{2,}/, numberStringify(attrs.month + 1)],
      [/M/, attrs.month + 1],
      [/D{2,}/, numberStringify(attrs.date)],
      [/D/, attrs.date],
      [/h{2,}/, numberStringify(attrs.hour)],
      [/h/, attrs.hour],
      [/m{2,}/, numberStringify(attrs.minute)],
      [/m/, attrs.minute],
      [/s{2,}/, numberStringify(attrs.second)],
      [/s/, attrs.second]
    ]

    formats.forEach(([reg, value]) => {
      _string = _string.replace(reg, value.toString())
    })

    return _string
  }

  function formatFromDateString(input: string, rule?: string) {
    const val = new Date(input)
    if (val.toString() === "Invalid Date") return ""
    return format(val, rule)
  }

  function semanticFromDateString(input: string, rule?: string) {
    const val = new Date(input)
    if (val.toString() === "Invalid Date") return ""

    const days = Math.floor((Date.now() - val.getTime()) / (24 * 3600 * 1000))
    if (days >= 7) {
      return formatFromDateString(input, rule)
    }

    if (days >= 1) {
      return `${days}天前`
    }

    const hours = Math.floor((Date.now() - val.getTime()) / (3600 * 1000))
    if (hours >= 1) {
      return `${hours}小时前`
    }

    const minutes = Math.floor((Date.now() - val.getTime()) / (60 * 1000))
    if (minutes >= 1) {
      return `${minutes}分钟前`
    }

    return "刚刚"
  }

  return {
    format,
    formatFromDateString,
    semanticFromDateString
  }
}