package serve

import (
	"log"
	"net/http"
	"strconv"
)

const (
	PORT = 9000
)

func HttpServe() {
	mux := http.NewServeMux()

	SetRout(mux)

	log.Println("Запуск веб-сервера на http://127.0.0.1:9000")
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), mux)
	if err != nil {
		log.Fatal(err)
	}
}

func setHeader(w http.ResponseWriter) {
	//w.WriteHeader(http.StatusCreated)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}
