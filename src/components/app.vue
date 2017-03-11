<template>
    <div id="app">
        <div class="main-panel">
            <b-navbar toggleable type="inverse" variant="inverse" class="dvi-navbar collapse">
                <b-collapse isNav id="nav_collapse">
                    <b-nav isNavBar>
                        <router-link to="/event">
                            <b-nav-item>custom event</b-nav-item>
                        </router-link>
                        <router-link to="/">
                            <b-nav-item>dvij</b-nav-item>
                        </router-link>
                        <router-link to="/users">
                            <b-nav-item>users</b-nav-item>
                        </router-link>
                        <router-link to="/events">
                            <b-nav-item>events</b-nav-item>
                        </router-link>
                        <router-link :to="{ name: 'bar', params: { id: 123 }}">bar</router-link>
                    </b-nav>
                </b-collapse>
            </b-navbar>
        </div>
        <div class="container-fluid">
            <router-view class="view"></router-view>
        </div>
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
