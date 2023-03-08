import { $storage } from "@/common"
import { ACCESS_TOKEN, SIGNIN_USER } from "@/constant/keys"
import { GlobalState, RootState, User } from "@/store/types/index"
import { ActionTree } from "vuex"
import { setGlobalState } from "@/constant/application/global-state"

const MUTATION_TYPE = {
  UPDATE_TOKEN: "UPDATE_TOKEN",
  UPDATE_USER_DATA: "UPDATE_USER_DATA"
}

const state: GlobalState = {
  token: "",
  user: null
}

const actions: ActionTree<GlobalState, RootState> = {
  updateToken: function({ commit }, token: string): void {
    if (!token) return
    commit(MUTATION_TYPE.UPDATE_TOKEN, token)
  },
  updateUserData: function({ commit }, data: User) {
    if (!data) return

    function setAttributeValue<O, T extends keyof O>(key: T, value: O[T], obj: O) {
      obj[key] = value
    }

    const user: User = {}
    const attributes = ["id", "name"] as (keyof User)[]
    attributes.forEach((key) => {
      setAttributeValue(key, data[key], user)
    })

    commit(MUTATION_TYPE.UPDATE_USER_DATA, user)
  },
  clearToken: function({ commit }) {
    commit(MUTATION_TYPE.UPDATE_TOKEN, "")
  },
  clearUserData: function({ commit }) {
    commit(MUTATION_TYPE.UPDATE_USER_DATA, null)
  },
  syncFromLocal: function({ dispatch }) {
    dispatch("updateToken", $storage.local.getItem(ACCESS_TOKEN))
    dispatch("updateUserData", $storage.local.getItem(SIGNIN_USER))
  }
}

const mutations = {
  [MUTATION_TYPE.UPDATE_TOKEN]: function(state: GlobalState, token: string) {
    state.token = token
    if (token !== $storage.local.getItem(ACCESS_TOKEN)) {
      token ? $storage.local.setItem(ACCESS_TOKEN, token) : $storage.local.removeItem(ACCESS_TOKEN)
    }
    setGlobalState(state)
  },
  [MUTATION_TYPE.UPDATE_USER_DATA]: function(state: GlobalState, value: User|null) {
    state.user = value
    if (JSON.stringify(value) !== $storage.local.getItem(SIGNIN_USER)) {
      value ? $storage.local.setItem(SIGNIN_USER, value) : $storage.local.removeItem(SIGNIN_USER)
    }

    setGlobalState(state)
  }
}

export default {
  namespaced: true,
  state,
  actions,
  mutations
}
