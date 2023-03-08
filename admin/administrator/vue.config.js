const { defineConfig } = require('@vue/cli-service')
const { name, version } = require('./package.json')
const path = require('path')
const microName = path.parse(name).name

let publicPath = process.env.VUE_APP_PUBLIC_PATH
if (process.env.NODE_ENV === 'production') {
  publicPath += `${version}/`
}

module.exports = defineConfig({
  outputDir: `dist/${version}`,
  publicPath: publicPath,
  devServer: {
    port: 8083,
    headers: {
      'Access-Control-Allow-Origin': '*',
    }
  },
  css: {
    loaderOptions: {
      sass: {
        additionalData: '@import "yuumi-ui-vue/packages/theme.scss";'
      }
    }
  },
  configureWebpack: {
    output: {
      library: `${microName}`,
      libraryTarget: 'umd' // 把微应用打包成 umd 库格式
    }
  },
  /**
   * transpileDependencies
   * 默认情况下 babel-loader 会忽略所有 node_modules 中的文件。你可以启用本选项，以避免构建后的代码中出现未转译的第三方依赖。
   * 不过，对所有的依赖都进行转译可能会降低构建速度。如果对构建性能有所顾虑，你可以只转译部分特定的依赖：
   * 给本选项传一个数组，列出需要转译的第三方包包名或正则表达式即可。
   */
  transpileDependencies: process.env.NODE_ENV !== 'production' ? false : [
    'yuumi-request',
    'yuumi-ui-vue',
  ],
  chainWebpack: function (config) {
    mergeHtmlConfig(config)
    mergeDefineConfig(config)
    return config
  }
})

function mergeHtmlConfig (config) {
  config.plugin('html').tap(args => {
    args[0].title = process.env.VUE_APP_NAME
    return args
  })
}

function mergeDefineConfig (config) {
  config.plugin('define').tap(args => {
    args[0]['process.env'].__APP_VERSION__ = JSON.stringify(version)
    args[0]['process.env'].__APP_MICRO_NAME__ = JSON.stringify(microName)
    return args
  })
}