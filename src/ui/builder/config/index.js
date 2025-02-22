'use strict'
// Template version: 1.3.1
// see http://vuejs-templates.github.io/webpack for documentation.

const path = require('path')
const fs = require('fs')
const parseArgs = require('minimist')

const config = {
  BUILD_TITLE: '',
  BUILD_OUTPUT: '../../Dockerfile/ui'
}

const argv = parseArgs(process.argv.slice(2))

process.argv.slice(2).forEach((str) => {
  const arg = str.split('=')
  if (Object.prototype.hasOwnProperty.call(config, arg[0])) {
    config[arg[0]] = arg.slice(1).join('=')
  }
})
process.CMDB_CONFIG = config
const dev = {
  // custom config
  config: Object.assign({}, config, {
    API_URL: JSON.stringify('http://localhost:9090/'),
    API_VERSION: JSON.stringify('v3'),
    API_LOGIN: JSON.stringify('login'),
    AGENT_URL: JSON.stringify(''),
    AUTH_SCHEME: JSON.stringify('internal'),
    AUTH_CENTER: JSON.stringify({}),
    BUILD_VERSION: JSON.stringify('dev'),
    USER_ROLE: JSON.stringify(1),
    USER_NAME: JSON.stringify('admin'),
    FULL_TEXT_SEARCH: JSON.stringify('off'),
    USER_MANAGE: JSON.stringify(''),
    HELP_DOC_URL: JSON.stringify(''),
    DISABLE_OPERATION_STATISTIC: false
  }),

  // Paths
  assetsSubDirectory: '',
  assetsPublicPath: '',
  proxyTable: {
    '/meta-ui/*': {
      logLevel: 'info',
      changeOrigin: true,
      target: 'http://172.22.50.25:32168/',
      pathRewrite: {
        '^/meta-ui/': ''
      }
    },
    '/api/*': {
      logLevel: 'info',
      changeOrigin: true,
      target: 'http://172.22.50.25:32168/'
    },
    '/ldap/*': {
      logLevel: 'info',
      changeOrigin: true,
      target: 'http://172.22.50.191:8090/'
    }
  },
  // Various Dev Server settings
  host: 'localhost', // can be overwritten by process.env.HOST
  port: 9090, // can be overwritten by process.env.PORT, if port is in use, a free one will be determined
  autoOpenBrowser: true,
  errorOverlay: true,
  notifyOnErrors: true,
  poll: false, // https://webpack.js.org/configuration/dev-server/#devserver-watchoptions-

  // Use Eslint Loader?
  // If true, your code will be linted during bundling and
  // linting errors and warnings will be shown in the console.
  useEslint: true,
  // If true, eslint errors and warnings will also be shown in the error overlay
  // in the browser.
  showEslintErrorsInOverlay: true,

  /**
     * Source Maps
     */

  // https://webpack.js.org/configuration/devtool/#development
  devtool: 'cheap-module-eval-source-map',

  // If you have problems debugging vue-files in devtools,
  // set this to false - it *may* help
  // https://vue-loader.vuejs.org/en/options.html#cachebusting
  cacheBusting: true,

  cssSourceMap: true
}

const customDevConfigPath = path.resolve(__dirname, `index.dev.${argv.env || 'ee'}.js`)
const isCustomDevConfigExist = fs.existsSync(customDevConfigPath)
if (isCustomDevConfigExist) {
  const customDevConfig = require(customDevConfigPath)
  Object.assign(dev, customDevConfig)
}

module.exports = {
  dev,

  build: {
    // custom config
    config: Object.assign({}, config, {
      API_URL: '{{.site}}',
      API_VERSION: '{{.version}}',
      BUILD_VERSION: '{{.ccversion}}',
      API_LOGIN: '{{.curl}}',
      AGENT_URL: '{{.agentAppUrl}}',
      AUTH_SCHEME: '{{.authscheme}}',
      AUTH_CENTER: '{{.authCenter}}',
      USER_ROLE: '{{.role}}',
      USER_NAME: '{{.userName}}',
      FULL_TEXT_SEARCH: '{{.fullTextSearch}}',
      USER_MANAGE: '{{.userManage}}',
      HELP_DOC_URL: '{{.helpDocUrl}}',
      DISABLE_OPERATION_STATISTIC: '{{.disableOperationStatistic}}'
    }),

    // Template for index.html
    index: `${path.resolve(config.BUILD_OUTPUT)}/web/index.html`,

    // Template for login.html
    login: `${path.resolve(config.BUILD_OUTPUT)}/web/login.html`,

    // Paths
    assetsRoot: `${path.resolve(config.BUILD_OUTPUT)}/web`,

    assetsSubDirectory: '',
    assetsPublicPath: '/meta/',

    /**
         * Source Maps
         */

    productionSourceMap: true,
    // https://webpack.js.org/configuration/devtool/#production
    devtool: '#source-map',

    // Gzip off by default as many popular static hosts such as
    // Surge or Netlify already gzip all static assets for you.
    // Before setting to `true`, make sure to:
    // npm install --save-dev compression-webpack-plugin
    productionGzip: false,
    productionGzipExtensions: ['js', 'css'],

    // Run the build command with an extra argument to
    // View the bundle analyzer report after build finishes:
    // `npm run build --report`
    // Set to `true` or `false` to always turn it on or off
    bundleAnalyzerReport: process.env.npm_config_report
  }
}
