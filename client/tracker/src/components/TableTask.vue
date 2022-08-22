<template>
    <!-- <ButtonRegular @click="getTable">Получить задачи</ButtonRegular> -->

    <div class="ComandPanel">
        <!-- <ButtonRegular @click="addTask">Создать задачу</ButtonRegular> -->
        <ButtonRegular>Редактировать выбранный</ButtonRegular>
        <ButtonRegular>Удалить выбранные</ButtonRegular>
    </div>

    <div class="tableTask">

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
            <tbody v-if="ReqSucces" >
                    <tr  v-for="(task, index) in TaskView" :key="index" @dblclick="dblclick_record(index)">
                        <td class="CheckBox"><input type="checkbox" v-bind:name="('task_' + index)" class="inputCheck" /></td>
                        <td>{{ task.id }}</td>
                        <td>{{ task.Company }}</td>
                        <td>{{ task.Project }}</td>
                        <td>{{ task.Name_task }}</td>
                        <td>{{ task.Description_task }}</td>
                        <td>{{ task.Hour }}</td>
                        <td>{{ task.Minute }}</td>
                        <td class="CheckBox"><input type="checkbox" v-bind:name="('taskClosed_' + index)" class="inputCheck" v-model="task.Closed"/></td>
                    </tr>
                    <tr class="NewTask">
                        <td></td>
                        <td></td>
                        <td>
                            <select name="Company" id="Company">
                                <option v-for="company in NewTask.ListCompany" value="{{company.id}}" :key="company.Id">{{company.Name}}</Option>
                            </select>
                        </td>
                        <td>
                            <select name="Project" id="Project">
                                <option v-for="project in NewTask.ListProject" value="{{project.Id}}" :key="project.Id">{{project.Name_project}}</Option>
                            </select>
                        </td>
                        <td><input type="text"></td>
                        <td><input type="text"></td>
                        <td><input type="number"></td>
                        <td><input type="number"></td>
                        <td><ButtonRegular @click="addTask">Добавить</ButtonRegular></td>
                    </tr>
            </tbody>
            <tbody v-else>    
                <tr>
                    <td style="text-align: center;" colspan="9">Ничего не вышло, попробуйте в следующий раз</td>
                </tr>
            </tbody>
            
        </table>

    <!-- <PopupTask></PopupTask> -->
    </div>

</template>

<script>
import ButtonRegular from './button/ButtonRegular.vue';
// import PopupTask from './PopupTask.vue';

import axios from 'axios';

export default {
    name: 'TableTask',
    components: {
        ButtonRegular,
        // PopupTask
    },
    data() {
        return {
            TaskView: [],
            ReqSucces: true,
            showModal: true,
            ListCompany:[],
            ListProject:[],
            NewTask:
            {
                ListCompany:[],
                ListProject: [],
                Name_task: "",
                Description_task:"",
                Hour: 0,
                Minute: 0,
                Link_our_tracker: "",
                Total_minutes: 0,
                Closed: false,
            }
        }
    },
    mounted() {
        const par = {id: 15, name: "privit"}
        axios.post("http://localhost:9000/TaskView", par)
        .then((response) => {
            this.TaskView = response.data
            this.ReqSucces = true
            console.log(response.data);
        })
        .catch(()=>{
            this.ReqSucces = false
        })

        axios.get("http://localhost:9000/Company")
        .then((response) => {
            this.NewTask.ListCompany = response.data
            console.log(response.data);
        })

        axios.get("http://localhost:9000/Project")
        .then((response) => {
            this.NewTask.ListProject = response.data
            console.log(this.NewTask.ListProject);
        })
    },
    methods: {
        addTask(){
            
        },
        dblclick_record(index){
            window.getSelection().removeAllRanges();
            console.log(this.TaskView[index]);
            this.showModal = true
        }
        
    }
}
</script>

<style>
.tableTask{
    margin-top: 15px;
    display: flex;
    justify-content: center;
}


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

tbody{
    cursor: pointer;
}

tbody tr:hover td{
    background: #AFCDE7;
}


.inputCheck {
    width: 20px;
    height: 20px;
}

.CheckBox {
    text-align: center;
    padding: 0;
}

.ComandPanel{
    margin-top: 15px;
}

.ComandPanel > *{
    margin: 0 2px;
}
</style>