<template>
  <div id="app">
    <div class="main-panel">
      <b-navbar toggleable type="inverse" variant="inverse" class="dvi-navbar collapse">
        <b-collapse isNav id="nav_collapse">
          <b-nav isNavBar>
            <router-link to="/">
              <b-nav-item>data easy</b-nav-item>
            </router-link>
            <router-link to="/datatable">
              <b-nav-item>data table</b-nav-item>
            </router-link>
            <router-link to="/datamap">
              <b-nav-item>data map</b-nav-item>
            </router-link>
          </b-nav>
        </b-collapse>
      </b-navbar>
    </div>
    <router-view class="view"></router-view>
  </div>
</template>

<script>
  export default {
  }
  </script>

  <style>
  html,
  body {
  }
  .main-panel {
    width: 100%;
    /*position: absolute;*/
    z-index: 2;
  }
  </style>
