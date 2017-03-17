<template>
  <div class="page-datatable">
    <h1>dfdff</h1>
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
     :items="items"
     :fields="fields"
     :current-page="currentPage"
     :perPage="perPage">
       <template slot="name" scope="item">
         {{item.value.first}} {{item.value.last}}
       </template>
      <template slot="actions" scope="item">
        <b-btn size="sm" @click="details(item.item)">Details</b-btn>
      </template>
    </b-table>
    <div class="justify-content-center row my-1">
      <b-pagination size="md"
       :total-rows="this.items.length"
       :per-page="perPage"
       v-model="currentPage"
       />
    </div>
  </div>
</template>

<script>

  export default {
    data() {
      return {
        docs: {
          component: 'bTable'
        },
        items: [
          {
            isActive: true,
            age: 40,
            name: {
              first: 'Dickerson',
              last: 'Macdonald'
            }
          },
          {
            isActive: false,
            age: 21,
            name: {
              first: 'Larsen',
              last: 'Shaw'
            }
          },
          {
            isActive: false,
            age: 26,
            state: 'success',
            name: {
              first: 'Mitzi',
              last: 'Navarro'
            }
          },
          {
            isActive: false,
            age: 22,
            name: {
              first: 'Geneva',
              last: 'Wilson'
            }
          },
          {
            isActive: true,
            age: 38,
            name: {
              first: 'Jami',
              last: 'Carney'
            }
          },
          {
            isActive: false,
            age: 27,
            name: {
              first: 'Essie',
              last: 'Dunlap'
            }
          },
          {
            isActive: true,
            age: 40,
            name: {
              first: 'Dickerson',
              last: 'Macdonald'
            }
          },
          {
            isActive: false,
            age: 21,
            name: {
              first: 'Larsen',
              last: 'Shaw'
            }
          },
          {
            isActive: false,
            age: 26,
            name: {
              first: 'Mitzi',
              last: 'Navarro'
            }
          },
          {
            isActive: false,
            age: 22,
            name: {
              first: 'Geneva',
              last: 'Wilson'
            }
          },
          {
            isActive: true,
            age: 38,
            name: {
              first: 'Jami',
              last: 'Carney'
            }
          },
          {
            isActive: false,
            age: 27,
            name: {
              first: 'Essie',
              last: 'Dunlap'
            }
          }
        ],
        fields: {
          name: {label: 'Person Full name', sortable: true},
          age: {label: 'Person age', sortable: true},
          isActive: {label: 'is Active'},
          actions: {label: 'Actions'}
        },
        currentPage: 1,
        perPage: 5,
        filter: null
      };
    },
    methods: {
      details(item) {
        // eslint-disable-next-line no-alert
        alert(JSON.stringify(item));
      }
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
  </style>
