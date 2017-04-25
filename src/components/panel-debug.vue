<template>
  <div>
  <div class="debug-panel">test
    <div class="container">{{ DATA }}</div>
    <b-button :size="size" :variant="variant"
      class="btn btn-primary" @click="GET_DATA">get data
    </b-button>
    <b-button :size="size" :variant="variant"
      class="btn btn-primary" @click="showConstants()">show constants
    </b-button>
  </div>
  <router-view></router-view>
  </div>
</template>

<script>

import { mapGetters, mapActions } from 'vuex'
import * as gets from '../constants/types.getters.js'
import * as acts from '../constants/types.actions.js'

export default {
  name: 'PanelDebug',

  data(){
    return {
    }
  },
  computed: {
    ...mapGetters([
      gets.DATA,
    ]),
  },
  methods: {
    ...mapActions([
      acts.GET_DATA,
    ]),
    showConstants: function(){
      console.log("## some msg");
    }
  },
}
</script>

<style>
	.debug-panel {
    width: 300px;
    height: 300px;
		top: 20%;
    right: 1%;
    position: absolute;
		float: right;
    z-index: 2;
		background: azure;
	}
</style>
