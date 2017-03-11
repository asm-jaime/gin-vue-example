import Vue from 'vue'
import Vuex from 'vuex'

import * as acts from './constants/types.actions.js'
import * as muts from './constants/types.mutations.js'
import * as gets from './constants/types.getters.js'

import * as gen from './api/api.gens.js'

import * as api_data from './api/api.dates.js'

Vue.use(Vuex)

const state = {
  data: [],
}

const mutations = {
  [muts.SET_DATA](state, data) {
    console.log('## inside mutations: ', data);
    state.data.push(data);
  },
}

const actions = {
  [acts.GET_DATA]({
    commit
  }) {
    api_data.getData.get().then(res => {
      console.log('## res: ', res.data.body);
      commit(muts.SET_DATA, res.data.body);
    });
  },
}

// getters are functions
const getters = {
  [gets.GET_DATA](state) {
    return state.data;
  },
}

// A Vuex instance is created by combining the state, mutations, actions,
// and getters.
export default new Vuex.Store({
  state,
  getters,
  actions,
  mutations
})
