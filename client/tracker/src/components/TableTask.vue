<template>
    <ButtonRegular @click="getTable">Получить задачи</ButtonRegular>

    <div>
        <ButtonRegular>Создать задачу</ButtonRegular>
        <ButtonRegular>Редактировать выбранный</ButtonRegular>
        <ButtonRegular>Удалить</ButtonRegular>
    </div>

    <table>
        <thead>
            <tr>
                <th></th>
                <th>id</th>
                <th>Организация</th>
                <th>Проект</th>
                <th>Задача</th>
                <th>Описание</th>
                <th>Часы</th>
                <th>Минуты</th>
                <th>Закрыта</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(task, index) in TaskView" :key="index">
                <td class="CheckBox"><input type="checkbox" name="name{{index}}" class="inputCheck" /></td>
                <td>{{ task.id }}</td>
                <td>{{ task.Company }}</td>
                <td>{{ task.Project }}</td>
                <td>{{ task.Name_task }}</td>
                <td>{{ task.Description_task }}</td>
                <td>{{ task.Hour }}</td>
                <td>{{ task.Minute }}</td>
                <td>{{ task.Closed }}</td>
            </tr>
        </tbody>
    </table>
</template>

<script>
import ButtonRegular from './button/ButtonRegular.vue';

import axios from 'axios';

export default {
    name: 'TableTask',
    components: {
        ButtonRegular,
    },
    data() {
        return {
            TaskView: [],
        }
    },

    methods: {
        getTable() {
            axios.get("http://localhost:9000/Task_View")
                .then((response) => {
                    this.TaskView = response.data
                })
        },
    }
}
</script>

<style>
table {
    font-family: "Lucida Sans Unicode", "Lucida Grande", Sans-Serif;
    font-size: 14px;
    border-collapse: collapse;
    text-align: center;
}

th,
td:first-child {
    background: #AFCDE7;
    color: white;
    /* padding: 10px 20px; */
}

th,
td {
    border-style: solid;
    border-width: 0 1px 1px 0;
    border-color: white;
    padding: 15px 20px;
}

td {
    padding: 5px;
    background: #D8E6F3;
}

th:first-child,
td:first-child {
    text-align: left;
}


.inputCheck{
    width: 20px;
    height: 20px;
}

.CheckBox{
    text-align: center;
    padding: 0;
}

</style>