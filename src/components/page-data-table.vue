<template>
  <div class="page-datatable">
    <h1>data table</h1>
    <div class="justify-content-centermy-1 row">
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
        <b-btn size="sm" @click="edit(item.item)"><i class="fa fa-edit"></i></b-btn>
        <b-btn size="sm" @click="del(item.item)"  v-bind:ref="item.item.id"><i class="fa fa-close"></i></b-btn>
      </template>
    </b-table>
    <div class="justify-content-center row my-1">
      <b-pagination size="md"
       :total-rows="this.DATA.length"
       :per-page="perPage"
       v-model="currentPage"
       />
    </div>
  </div>
</template>

<script>
  import { mapGetters, mapActions } from 'vuex'
  import * as gets from '../constants/types.getters.js'
  import * as acts from '../constants/types.actions.js'

  export default {
    data() { // {{{
      return {
        docs: {
          component: 'bTable'
        },
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
      this.GET_DATA().then(()=>{
        //this.items = this.DATA;
        //console.log('good');
      })
    }, // }}}
    methods: {
      ...mapActions([
        acts.GET_DATA,
        acts.PST_DATA,
        acts.PUT_DATA,
        acts.DEL_DATA,
      ]),
      del(item) {
        console.log('delete');
        try {
          this.DEL_DATA(item)
        } catch(err) {
          console.log('some error: ', err)
          this.$refs[item.id].$el.className = this.$refs[item.id].$el.className+' show-blinking';
          setTimeout(()=>{
            this.$refs[item.id].$el.className = this.$refs[item.id].classObject.join(' ');
          }, 1000)
        }
      },
      edit(item) {
       // try {
       //   this.DEL_DATA(item)
       // } catch(err) {
       //   console.log('bad request')
       // }
      },
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
    background: azure;
  }
  .show-blinking {
    background: #F00;
    transition-property: background;
    transition-duration: 0.6s;
  }
  </style>
