const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  // 关闭语法监测
  lintOnSave:false,
  outputDir: 'weapp.shujingyun.com',
  assetsDir: 'static',
  // publicPath: '/admin/'
  publicPath: process.env.VUE_APP_PUBLIC_PATH
  // publicPath: "./",
  // publicPath: "/weapp.shujingyun.com/",
})
