export interface MicroApplicationConfig {
  name: string
  entry: string
  activeRule: (location: Location) => boolean
}

export const APPLICATIONS_CONFIG = [{
  name: "administrator",
  entry: process.env.VUE_APP_MICRO_ENTRY_ADMINISTRATOR
}, {
  name: "article",
  entry: process.env.VUE_APP_MICRO_ENTRY_ARTICLE
}]
