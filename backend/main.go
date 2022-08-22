package main

import (
	// пакет для форматированного ввода вывода
	"encoding/json"
	"fmt"
	"log"      // пакет для логирования
	"net/http" // пакет для поддержки HTTP протокола

	// пакет для работы с  UTF-8 строками

	"main_mod/internal/bd/controller"

	env "github.com/joho/godotenv"
)

const (
	PORT = "9000"
)

func set_env() {
	err := env.Load()
	if err != nil {
		log.Panicf("Error loading env:")
	}
}

func main() {
	set_env()

	http.HandleFunc("/", HomeRouterHandler)               // установим роутер
	http.HandleFunc("/Company", CompanyRouterHandler)     // установим роутер
	http.HandleFunc("/Task", TaskRouterHandler)           // установим роутер
	http.HandleFunc("/Task_View", Task_ViewRouterHandler) // установим роутер
	http.HandleFunc("/Project", Project_RouterHandler)    // установим роутер

	err := http.ListenAndServe(":"+PORT, nil) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("Run serve: localhost:" + PORT)
}

func cors(w http.ResponseWriter) {
	//++CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//--CORS
}

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	defer fmt.Println("Запрос прошел")

	cors(w)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["message"] = "Status Created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)

}

func CompanyRouterHandler(w http.ResponseWriter, r *http.Request) {
	defer fmt.Println("Запрос company прошел")

	cors(w)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(controller.AllCompany())

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func TaskRouterHandler(w http.ResponseWriter, r *http.Request) {
	defer fmt.Println("Запрос Task прошел")

	cors(w)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(controller.AllTask())

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func Task_ViewRouterHandler(w http.ResponseWriter, r *http.Request) {
	defer fmt.Println("Запрос Task_View прошел")

	cors(w)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(controller.AllTask_view())

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func Project_RouterHandler(w http.ResponseWriter, r *http.Request) {
	defer fmt.Println("Запрос Project прошел")

	cors(w)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(controller.AllProgect())

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
