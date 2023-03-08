export const DYNAMIC_ROUTES = [{
  path: "system/article/list",
  name: "ArticleList",
  component: () => import("../pages/article/list/index.vue"),
  meta: {}
}, {
  path: "system/article/detail/:id",
  name: "ArticleDetail",
  component: () => import("../pages/article/detail/index.vue"),
  meta: {}
}]
