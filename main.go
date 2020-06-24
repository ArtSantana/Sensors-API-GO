package main

import (
	"fmt"
	"log"
	"net/http"
	"sensores/controller"
	"time"

	"github.com/gorilla/mux"
)

func doEvery(ms time.Duration, repeatFunc func(time.Time)) {
	for x := range time.Tick(ms) {
		repeatFunc(x)
	}
}

func main() {
	router := mux.NewRouter()
	const port string = ":3000"
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Why are u here?")
	})
	router.HandleFunc("/sensors", controller.GetSensors).Methods("GET")

	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
