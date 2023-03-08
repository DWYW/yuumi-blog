import { ActionTree } from "vuex"
import { ApplicationState, RootState } from "@/store/types/index"
import { MenuAPI } from "@/api"
import router from "@/router"
import { $helper } from "@/common"

const MUTATION_TYPE = {
  UPDATE_LOADED: "UPDATE_LOADED",
  UPDATE_NAVS: "UPDATE_NAVS"
}

const state: ApplicationState = {
  navsLoaded: false,
  navs: []
}

const actions: ActionTree<ApplicationState, RootState> = {
  updateNavs: function({ commit }) {
    return MenuAPI.getMineNavMenus().then(({ menus }) => {
      commit(MUTATION_TYPE.UPDATE_LOADED, true)
      commit(MUTATION_TYPE.UPDATE_NAVS, $helper.list2tree(menus))
    })
  },
  signout: function({ dispatch }) {
    dispatch("global/clearToken", undefined, { root: true })
    dispatch("global/clearUserData", undefined, { root: true })
    router.replace("/signin")
  }
}

const mutations = {
  [MUTATION_TYPE.UPDATE_LOADED]: function(state: ApplicationState, value: boolean) {
    state.navsLoaded = value
  },
  [MUTATION_TYPE.UPDATE_NAVS]: function(state: ApplicationState, navs: any[]) {
    state.navs = navs
  }
}

export default {
  namespaced: true,
  state,
  actions,
  mutations
}
