<template>
  <div class="page-datatable">
    <h1>data table</h1>
    <div class="justify-content-center my-1 row">
    <b-form-fieldset horizontal label="Rows per page" class="col-6" :label-size="6">
      <b-form-select :options="[{text:5,value:5},{text:10,value:10},{text:15,value:15}]"
      v-model="perPage">
      </b-form-select>
    </b-form-fieldset>
    <b-form-fieldset horizontal label="Filter" class="col-6" :label-size="2">
      <b-form-input v-model="filter" placeholder="Type to Search"></b-form-input>
    </b-form-fieldset>
    </div>
    <b-table
     striped
     hover
     :filter="filter"
     :items="DATA"
     :fields="fields"
     :current-page="currentPage"
     :perPage="perPage">
       <template slot="location" scope="item">
         lng: {{item.value.coordinates[0]}}, lat: {{item.value.coordinates[1]}}
       </template>
      <template slot="actions" scope="item">
				<b-button :variant="success" size="sm" @click="open_edit(item.item)"><i class="fa fa-edit"></i></b-button>
        <b-button size="sm" @click="del(item.item)"  v-bind:ref="item.item.id"><i class="fa fa-close"></i></b-button>
      </template>
    </b-table>
		<div class="row justify-content-end" style="padding-right: 41px;">
			<b-button ref="addData" @click="open_add" style="width: 64px;"><i class="fa fa-plus"></i></b-button>
    </div>
    <div class="justify-content-center row my-1">
      <b-pagination size="md"
       :total-rows="this.DATA.length"
       :per-page="perPage"
       v-model="currentPage"/>
    </div>

		<!-- Modal for edit data {{{ -->
    <b-modal ref="modal1" title="edit data" @ok="update(cur_item)">
      <form v-on:submit.prevent="update">
        <b-form-input type="text" placeholder="enter data" v-model="cur_item.data"></b-form-input>
        <small class="text-muted">change some data</small>

				<b-form-select v-model="cur_item.type" :options="type_geos" class="col-10"></b-form-select>
				<small vertical class="text-muted">change type data</small>

        <b-button-group class="row">
				<div class="col">
        <b-form-input type="number" step="0.01" placeholder="enter latitude"
					v-model="cur_item.latitude"></b-form-input>
        <small class="text-muted">change lat position</small>
				</div>
				<div class="col">
        <b-form-input type="number" step="0.01" placeholder="enter longitude" v-model="cur_item.longitude"></b-form-input>
        <small class="text-muted">change long position</small>
				</div>
        </b-button-group>
      </form>
    </b-modal>
		<!-- }}} -->

		<!-- Modal for add data {{{ -->
    <b-modal ref="modal2" title="edit data" @ok="add(cur_item)">
      <form v-on:submit.prevent="add">
        <b-form-input type="text" placeholder="enter data" v-model="cur_item.data"></b-form-input>
        <small class="text-muted">change some data</small>

				<b-form-select v-model="cur_item.type" :options="type_geos" class="col-10"></b-form-select>
				<small vertical class="text-muted">change type data</small>

        <b-button-group class="row">
				<div class="col">
        <b-form-input type="number" step="0.01" placeholder="enter latitude"
					v-model="cur_item.latitude"></b-form-input>
        <small class="text-muted">change lat position</small>
				</div>
				<div class="col">
        <b-form-input type="number" step="0.01" placeholder="enter longitude" v-model="cur_item.longitude"></b-form-input>
        <small class="text-muted">change long position</small>
				</div>
        </b-button-group>
      </form>
    </b-modal>
		<!-- }}} -->

  </div>
</template>

<script>
  import { mapGetters, mapActions } from 'vuex'
  import * as gets from '../constants/types.getters.js'
  import * as acts from '../constants/types.actions.js'
  import * as geos from '../constants/types.some.js'

  export default {
    data() { // {{{
      return {
        docs: {
          component: 'bTable'
        },
				selected: null,
				type_geos: [],
        cur_item: {id:'dfdfdf', data:'dfdfdf', location:{type: '', coordinates:[0.0,0.0]}},
        items: [{id:'dfdfdf', data:'dfdfdf', location:{type:'point', coordinates:[1,2]}}],
        fields: {
          id: {label: 'Id', sortable: true},
          data: {label: 'Data', sortable: false},
          location: {label: 'Geo location', sortable: false},
          actions: {label: 'Actions', sortable: false},
        },
        currentPage: 1,
        perPage: 5,
        filter: null
      };
    }, // }}}
    computed: { // {{{
      ...mapGetters([
        gets.DATA,
      ]),
    }, // }}}
    mounted: function() { // {{{
			console.log(geos.GEO_OBJECT);
			this.type_geos = geos.GEO_OBJECT;
      this.GET_DATA().then(()=>{
        //this.items = this.DATA;
        //console.log('good');
      })
    }, // }}}
		methods: {
			...mapActions([ // {{{
        acts.GET_DATA,
        acts.PST_DATA,
        acts.PUT_DATA,
        acts.DEL_DATA,
			]), // }}}
			open_edit(item) { // {{{
				//this.selected = 
        this.cur_item = {
          id: item.id,
          data: item.data, type: item.location.type,
          latitude: item.location.coordinates[0],
          longitude: item.location.coordinates[1],
        };
				this.$refs.modal1.show();
			}, // }}}
			open_add(item) { // {{{
        this.cur_item = {
          id: '',
          data: '', type: '',
          latitude: 0.0,
          longitude: 0.0,
        };
				this.$refs.modal2.show();
			}, // }}}
      del(item) { // {{{
        this.DEL_DATA(item)
          .then((e)=>{
            console.log('deleted');
          })
          .catch(e => {
            console.log(e);
						this.err_blink(this.$refs[item.id]);
          })
      }, // }}}
			update(cur_item){ // {{{
				const item = {
					id: cur_item.id, data: cur_item.data,
					location:{ type: cur_item.type,
						coordinates:[ parseFloat(cur_item.latitude), parseFloat(cur_item.longitude) ]
					}
				};
				this.PUT_DATA(item)
					.then((e)=>{
						console.log('putted');
					})
					.catch(e => {
						console.log(e);
						this.err_blink(this.$refs[item.id]);
					})
			}, // }}}
			add(cur_item){ // {{{
        const item = {
					id: cur_item.id, data: cur_item.data,
					location:{
						type: cur_item.type,
						coordinates:[parseFloat(cur_item.latitude), parseFloat(cur_item.longitude)],
					}
				};
				console.log(JSON.stringify(item));
				console.log(this.$refs);
        this.PST_DATA(item)
          .then((e)=>{
            console.log('posted');
          })
          .catch(e => {
            console.log(e);
						this.err_blink(this.$refs.addData);
          })
			}, // }}}
			err_blink(e){ // {{{
        e.$el.className = e.$el.className+' show-blinking';
        setTimeout(()=>{
          e.$el.className = e.classObject.join(' ');
        }, 1000);
			} // }}}
    }
  }
  </script>

  <style>
  .page-datatable {
    width: 100%;
    height: 100%;
    position: absolute;
    float: right;
    z-index: 2;
  }
  .modal-data {
    display: block;
  }
  .show-blinking {
    background: #F00;
    transition-property: background;
    transition-duration: 0.6s;
  }
  </style>
