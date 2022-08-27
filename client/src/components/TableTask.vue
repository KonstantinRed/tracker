<template>
    <!-- <ButtonRegular @click="getTable">Получить задачи</ButtonRegular> -->

    <div class="ComandPanel">
        <!-- <ButtonRegular @click="addTask">Создать задачу</ButtonRegular> -->
        <ButtonRegular>Редактировать выбранный</ButtonRegular>
        <ButtonRegular @click="deleteTask">Удалить выбранные</ButtonRegular>
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
            <tbody v-if="ErrReq == undefined" >
                    <tr  v-for="(task, index) in TaskView" :key="index" @dblclick="dblclick_record(index)">
                        <td class="CheckBox"><input type="checkbox" v-bind:value="(task.id)" class="inputCheck" v-model="arrayCheckedTask"/></td>
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
                            <select name="Company" id="Company" v-model="NewTask.id_company">
                                <option v-for="company in ListCompany" :value="company.Id" :key="company.Id">{{company.Name}}</Option>
                            </select>
                        </td>
                        <td>
                            <select name="Project" id="Project" v-model="NewTask.id_project">
                                <option v-for="project in ListProject" :value="project.Id" :key="project.Id">{{project.Name_project}}</Option>
                            </select>
                        </td>
                        <td><input type="text" v-model="NewTask.Name_task"></td>
                        <td><input type="text" v-model="NewTask.Description_task"></td>
                        <td><input type="number" v-model="NewTask.Hour"></td>
                        <td><input type="number" v-model="NewTask.Minute"></td>
                        <td><ButtonRegular @click="addTask">Добавить</ButtonRegular></td>
                        <TD><ButtonRegular @click="Test">Тест</ButtonRegular></TD>
                    </tr>
            </tbody>
            <tbody v-else>    
                <tr>
                    <td style="text-align: center;" colspan="9">Ничего не вышло, попробуйте в следующий раз <span>{{ErrReq}}</span></td>
                </tr>
            </tbody>
            
        </table>

    <!-- <PopupTask></PopupTask> -->
    </div>

</template>

<script>
import ButtonRegular from './button/ButtonRegular.vue';
import axios from 'axios';

export default {
    name: 'TableTask',
    components: {
        ButtonRegular,
        // PopupTask
    },
    data() {
        return {
            arrayCheckedTask: [],
            TaskView: [],
            ErrReq: undefined,
            showModal: true,
            ListCompany:[
                    {Id: 0, Name: "---"}
            ],
            ListProject: [
                    {Id: 0, Name_project: "---"}
            ],
            NewTask:
            {
                
                id_company: 0,
                id_project: 0,
                Name_task: "Задача1",
                Description_task:"Простейшее описание задачи",
                Hour: 2,
                Minute: 30,
                Link_our_tracker: "",
                Closed: false,
                Total_minutes: 0,                
            }
        }
    },
    mounted() {
        this.startComponent()
    },
    methods: {
        Test(){
            console.log(this.NewTask);
        },
        addTask(){
            let NewTask = this.NewTask
            axios.post("http://localhost:9000/AddTask", 
                NewTask, 
                {headers: {'Content-Type': 'application/json'}})
            .then((res)=>{
                if (!res.data.msg){
                    console.log("Добавлена запись")
                }else{
                    throw new Error(res.data.msg)
                }
                this.startComponent()

            })
            .catch(err=>console.log(err))
        },
        deleteTask(){
            axios.post("http://localhost:9000/DeleteTask", 
                this.arrayCheckedTask, 
                {headers: {'Content-Type': 'application/json'}})
            .then((res)=>{
                if (!res.data.msg){
                    console.log("Удалены записи")
                }else{
                    throw new Error(res.data.msg)
                }
                this.startComponent()

            })
            .catch(err=>console.log(err))
        },
        dblclick_record(index){
            window.getSelection().removeAllRanges();
            console.log(this.TaskView[index]);
            this.showModal = true
        },
        startComponent(){
            axios.get("http://localhost:9000/TaskView")
            .then((response) => {
                this.TaskView = response.data
                this.ErrReq = undefined
                this.arrayCheckedTask = []
            })
            .catch((err)=>{
                this.ErrReq = err;
            })

            axios.get("http://localhost:9000/Company")
            .then((response) => {
                this.ListCompany = [ {Id: 0, Name: "---"}]
                response.data.map((el)=>{this.ListCompany.push(el)})
            })

            axios.get("http://localhost:9000/Project")
            .then((response) => {
                this.ListProject = [ {Id: 0, Name_project: "---"}]
                response.data.map((el)=>{this.ListProject.push(el)})
            })
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