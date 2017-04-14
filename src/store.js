import Vue from 'vue'
import Vuex from 'vuex'

import * as acts from './constants/types.actions.js'
import * as muts from './constants/types.mutations.js'
import * as gets from './constants/types.getters.js'

import * as gen from './api/api.gens.js'

import * as api_data from './api/api.dates.js'

Vue.use(Vuex)

const state = {
  dates: [],
}

const mutations = {
  [muts.SET_DATA](state, dates) {
    // console.log('## inside mutations: ', data);
    // const dates = {}
    // data.forEach((e) => {
        // dates[e.id] = e;
      // })
      // console.log(state.data)
    state.dates = dates
  },
}

const actions = {
  [acts.GET_DATA]({
    commit
  }) {
    console.log('## acts: ');
    api_data.resData.get().then(res => {
      console.log('## res: ', res.data.body);
      commit(muts.SET_DATA, res.data.body);
    });
  },
}

// getters are functions
const getters = {
  [gets.DATA](state) {
    return state.dates;
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
