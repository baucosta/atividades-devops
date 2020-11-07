package main

import (
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/gorilla/mux"
)

func squareRoot(value float64) float64 {
	var res float64

	for i := 0; i < 10000000; i++ {
		res += math.Sqrt(value)
	}

	return res
}

func majorRoute(w http.ResponseWriter, r *http.Request) {
	var value float64

	value = 0.0001

	res := fmt.Sprintf("%.2f", squareRoot(value))

	fmt.Fprintf(w, "<h1>Result: "+res+"</h1><br /><h2>Code education rocks!</h2>")
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
