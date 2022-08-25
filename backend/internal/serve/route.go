package serve

import (
	"encoding/json"
	"fmt"
	"log"
	"main_mod/internal/bd/controller"
	"net/http"
)

func rout() map[string]func(http.ResponseWriter, *http.Request) {

	rout := make(map[string]func(http.ResponseWriter, *http.Request))
	rout["/"] = r_Home
	rout["/Company"] = r_Company
	rout["/Task"] = r_Task
	rout["/TaskView"] = r_TaskView
	rout["/Project"] = r_Project
	rout["/AddTask"] = r_AddTask

	return rout
}

func SetRout(mux *http.ServeMux) {
	route := rout()
	for i, rout := range route {
		mux.HandleFunc(i, rout)
	}
}

// ++Router

func r_Home(w http.ResponseWriter, r *http.Request) {
	defer fmt.Println("Запрос прошел")
	setHeader(w)

	resp := make(map[string]string)
	resp["message"] = "Status Created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func r_Company(w http.ResponseWriter, r *http.Request) {
	defer fmt.Println("Запрос company прошел")
	setHeader(w)

	jsonResp, err := json.Marshal(controller.AllCompany())

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func r_Task(w http.ResponseWriter, r *http.Request) {
	defer fmt.Println("Запрос Task прошел")
	setHeader(w)

	jsonResp, err := json.Marshal(controller.AllTask())

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func r_TaskView(w http.ResponseWriter, r *http.Request) {
	defer fmt.Println("Запрос Task_View прошел")

	setHeader(w)

	jsonResp, err := json.Marshal(controller.AllTask_view())

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func r_Project(w http.ResponseWriter, r *http.Request) {
	defer fmt.Println("Запрос Project прошел")
	setHeader(w)

	jsonResp, err := json.Marshal(controller.AllProgect())

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func r_AddTask(w http.ResponseWriter, r *http.Request) {
	defer fmt.Println("Запрос AddTask прошел")
	setHeader(w)

	r.ParseForm()
	x := r.Form.Get("Name_task")
	fmt.Println(x)
	// jsonResp, err := json.Marshal(controller.AddTask())

	// if err != nil {
	// 	log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	// }
	// w.Write(jsonResp)
}

// --Router
