package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	serviceName := "undefined"
	if os.Getenv("SERVICE_NAME") != "" {
		serviceName = os.Getenv("SERVICE_NAME")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello " + serviceName))
		log.Println("Hello", serviceName)
	}).Methods("GET")

	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Println(err)
	}
}
