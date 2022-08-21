<template>
    <table>
        <tr>
            <th>Id</th>
            <th>Имя</th>
        </tr>
        <tr v-for="item in items" :key="item.Id">
            <td>{{ item.Id }}</td>
            <td>{{ item.Name }}</td>
        </tr>
    </table>
    <button @click="getCompany">Получить организации</button>
</template>

<script>
// import axios from 'axios'

import axios from 'axios'


export default {
    data() {
        return {
            items: []
        }
    },
    mounted() {

    },

    methods: {
        getCompany() {
            axios.get("http://localhost:9000/Company")
            .then(response => {
                this.items = response.data
            })
        }
    }
}
</script>

<style>
</style>