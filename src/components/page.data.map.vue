<template>
  <div class="page-map">
    <div>dataMap</div>
  </div>
</template>

<script>

  import { mapGetters, mapActions } from 'vuex'
import * as gets from '../constants/types.getters.js'
import * as acts from '../constants/types.actions.js'

export default {
  name: 'PageDataMap',

  data(){
    return {
    }
  },
  computed: {
    ...mapGetters([
    ]),
  },
  methods: {
    ...mapActions([
    ]),
    showConstants: function(){
      console.log("## some msg");
    }
  }
}

</script>

<style>

	.page-map {
    width: 300px;
    height: 300px;
		top: 20%;
    left: 1%;
    position: absolute;
		float: right;
    z-index: 2;
		background: azure;
	}
</style>
