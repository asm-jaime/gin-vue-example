import Vue from 'vue'
import VueRouter from 'vue-router'
import BootstrapVue from 'bootstrap-vue'

import App from './components/app.vue'

import store from './store.js'
import PanelDebug from './components/panel.debug.vue'

Vue.use(VueRouter)
Vue.use(BootstrapVue)

const Users = { template: '<div>users</div>' }
const Bar = { template: '<div>This is Bar {{ $route.params.id }}</div>' }

const router = new VueRouter({
  mode: 'history',
  base: __dirname,
  routes: [
    {
      path: '/', component: PanelDebug,
      children: [
        {path: '/bar/:id', name: 'bar', component: Bar}
      ],
    },
  ]
})

new Vue({
  el: '#app',
  store,
  router,
  render: h => h(App),
})
