<template>
  <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="Welcome to Your Vue.js App" Hidden/>
  <div>
    <button @click="inc" id="counter">Кнопка {{counter}}</button>
    <button @click="reqHome">API /</button>
    <button @click="reqCompany">API /Company</button>
    <div>{{info}}</div>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

import axios from 'axios';

export default {
  name: 'App',
  components: {
    HelloWorld
  },
  data(){
    return{
      counter: 0,
      info: ""
    }
  },
  methods: {
    inc(){
      return this.counter++
    },
    reqHome(){
      // axios
      // .get('https://api.coindesk.com/v1/bpi/currentprice.json')
      // .then(response => (this.info = response));  
      axios
      .get('http://localhost:9000/')
      .then(response => (this.info = response.data))
    },
    reqCompany(){
      axios
      .get("http://localhost:9000/Company")
      .then(response => (this.info = response.data))
    }
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
