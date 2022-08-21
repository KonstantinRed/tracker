<template>
  <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
  <!-- <HelloWorld msg="Welcome to Your Vue.js App" Hidden/> -->
  
  <TableTask/>
  <!-- <TestAPI/> -->
</template>

<script>
// import TableCompany from './components/TableCompany.vue'
// import TestAPI from './components/TestAPI.vue'
import TableTask from './components/TableTask.vue'

import axios from 'axios';

export default {
  name: 'App',
  components: {
    // TestAPI,
    // TableCompany,
    TableTask,
  },
  data(){
    return{
      counter: 0,
      info: {}
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
      .then((response) => {
        this.info = response.data
        console.log(this.info)
      })
    },
    reqCompany(){
      axios
      .get("http://localhost:9000/Company")
      .then((response) => {
        this.info = response.data
        console.log(this.info)
      })
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
