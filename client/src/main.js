// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import VueAuth from '@websanova/vue-auth'
import NProgress from 'vue-nprogress'
import { sync } from 'vuex-router-sync'
import App from './App'
import router from './router'

Vue.config.productionTip = false

Vue.router = router
Vue.use(VueAxios, axios)
Vue.use(VueAuth, {
  auth: {
    request: function (req, token) {
      this.options.http._setheaders.call(this, req, {Authorization: 'Bearer ' + token})
    },
    response: function (res) {
      return res.data
    }
  },
  http: require('@websanova/vue-auth/drivers/http/axios.1.x.js'),
  router: require('@websanova/vue-auth/drivers/router/vue-router.2.x.js'),
  loginData: { url: 'http://localhost:8099/login', fetchUser: false },
  refreshData: { enabled: false }
})

Vue.use(NProgress)

Vue.config.devtools = true
sync(store, router)

const nprogress = new NProgress({ parent: '.nprogress-container' })
const { state } = store

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})
