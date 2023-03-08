import { $storage } from "@/common"
import { ACCESS_TOKEN } from "@/constant/keys"
import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router"

const routes: RouteRecordRaw[] = [{
  path: "/signin",
  name: "signin",
  component: () => import("../views/signin.vue")
}, {
  path: "/",
  name: "home",
  component: () => import("../views/home.vue"),
  children: [{
    path: "micro/:childpath*",
    name: "microApp",
    component: () => import("../views/micro.vue")
  }]
}, {
  path: "/:childpath*",
  name: "404",
  component: () => import("../views/404.vue")
}]

const router = createRouter({
  routes,
  history: createWebHashHistory()
})

router.beforeEach((to, from, next) => {
  const token = $storage.local.getItem(ACCESS_TOKEN)

  if (!token && !/^\/signin/.test(to.path)) {
    next("/signin")
  } else if (token && /^\/signin/.test(to.path)) {
    next(from.fullPath)
  } else {
    next()
  }
})

export default router
