import Vue from 'vue'
import VueResource from 'vue-resource'

import {DATES as api_data} from '../constants/paths.api.js'

Vue.use(VueResource)
// Vue.http.options.emulateJSON = true
// Vue.http.options.emulateHTTP = true
Vue.http.options.crossOrigin = true
// Vue.http.options.credentials = true

export const resData = Vue.resource(api_data.DATA)
