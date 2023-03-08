import { createApp, App } from "vue"
import { createRouter, createWebHashHistory, Router, RouterHistory } from "vue-router"
import { routes } from "@/router"
import { store, key } from "@/store"
import { GlobalState } from "@/store/types"
import { MicroAppStateActions } from "qiankun"
import { DYNAMIC_ROUTES } from "./router/dynamic"
import i18n from "@/i18n"
import AppComponent from "@/App.vue"
import yuumi from "yuumi-ui-vue"

export let router: Router|null = null
export let app: App<Element>|null = null
export let history: RouterHistory|null = null

let microAppStateActions: MicroAppStateActions|null = null
let signoutFun: any

export function siginout() {
  signoutFun && signoutFun()
}

export function setGlobalState(state: Record<string, any>) {
  microAppStateActions?.setGlobalState(state)
}

export function render(props?: any) {
  history = createWebHashHistory(process.env.VUE_APP_PUBLIC_PATH)
  router = createRouter({
    history,
    routes: window.__POWERED_BY_QIANKUN__
      ? routes.map((route) => {
        // 自动添加路由前缀
        if (!route.path.startsWith(process.env.VUE_APP_MICRO_ROUTE_PREFIX)) {
          route.path = `${process.env.VUE_APP_MICRO_ROUTE_PREFIX}${route.path}`
        }
        return route
      })
      : routes
  })

  const { container } = props || {}
  app = createApp(AppComponent)
  app.use(router)
  app.use(store, key)
  app.use(i18n)
  app.use(yuumi)
  app.mount(container ? container.querySelector("#app") as HTMLElement : "#app")

  let registered = false
  // 路由导航守卫
  router.beforeEach((to, from, next) => {
    if (window.__POWERED_BY_QIANKUN__ && !to.path.startsWith(process.env.VUE_APP_MICRO_ROUTE_PREFIX)) {
      next(`${process.env.VUE_APP_MICRO_ROUTE_PREFIX}${to.fullPath}`)
    } else {
      if (!registered) {
        registered = true

        DYNAMIC_ROUTES.forEach((route) => {
          router?.addRoute("home", route)
        })

        next(to.fullPath)
      } else {
        next()
      }
    }
  })
}

export async function bootstrap() {
  console.log("%c%s", "color: green", `${process.env.__APP_MICRO_NAME__} bootstraped`)
}

export async function mount(props: any) {
  signoutFun = props.getSignoutFun()
  const globalState = props.getGlobalState()
  store.dispatch("global/updateToken", globalState.token)
  store.dispatch("global/updateUserData", globalState.user)

  render(props)

  microAppStateActions = {
    onGlobalStateChange: props.onGlobalStateChange,
    offGlobalStateChange: props.offGlobalStateChange,
    setGlobalState: props.setGlobalState
  }

  props.onGlobalStateChange((value: any) => {
    (Object.keys(store.state.global) as unknown as (keyof GlobalState)[]).forEach((key) => {
      if (JSON.stringify(value[key]) === JSON.stringify(store.state.global[key])) return

      switch (key) {
        case "token":
          store.dispatch("global/updateToken", value[key])
          break
        case "user":
          store.dispatch("global/updateUserData", value[key])
          break
      }
    })
  }, true)
}

export async function unmount() {
  app?.unmount()

  if (app?._container) {
    app._container.innerHTML = ""
  }

  app = null
  router = null
  history?.destroy()
}
