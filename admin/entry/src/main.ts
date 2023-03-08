import { createApp } from "vue"
import App from "./App.vue"
import router from "./router"
import { store, key } from "./store"
import i18n from "./i18n"
import yuumi from "yuumi-ui-vue"

export const app = createApp(App)

app.use(router)
app.use(store, key)
app.use(i18n)
app.use(yuumi)

app.mount("#app")
