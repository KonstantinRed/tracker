package serve

import (
	"log"
	"net/http"
)

func HttpServe() {
	mux := http.NewServeMux()
	//++Роутинг
	SetRout(mux)
	// mux.HandleFunc("/", HomeRouterHandler)
	// mux.HandleFunc("/Company", CompanyRouterHandler)
	// mux.HandleFunc("/Task", TaskRouterHandler)
	// mux.HandleFunc("/Task_View", Task_ViewRouterHandler)
	// mux.HandleFunc("/Project", Project_RouterHandler)
	//--Роутинг

	log.Println("Запуск веб-сервера на http://127.0.0.1:9000")
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func setHeader(w http.ResponseWriter) {
	// w.WriteHeader(http.StatusCreated)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}
