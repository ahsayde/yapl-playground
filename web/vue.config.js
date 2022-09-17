const { defineConfig } = require('@vue/cli-service')
const monacoWebpackPlugin = require("monaco-editor-webpack-plugin");

module.exports = defineConfig({
  publicPath: process.env.NODE_ENV === 'production' ? '/yapl-playground' : '/',
  transpileDependencies: [
    'vuetify'
  ],
  configureWebpack: {
    plugins: [new monacoWebpackPlugin()]
  }
})

