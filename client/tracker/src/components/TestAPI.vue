<template>
    <div>
        <br>
        <button v-for="(but, ind) in butSwitch" :key="ind" @click="setApi(but)">{{but}}</button>
        <br>
        <input id="api_str" v-model="api_str" type="text" name="api" placeholder="api str">
        <br>
        <button @click="req">Отправить GET {{'http://localhost:9000' + api_str}}</button>
        
        <div>{{ info }}</div>
    </div>
</template>
<script>
import axios from 'axios';

export default {
    // name: 'TestAPI',
    data() {
        return {
            info: {},
            api_str: "",
            butSwitch: [
                "/", 
                "/Company",
                "/Task",
                "/Task_View",
            ],   
        }
    },
    methods: {
        setApi(api_args){
            this.api_str = api_args;
        },
        req() {
            let test = 'http://localhost:9000' + this.api_str
            console.log(test)
            axios.get(test)
            .then((response) => {
                this.info = response.data
                console.log(this.info)
            })
        },
    }
}
</script>
