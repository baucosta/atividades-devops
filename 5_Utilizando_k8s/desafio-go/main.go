package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func greeting(msg string) string {
	return "<b>" + msg + "</b>"
}

func majorRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, greeting("Code.education Rocks!"))
}

func setRoutes(router *mux.Router) {
	router.HandleFunc("/", majorRoute)
}

func main() {
	var errorServer error
	var router *mux.Router

	log.Println("Server started on: https://localhost:8000")

	router = mux.NewRouter()

	setRoutes(router)

	errorServer = http.ListenAndServe(":8000", router)
	if errorServer != nil {
		log.Println(errorServer)
	}

}
