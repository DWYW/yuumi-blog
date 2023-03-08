import { store } from "@/store"
import { registerMicroApps, start } from "qiankun"
import { YuumiRequest } from "yuumi-request"
import { MicroApplicationConfig } from "./config"
import { listenGlobalStateChange } from "./global-state"

export function registerApplications(apps: MicroApplicationConfig[]) {
  const yuumi = new YuumiRequest()

  Promise.all(apps.map(item => {
    // 解析manifest
    return new Promise<MicroApplicationConfig>((resolve, reject) => {
      if (/manifest\.json$/.test(item.entry)) {
        yuumi.get(item.entry, { _t: Date.now() }).then(({ response }) => {
          if (typeof response === "string") {
            response = JSON.parse(response)
          }
          item.entry = response.index

          resolve(item)
        }).catch(err => {
          reject(err)
        })
      } else {
        resolve(item)
      }
    })
  })).then((res) => res.map((item) => Object.assign({
    container: `#${item.name}`,
    props: {
      getGlobalState: () => store.state.global,
      getSignoutFun: () => () => store.dispatch("application/signout")
    }
  }, item))).then((res) => {
    // 注册应用
    registerMicroApps(res)
    // 启动
    start()
    // 监听 global state change
    listenGlobalStateChange()
  }).catch((err) => {
    console.error(err)
  })
}
