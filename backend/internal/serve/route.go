package serve

import (
	"encoding/json"
	"fmt"
	"log"
	"main_mod/internal/bd/controller"
	"main_mod/internal/bd/models"
	"net/http"
	"os"
)

func writeCommentAfterRequest(text string) {

	if WriteHttpComment := os.Getenv("WriteHttpComment"); WriteHttpComment == "1" {
		fmt.Println(text)
	}
}

func rout() map[string]func(http.ResponseWriter, *http.Request) {

	rout := make(map[string]func(http.ResponseWriter, *http.Request))
	rout["/"] = r_Home
	rout["/Company"] = r_Company
	rout["/Task"] = r_Task
	rout["/TaskView"] = r_TaskView
	rout["/Project"] = r_Project
	rout["/AddTask"] = r_AddTask
	rout["/DeleteTask"] = r_DeleteTask

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
	defer writeCommentAfterRequest("Запрос прошел")
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
	defer writeCommentAfterRequest("Запрос company прошел")
	setHeader(w)

	jsonResp, err := json.Marshal(controller.AllCompany())

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func r_Task(w http.ResponseWriter, r *http.Request) {
	defer writeCommentAfterRequest("Запрос Task прошел")
	setHeader(w)

	jsonResp, err := json.Marshal(controller.AllTask())

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func r_TaskView(w http.ResponseWriter, r *http.Request) {
	defer writeCommentAfterRequest("Запрос Task_View прошел")

	setHeader(w)

	jsonResp, err := json.Marshal(controller.AllTask_view())

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func r_Project(w http.ResponseWriter, r *http.Request) {
	defer writeCommentAfterRequest("Запрос Project прошел")
	setHeader(w)

	jsonResp, err := json.Marshal(controller.AllProgect())

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func r_AddTask(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	if r.Method == http.MethodOptions {
		defer writeCommentAfterRequest("Запрос AddTask прошел (Options)")
		return
	}
	writeCommentAfterRequest("Запрос AddTask прошел")

	decoder := json.NewDecoder(r.Body)

	var NewTask models.FTask
	err := decoder.Decode(&NewTask)
	if err != nil {
		log.Printf("%#v\n", err)
	}

	err = controller.AddTask(NewTask)

	var returnStruct struct {
		Msg string `json:"msg"`
	}
	if err != nil {
		returnStruct.Msg = err.Error()
	}

	jsonResp, err := json.Marshal(returnStruct)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func r_DeleteTask(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	if r.Method == http.MethodOptions {
		defer writeCommentAfterRequest("Запрос r_DeleteTask прошел (Options)")
		return
	}
	writeCommentAfterRequest("Запрос DeleteTask прошел")

	decoder := json.NewDecoder(r.Body)

	var sliceInt []int
	err := decoder.Decode(&sliceInt)
	if err != nil {
		log.Printf("%#v\n", err)
	} else {
		fmt.Println(sliceInt)
	}

	err = controller.DeleteTask(sliceInt)

	var returnStruct struct {
		Msg string `json:"msg"`
	}
	if err != nil {
		returnStruct.Msg = err.Error()
	}

	jsonResp, err := json.Marshal(returnStruct)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

// --Router
