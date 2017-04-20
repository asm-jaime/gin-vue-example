import Vue from 'vue'
import VueRouter from 'vue-router'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

import App from './components/app.vue'

import store from './store.js'
import PanelDebug from './components/panel.debug.vue'

import PanelData from './components/panel.data.vue'
import PageDataTable from './components/page-data-table.vue'
import PageDataMap from './components/page.data.map.vue'

// require('../node_modules/bootstrap-vue/dist/bootstrap-vue.css')

Vue.use(VueRouter)
Vue.use(BootstrapVue)

const router = new VueRouter({
  mode: 'history',
  base: __dirname,
  routes: [
    {path: '/', component: PanelDebug, children: [
      {path: 'datamap', component: PageDataMap, children: [
        {path: ':id', component: PanelData }
      ]},
    ]},
    {path: '/datatable', component: PageDataTable},
  ]
})

new Vue({
  el: '#app',
  store,
  router,
  render: h => h(App),
})
