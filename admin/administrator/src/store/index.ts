import { InjectionKey } from "vue"
import { createStore, Store, useStore as baseUseStore } from "vuex"
import { RootState } from "./types"
import global from "./modules/global"

export const key: InjectionKey<Store<RootState>> = Symbol("")

export const store = createStore<RootState>({
  modules: {
    global
  },
  strict: process.env.NODE_ENV !== "production"
})

export function useStore() {
  return baseUseStore(key)
}
