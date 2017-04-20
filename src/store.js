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
  [muts.DEL_DATA](state, data) {
    const index = state.dates.findIndex(e => e.id === data.id);
    state.dates.splice(index, 1);
    console.log(state.dates);
  },
}

const actions = {
  [acts.GET_DATA]({
    commit
  }) {
    return api_data.resData.get().then(res => {
      console.log('## get data: ', res.data);
      commit(muts.SET_DATA, res.data.body);
    });
  },
  [acts.PST_DATA]({
    commit
  }, data) {
    api_data.resData.save(data).then(res => {
      console.log('## post data: ', res.data);
      // commit(muts.SET_DATA, res.data.body);
    });
  },
  [acts.PUT_DATA]({
    commit
  }, data) {
    api_data.resData.update(data).then(res => {
      console.log('## update data: ', res.data);
      // commit(muts.SET_DATA, res.data.body);
    });
  },
  [acts.DEL_DATA]({
    commit
  }, data) {
    console.log('data: ', data)
    return api_data.resData.remove({}, {id: data.id}).then(res => {
      if(res.status === 200){
        commit(muts.DEL_DATA, data);
      } else {
        throw new Error('can\'t deleted this data');
      }
    })
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
