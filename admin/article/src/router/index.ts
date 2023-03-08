import { RouteRecordRaw } from "vue-router"

export const routes: Array<RouteRecordRaw> = [{
  path: "/",
  name: "home",
  component: () => import(/* webpackChunkName: "home" */ "../views/home.vue")
}, {
  path: "/:childpath*",
  name: "404",
  component: () => import("../views/404.vue")
}]
