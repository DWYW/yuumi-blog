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
        { name: "keywords", content: "blog,dwyw,scolint" },
        { name: "description", content: "scolint's blog" },
        { name: "viewport", content: "width=device-width, initial-scale=1.0" },
        { name: "author", content: "scolint" },
        // 360，优先使用解析内核
        { name: "renderer", content: "webkit|ie-comp|ie-stand" },
        // Google Chrome
        // 优先使用最新的chrome版本
        { "http-equiv": "X-UA-Compatible", content: "chrome=1" },
        // 禁止自动翻译
        { name: "google", content: "notranslate" }
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
