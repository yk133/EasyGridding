<template>

  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <Search msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>

import Search from './components/Search.vue'
import Vue from 'vue';
import ElementUI from 'element-ui';

import 'element-ui/lib/theme-chalk/index.css';

Vue.use(ElementUI);

export default {
  name: 'App',
  components: {
    Search
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
