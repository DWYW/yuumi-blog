import i18n from "@/i18n"
const { t } = i18n.global

const TEXT = {
  today: t("DAY.TODAY"),
  yesterday: t("DAY.YESTERDAY"),
  tomorrow: t("DAY.TOMORROW")
}

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

export function stamp2ms(stamp: number|string): number {
  if (typeof stamp !== "number" && typeof stamp !== "string") return 0

  stamp = stamp.toString().slice(0, 13)
  while (stamp.length < 13) {
    stamp += "0"
  }

  return Number(stamp)
}

export function dateParse(input: Date): DateAttrs|null {
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

export function formatTimeFromDateString(input: string, rule?: string): string {
  const val = new Date(input)
  if (val.toString() === "Invalid Date") return ""
  return formatTime(val, rule)
}

export function formatTime(input: Date, rule = "YYYY/MM/DD hh:mm:ss"): string {
  return format(input, rule)
}

export function format(input: Date, rule = "YYYY/MM/DD"): string {
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

export function stampFormat(stamp: number|string, rule?: string): string {
  if (typeof stamp !== "number" && typeof stamp !== "string") return ""

  const date = new Date(stamp2ms(stamp))
  return format(date, rule)
}

export function semantic(date: Date, rule?: string) {
  if ((date instanceof Date) === false) return ""

  const parse = dateParse(new Date()) as DateAttrs
  const todayStamp = new Date(parse.year, parse.month, parse.date, 0, 0, 0).getTime()

  const days = (date.getTime() - todayStamp) / (24 * 3600 * 1000)

  if (days >= 1 && days < 2) return TEXT.tomorrow
  if (days >= 0 && days < 1) return TEXT.today
  if (days >= -1 && days < 0) return TEXT.yesterday

  return format(date, rule)
}

export function stamp2semantic(stamp: number|string, rule?: string) {
  const date = new Date(stamp2ms(stamp))
  return semantic(date, rule)
}

export function seconds2semantic(seconds: number): string {
  const arr = ["--", "--"]

  if (seconds >= 0) {
    seconds = Math.ceil(seconds)

    const second = seconds % 60
    arr[1] = numberStringify(second)

    const minute = Math.floor(seconds / 60) % 60
    arr[0] = numberStringify(minute)

    const hour = Math.floor(seconds / 60 / 60)

    if (hour) {
      arr.unshift(numberStringify(hour))
    }
  }

  return arr.join(":")
}

export function string2date(input: string, rule?: number[][]): Date {
  const [
    [yearStart, yearEnd],
    [monthStart, monthEnd],
    [dayStart, dayEnd]
  ] = rule || [[0, 4], [4, 6], [6, 8]]

  return new Date(
    Number(input.slice(yearStart, yearEnd)),
    Number(input.slice(monthStart, monthEnd)) - 1,
    Number(input.slice(dayStart, dayEnd)),
    0,
    0,
    0
  )
}

export function stamp2date(value: number) {
  return new Date(stamp2ms(value))
}
