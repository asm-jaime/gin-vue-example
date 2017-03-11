import Vue from 'vue'
import VueResource from 'vue-resource'

import {API_DATA as datas} from '../constants/paths.api.js'

Vue.use(VueResource)

Vue.http.options.crossOrigin = true
// Vue.http.options.credentials = true

export const getData = Vue.resource(points.GET_DATA)
