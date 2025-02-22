<template>
  <div id="app" v-bkloading="{ isLoading: globalLoading }" :bk-language="$i18n.locale"
    :class="{
      'no-breadcrumb': hideBreadcrumbs,
      'main-full-screen': mainFullScreen
    }">
    <div class="browser-tips" v-if="showBrowserTips">
      <span class="tips-text">{{$t('您的浏览器非Chrome，建议您使用最新版本的Chrome浏览，以保证最好的体验效果')}}</span>
      <i class="tips-icon bk-icon icon-close-circle-shape" @click="showBrowserTips = false"></i>
    </div>
    <the-header v-if="currentRoute !== '/login'"></the-header>
    <router-view class="views-layout" :name="topView" ref="topView"></router-view>
    <the-permission-modal ref="permissionModal"></the-permission-modal>
    <the-login-modal ref="loginModal"
      v-if="loginUrl"
      :login-url="loginUrl"
      :success-url="loginSuccessUrl">
    </the-login-modal>
  </div>
</template>

<script>
  import theHeader from '@/components/layout/header'
  import thePermissionModal from '@/components/modal/permission'
  import theLoginModal from '@blueking/paas-login'
  import { addResizeListener, removeResizeListener } from '@/utils/resize-events'
  import { mapGetters } from 'vuex'
  export default {
    name: 'app',
    components: {
      theHeader,
      thePermissionModal,
      theLoginModal
    },
    data() {
      const showBrowserTips = window.navigator.userAgent.toLowerCase().indexOf('chrome') === -1
      return {
        showBrowserTips,
        loginSuccessUrl: `${window.location.origin}/static/login_success.html`
      }
    },
    computed: {
      ...mapGetters(['globalLoading', 'mainFullScreen']),
      ...mapGetters('userCustom', ['usercustom', 'firstEntryKey', 'classifyNavigationKey']),
      hideBreadcrumbs() {
        return !(this.$route.meta.layout || {}).breadcrumbs
      },
      topView() {
        const [topRoute] = this.$route.matched
        return (topRoute && topRoute.meta.view) || 'default'
      },
      currentRoute() {
        return this.$route.fullPath
      },
      loginUrl() {
        if (process.env.NODE_ENV === 'development') {
          return ''
        }
        const siteLoginUrl = this.$Site.login || ''
        const [loginBaseUrl] = siteLoginUrl.split('?')
        if (loginBaseUrl) {
          return `${loginBaseUrl}plain`
        }
        return ''
      }
    },
    mounted() {
      addResizeListener(this.$el, this.calculateAppHeight)
      window.permissionModal = this.$refs.permissionModal
      window.loginModal = this.$refs.loginModal
    },
    beforeDestroy() {
      removeResizeListener(this.$el, this.calculateAppHeight)
    },
    methods: {
      calculateAppHeight() {
        this.$store.commit('setAppHeight', this.$el.offsetHeight)
      }
    }
  }
</script>
<style lang="scss" scoped>
    #app{
        height: 100%;
    }
    .browser-tips{
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 40px;
        line-height: 40px;
        text-align: center;
        color: #ff5656;
        background-color: #f8f6db;
        z-index: 99999;
        .tips-text{
            margin: 0 20px 0 0 ;
        }
        .tips-icon{
            cursor: pointer;
        }
    }
    .views-layout{
        height: calc(100% - 58px);
    }
    // 主内容区全屏
    .main-full-screen {
        /deep/ {
            .header-layout,
            .nav-layout,
            .breadcrumbs-layout {
                display: none;
            }
        }
        .views-layout {
            height: 100%;
        }
    }
    .no-breadcrumb {
        /deep/ {
            .main-layout {
                margin-top: 0
            }
            .main-views {
                height: 100%;
                margin-top: 0;
            }
        }
    }
</style>
