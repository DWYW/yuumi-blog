import { GlobalState } from "@/store/types"
import { initGlobalState, MicroAppStateActions } from "qiankun"
import { store } from "@/store"

const microAppStateActions: MicroAppStateActions = initGlobalState({})

export function setGlobalState(state: Record<string, any>) {
  microAppStateActions.setGlobalState(state)
}

export function listenGlobalStateChange() {
  microAppStateActions.onGlobalStateChange((value) => {
    (Object.keys(store.state.global) as unknown as (keyof GlobalState)[]).forEach((key) => {
      if (JSON.stringify(value[key]) === JSON.stringify(store.state.global[key])) return

      switch (key) {
        case "token":
          store.dispatch("global/updateToken", value[key])
          break
        case "user":
          store.dispatch("global/updateUser", value[key])
          break
      }
    })
  }, true)
}
