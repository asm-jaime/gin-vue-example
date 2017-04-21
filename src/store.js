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
    state.dates = dates;
  },
  [muts.DEL_DATA](state, data) {
    const index = state.dates.findIndex((e) => e.id === data.id);
    state.dates.splice(index, 1);
  },
  [muts.PUT_DATA](state, data) {
    console.log(data);
    const index = state.dates.findIndex((e) => e.id === data.id);
    if (index > -1) {
      state.dates.splice(index, 1, data);
    } else {
      state.dates.unshift(data);
    }
  },
}

const actions = {
  [acts.GET_DATA]({
    commit
  }) { // {{{
    return api_data.resData.get().then(res => {
      commit(muts.SET_DATA, res.data.body);
    });
  }, // }}}
  [acts.PST_DATA]({
    commit
  }, data) { // {{{
    return api_data.resData.save({}, data).then(res => {
      if (res.status === 200) {
        commit(muts.PUT_DATA, res.data.body);
      } else {
        throw new Error('can\'t set this data');
      }
    });
  }, // }}}
  [acts.PUT_DATA]({
    commit
  }, data) { // {{{
    return api_data.resData.update({}, data).then(res => {
      if (res.status === 200) {
        commit(muts.PUT_DATA, res.data.body);
      } else {
        throw new Error('can\'t put this data');
      }
    });
  }, // }}}
  [acts.DEL_DATA]({
    commit
  }, data) { // {{{
    console.log('data: ', data)
    return api_data.resData.remove({}, {
      id: data.id
    }).then(res => {
      if (res.status === 200) {
        commit(muts.DEL_DATA, data);
      } else {
        throw new Error('can\'t deleted this data');
      }
    })
  }, // }}}
}

const getters = {
  [gets.DATA](state) {
    return state.dates;
  },
}

export default new Vuex.Store({
  state,
  getters,
  actions,
  mutations
})
