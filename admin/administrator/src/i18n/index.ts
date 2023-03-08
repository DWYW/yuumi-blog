import { createI18n } from "vue-i18n"
import * as en from "./en"
import * as zh from "./zh"

export default createI18n({
  legacy: false,
  locale: "zh",
  messages: {
    en,
    zh
  }
})
