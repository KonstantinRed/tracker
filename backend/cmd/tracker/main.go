package main

import (
	// пакет для форматированного ввода вывода
	"encoding/json"
	"fmt"
	"log"      // пакет для логирования
	"net/http" // пакет для поддержки HTTP протокола

	// пакет для работы с  UTF-8 строками

	"main_mod/internal/tracker/bd/controller"
)

func main() {
	http.HandleFunc("/", HomeRouterHandler)           // установим роутер
	http.HandleFunc("/Company", CompanyRouterHandler) // установим роутер
	err := http.ListenAndServe(":9000", nil)          // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
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

	resp := controller.AllCompany()

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func cors(w http.ResponseWriter) {
	//++CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//--CORS
}
