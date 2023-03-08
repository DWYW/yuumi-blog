// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  typescript: {
    shim: false
  },
  app: {
    head: {
      htmlAttrs: {
        lang: "zh"
      },
      charset: "utf-16",
      title: "SCOLINT‘S BLOG",
      meta: [
        { name: "keyword", content: "blog,dwyw,scolit" },
        { name: "description", content: "scolint's blog" }
      ]
    }
  },
  devServer: {
    port: 8080
  },
  runtimeConfig: {
    public: {
      ARTICLE_API_PRIFIX: "http://127.0.0.1:9007",
      PASSPORT_API_PRIFIX: "http://127.0.0.1:9005",
      UA: "google analytics 衡量 ID",
      // https://github.com/settings/developers
      GITHUB_CLIENTID: "client_id",
      GITHUB_HOME: "https://github.com/DWYW"
    }
  },
  modules: [
    ['@pinia/nuxt', {
      autoImports: [
        // 自动引入 `usePinia()`
        'defineStore',
        // 自动引入 `usePinia()` 并重命名为 `usePiniaStore()`
        ['defineStore', 'definePiniaStore'],
      ],
    }]
  ],
  plugins: [
    "~/plugins/ga.client.ts",
    "~/plugins/store.client.ts",
  ]
})
